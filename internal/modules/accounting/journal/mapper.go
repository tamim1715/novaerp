package journal

import (
	"github.com/tamim1715/novaerp/internal/modules/accounting/account"
)

func ToJournalEntryLineResponse(l *JournalEntryLine) JournalEntryLineResponse {
	if l == nil {
		return JournalEntryLineResponse{}
	}
	res := JournalEntryLineResponse{
		ID:          l.ID,
		AccountID:   l.AccountID,
		Debit:       l.Debit,
		Credit:      l.Credit,
		Description: l.Description,
		PartnerType: l.PartnerType,
		PartnerID:   l.PartnerID,
	}
	if l.Account.ID != l.AccountID && l.Account.ID.String() != "00000000-0000-0000-0000-000000000000" {
		res.AccountCode = l.Account.Code
		res.AccountName = l.Account.Name
		res.AccountType = l.Account.Type
		accResp := account.ToAccountResponse(&l.Account)
		res.Account = &accResp
	}
	return res
}

func ToJournalEntryLineResponseList(lines []JournalEntryLine) []JournalEntryLineResponse {
	res := make([]JournalEntryLineResponse, len(lines))
	for i, l := range lines {
		res[i] = ToJournalEntryLineResponse(&l)
	}
	return res
}

func ToJournalEntryResponse(e *JournalEntry) JournalEntryResponse {
	if e == nil {
		return JournalEntryResponse{}
	}
	res := JournalEntryResponse{
		ID:           e.ID,
		EntryNumber:  e.EntryNumber,
		EntryDate:    e.EntryDate,
		PostingDate:  e.PostingDate,
		PeriodID:     e.PeriodID,
		Reference:    e.Reference,
		SourceType:   e.SourceType,
		SourceID:     e.SourceID,
		Description:  e.Description,
		Status:       e.Status.String(),
		TotalDebit:   e.TotalDebit,
		TotalCredit:  e.TotalCredit,
		PostedBy:     e.PostedBy,
		VoidReason:   e.VoidReason,
		ReversalOfID: e.ReversalOfID,
		CreatedAt:    e.CreatedAt,
		UpdatedAt:    e.UpdatedAt,
	}
	if len(e.Lines) > 0 {
		res.Lines = ToJournalEntryLineResponseList(e.Lines)
	}
	return res
}

func ToJournalEntryResponseList(entries []JournalEntry) []JournalEntryResponse {
	res := make([]JournalEntryResponse, len(entries))
	for i, e := range entries {
		res[i] = ToJournalEntryResponse(&e)
	}
	return res
}
