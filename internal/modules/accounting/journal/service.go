package journal

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/tamim1715/novaerp/internal/common/pagination"
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
	"github.com/tamim1715/novaerp/internal/modules/accounting/period"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Service interface {
	CreateJournalEntry(ctx context.Context, req CreateJournalEntryRequest, postedBy *uuid.UUID) (*JournalEntry, error)
	FindAllJournalEntries(ctx context.Context, req pagination.PageRequest, status, sourceType string, startDate, endDate *time.Time) ([]JournalEntry, int64, error)
	FindJournalEntryByID(ctx context.Context, id string) (*JournalEntry, error)
	PostJournalEntry(ctx context.Context, id string) (*JournalEntry, error)
	VoidJournalEntry(ctx context.Context, id string, req VoidJournalEntryRequest, voidedBy *uuid.UUID) (*JournalEntry, *JournalEntry, error)
}

type service struct {
	repo        Repository
	accountRepo account.Repository
	periodRepo  period.Repository
	logger      *zap.Logger
}

func NewService(repo Repository, accountRepo account.Repository, periodRepo period.Repository, logger *zap.Logger) Service {
	return &service{
		repo:        repo,
		accountRepo: accountRepo,
		periodRepo:  periodRepo,
		logger:      logger,
	}
}

func (s *service) CreateJournalEntry(ctx context.Context, req CreateJournalEntryRequest, postedBy *uuid.UUID) (*JournalEntry, error) {
	if len(req.Lines) < 2 {
		return nil, errors.New("journal entry must contain at least 2 balanced line items (debit and credit)")
	}

	var totalDebit, totalCredit float64
	lines := make([]JournalEntryLine, len(req.Lines))

	for i, l := range req.Lines {
		if l.Debit < 0 || l.Credit < 0 {
			return nil, fmt.Errorf("line %d: debit and credit amounts cannot be negative", i+1)
		}
		if (l.Debit > 0 && l.Credit > 0) || (l.Debit == 0 && l.Credit == 0) {
			return nil, fmt.Errorf("line %d: line must contain either a debit or credit amount, not both or zero", i+1)
		}

		// Verify account exists and is active
		acc, err := s.accountRepo.FindByID(ctx, l.AccountID.String())
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, fmt.Errorf("line %d: account not found", i+1)
			}
			return nil, err
		}
		if !acc.IsActive {
			return nil, fmt.Errorf("line %d: account %s (%s) is inactive", i+1, acc.Code, acc.Name)
		}

		pType := l.PartnerType
		if pType == "" {
			pType = PartnerNone
		}

		lines[i] = JournalEntryLine{
			AccountID:   l.AccountID,
			Debit:       math.Round(l.Debit*100) / 100,
			Credit:      math.Round(l.Credit*100) / 100,
			Description: l.Description,
			PartnerType: pType,
			PartnerID:   l.PartnerID,
		}

		totalDebit += lines[i].Debit
		totalCredit += lines[i].Credit
	}

	totalDebit = math.Round(totalDebit*100) / 100
	totalCredit = math.Round(totalCredit*100) / 100

	// Double-entry balancing invariant
	if math.Abs(totalDebit-totalCredit) >= 0.001 {
		return nil, fmt.Errorf("unbalanced journal entry: total debits ($%.2f) must equal total credits ($%.2f)", totalDebit, totalCredit)
	}

	if totalDebit <= 0 {
		return nil, errors.New("total debit amount must be greater than zero")
	}

	srcType := req.SourceType
	if srcType == "" {
		srcType = SourceManual
	}

	// Check if matching accounting period is open
	var periodID *uuid.UUID
	p, err := s.periodRepo.FindPeriodByDate(ctx, req.EntryDate)
	if err == nil && p != nil {
		if p.Status == period.StatusClosed {
			return nil, fmt.Errorf("cannot create journal entry in closed accounting period (%s)", p.Name)
		}
		periodID = &p.ID
	}

	status := StatusDraft
	var postingDate *time.Time
	if req.AutoPost {
		status = StatusPosted
		now := time.Now().UTC()
		postingDate = &now
	}

	entry := &JournalEntry{
		EntryDate:   req.EntryDate,
		PostingDate: postingDate,
		PeriodID:    periodID,
		Reference:   req.Reference,
		SourceType:  srcType,
		SourceID:    req.SourceID,
		Description: req.Description,
		Status:      status,
		TotalDebit:  totalDebit,
		TotalCredit: totalCredit,
		PostedBy:    postedBy,
		Lines:       lines,
	}

	if err := s.repo.Create(ctx, entry); err != nil {
		return nil, err
	}

	return s.repo.FindByID(ctx, entry.ID.String())
}

func (s *service) FindAllJournalEntries(ctx context.Context, req pagination.PageRequest, status, sourceType string, startDate, endDate *time.Time) ([]JournalEntry, int64, error) {
	return s.repo.FindAll(ctx, req, status, sourceType, startDate, endDate)
}

func (s *service) FindJournalEntryByID(ctx context.Context, id string) (*JournalEntry, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) PostJournalEntry(ctx context.Context, id string) (*JournalEntry, error) {
	entry, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if entry.Status == StatusPosted {
		return nil, errors.New("journal entry is already posted")
	}
	if entry.Status == StatusVoid {
		return nil, errors.New("cannot post a voided journal entry")
	}

	// Check period status
	var periodID *uuid.UUID
	p, err := s.periodRepo.FindPeriodByDate(ctx, entry.EntryDate)
	if err == nil && p != nil {
		if p.Status == period.StatusClosed {
			return nil, fmt.Errorf("cannot post to closed accounting period (%s)", p.Name)
		}
		periodID = &p.ID
	}

	return s.repo.PostEntry(ctx, id, periodID, time.Now().UTC())
}

func (s *service) VoidJournalEntry(ctx context.Context, id string, req VoidJournalEntryRequest, voidedBy *uuid.UUID) (*JournalEntry, *JournalEntry, error) {
	orig, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, nil, err
	}

	if orig.Status != StatusPosted {
		return nil, nil, errors.New("only posted journal entries can be voided via reversal entry")
	}

	// Build reversal entry with swapped debits & credits
	reversalLines := make([]JournalEntryLine, len(orig.Lines))
	for i, l := range orig.Lines {
		reversalLines[i] = JournalEntryLine{
			AccountID:   l.AccountID,
			Debit:       l.Credit, // Swapped!
			Credit:      l.Debit,  // Swapped!
			Description: fmt.Sprintf("Reversal of %s - %s", orig.EntryNumber, l.Description),
			PartnerType: l.PartnerType,
			PartnerID:   l.PartnerID,
		}
	}

	now := time.Now().UTC()
	reversalEntry := &JournalEntry{
		EntryDate:   now,
		PostingDate: &now,
		PeriodID:    orig.PeriodID,
		Reference:   fmt.Sprintf("REV-%s", orig.EntryNumber),
		SourceType:  SourceReversal,
		SourceID:    &orig.ID,
		Description: fmt.Sprintf("Reversal Entry for %s: %s", orig.EntryNumber, req.Reason),
		Status:      StatusPosted,
		TotalDebit:  orig.TotalCredit, // Equal to original
		TotalCredit: orig.TotalDebit,
		PostedBy:    voidedBy,
		Lines:       reversalLines,
	}

	return s.repo.VoidEntry(ctx, id, req.Reason, reversalEntry)
}
