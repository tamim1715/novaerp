package journal

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	Create(ctx context.Context, entry *JournalEntry) error
	FindAll(ctx context.Context, req pagination.PageRequest, status, sourceType string, startDate, endDate *time.Time) ([]JournalEntry, int64, error)
	FindByID(ctx context.Context, id string) (*JournalEntry, error)
	FindByEntryNumber(ctx context.Context, num string) (*JournalEntry, error)
	PostEntry(ctx context.Context, id string, periodID *uuid.UUID, postingDate time.Time) (*JournalEntry, error)
	VoidEntry(ctx context.Context, originalID string, voidReason string, reversalEntry *JournalEntry) (*JournalEntry, *JournalEntry, error)
	NextEntryNumber(ctx context.Context, date time.Time) (string, error)
	GetGeneralLedgerLines(ctx context.Context, accountID *uuid.UUID, startDate, endDate time.Time) ([]JournalEntryLine, error)
	GetOpeningBalance(ctx context.Context, accountID uuid.UUID, beforeDate time.Time) (float64, float64, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) NextEntryNumber(ctx context.Context, date time.Time) (string, error) {
	prefix := fmt.Sprintf("JE-%s-", date.Format("200601"))
	var count int64
	err := r.db.WithContext(ctx).Model(&JournalEntry{}).
		Where("entry_number LIKE ?", prefix+"%").
		Count(&count).Error
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%05d", prefix, count+1), nil
}

func (r *repository) Create(ctx context.Context, entry *JournalEntry) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if entry.EntryNumber == "" {
			num, err := r.NextEntryNumber(ctx, entry.EntryDate)
			if err != nil {
				return err
			}
			entry.EntryNumber = num
		}
		return tx.Create(entry).Error
	})
}

func (r *repository) FindAll(ctx context.Context, req pagination.PageRequest, status, sourceType string, startDate, endDate *time.Time) ([]JournalEntry, int64, error) {
	var entries []JournalEntry
	var total int64

	req.Normalize()
	db := r.db.WithContext(ctx).Model(&JournalEntry{}).
		Preload("Lines.Account").
		Preload("Period")

	if status != "" {
		db = db.Where("status = ?", status)
	}
	if sourceType != "" {
		db = db.Where("source_type = ?", sourceType)
	}
	if startDate != nil {
		db = db.Where("entry_date >= ?", startDate.Format("2006-01-02"))
	}
	if endDate != nil {
		db = db.Where("entry_date <= ?", endDate.Format("2006-01-02"))
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := req.Offset()
	err := db.Order(req.SortBy + " " + req.Order).
		Limit(req.Size).
		Offset(offset).
		Find(&entries).Error

	return entries, total, err
}

func (r *repository) FindByID(ctx context.Context, id string) (*JournalEntry, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}
	var entry JournalEntry
	err = r.db.WithContext(ctx).
		Preload("Lines.Account").
		Preload("Period").
		First(&entry, "id = ?", uid).Error
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *repository) FindByEntryNumber(ctx context.Context, num string) (*JournalEntry, error) {
	var entry JournalEntry
	err := r.db.WithContext(ctx).
		Preload("Lines.Account").
		Preload("Period").
		First(&entry, "entry_number = ?", num).Error
	if err != nil {
		return nil, err
	}
	return &entry, nil
}

func (r *repository) PostEntry(ctx context.Context, id string, periodID *uuid.UUID, postingDate time.Time) (*JournalEntry, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, gorm.ErrRecordNotFound
	}

	var entry JournalEntry
	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Preload("Lines.Account").
			First(&entry, "id = ?", uid).Error; err != nil {
			return err
		}

		if entry.Status == StatusPosted {
			return errors.New("journal entry is already posted and immutable")
		}
		if entry.Status == StatusVoid {
			return errors.New("cannot post a voided journal entry")
		}

		entry.Status = StatusPosted
		entry.PostingDate = &postingDate
		if periodID != nil {
			entry.PeriodID = periodID
		}

		return tx.Save(&entry).Error
	})

	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (r *repository) VoidEntry(ctx context.Context, originalID string, voidReason string, reversalEntry *JournalEntry) (*JournalEntry, *JournalEntry, error) {
	uid, err := uuid.Parse(originalID)
	if err != nil {
		return nil, nil, gorm.ErrRecordNotFound
	}

	var original JournalEntry
	err = r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).
			Preload("Lines").
			First(&original, "id = ?", uid).Error; err != nil {
			return err
		}

		if original.Status != StatusPosted {
			return errors.New("only posted journal entries can be voided with a reversal entry")
		}

		// Update original entry
		original.Status = StatusVoid
		original.VoidReason = voidReason
		if err := tx.Save(&original).Error; err != nil {
			return err
		}

		// Create reversal entry
		if reversalEntry.EntryNumber == "" {
			num, err := r.NextEntryNumber(ctx, reversalEntry.EntryDate)
			if err != nil {
				return err
			}
			reversalEntry.EntryNumber = num
		}
		reversalEntry.ReversalOfID = &original.ID

		return tx.Create(reversalEntry).Error
	})

	if err != nil {
		return nil, nil, err
	}

	return &original, reversalEntry, nil
}

func (r *repository) GetGeneralLedgerLines(ctx context.Context, accountID *uuid.UUID, startDate, endDate time.Time) ([]JournalEntryLine, error) {
	var lines []JournalEntryLine

	query := r.db.WithContext(ctx).
		Joins("JOIN journal_entries ON journal_entries.id = journal_entry_lines.journal_entry_id").
		Where("journal_entries.status = ?", StatusPosted).
		Where("journal_entries.entry_date >= ?", startDate.Format("2006-01-02")).
		Where("journal_entries.entry_date <= ?", endDate.Format("2006-01-02")).
		Preload("Account").
		Order("journal_entries.entry_date ASC, journal_entries.entry_number ASC")

	if accountID != nil {
		query = query.Where("journal_entry_lines.account_id = ?", *accountID)
	}

	err := query.Find(&lines).Error
	return lines, err
}

func (r *repository) GetOpeningBalance(ctx context.Context, accountID uuid.UUID, beforeDate time.Time) (float64, float64, error) {
	type Result struct {
		TotalDebit  float64 `gorm:"column:total_debit"`
		TotalCredit float64 `gorm:"column:total_credit"`
	}
	var res Result
	err := r.db.WithContext(ctx).Table("journal_entry_lines").
		Select("COALESCE(SUM(debit), 0) AS total_debit, COALESCE(SUM(credit), 0) AS total_credit").
		Joins("JOIN journal_entries ON journal_entries.id = journal_entry_lines.journal_entry_id").
		Where("journal_entries.status = ?", StatusPosted).
		Where("journal_entry_lines.account_id = ?", accountID).
		Where("journal_entries.entry_date < ?", beforeDate.Format("2006-01-02")).
		Scan(&res).Error

	return res.TotalDebit, res.TotalCredit, err
}
