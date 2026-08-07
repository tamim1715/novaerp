package journal

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateJournalEntryDoc creates a balanced double-entry transaction
// @Summary Create Journal Entry
// @Description Creates a multi-line balanced debit/credit journal transaction (draft or auto-posted)
// @Tags Accounting - Journal Entries & Ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateJournalEntryRequest true "Journal Entry Payload"
// @Success 201 {object} response.APIResponse{data=JournalEntryResponse} "Journal entry created successfully"
// @Failure 400 {object} response.APIResponse "Unbalanced debits/credits or invalid line items"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/journals [post]
func (h *Handler) CreateJournalEntryDoc(c *gin.Context) {
	h.CreateJournalEntry(c)
}

// FindAllJournalEntriesDoc retrieves paginated journal entries
// @Summary List Journal Entries
// @Description Retrieve a paginated list of journal entries with date, status, and source filters
// @Tags Accounting - Journal Entries & Ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param size query int false "Items per page" default(10)
// @Param sortBy query string false "Sort by field" default(entry_date)
// @Param order query string false "Sort order (asc/desc)" default(desc)
// @Param status query string false "Filter by Status (DRAFT, POSTED, VOID)"
// @Param sourceType query string false "Filter by Source (MANUAL, PAYROLL, INVOICE, PAYMENT, STOCK_ADJUSTMENT)"
// @Param startDate query string false "Filter start date (YYYY-MM-DD)"
// @Param endDate query string false "Filter end date (YYYY-MM-DD)"
// @Success 200 {object} response.APIResponse{data=pagination.PageResponse{data=[]JournalEntryResponse}} "Journal entries retrieved"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/journals [get]
func (h *Handler) FindAllJournalEntriesDoc(c *gin.Context) {
	h.FindAllJournalEntries(c)
}

// FindJournalEntryByIDDoc retrieves a single journal entry
// @Summary Get Journal Entry by ID
// @Description Retrieve a journal entry with all line items and ledger details
// @Tags Accounting - Journal Entries & Ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Journal Entry UUID"
// @Success 200 {object} response.APIResponse{data=JournalEntryResponse} "Journal entry fetched successfully"
// @Failure 404 {object} response.APIResponse "Journal entry not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/journals/{id} [get]
func (h *Handler) FindJournalEntryByIDDoc(c *gin.Context) {
	h.FindJournalEntryByID(c)
}

// PostJournalEntryDoc commits and locks a draft journal entry to the General Ledger
// @Summary Post Journal Entry
// @Description Commits a draft journal entry, stamps the posting timestamp, and freezes it from modification
// @Tags Accounting - Journal Entries & Ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Journal Entry UUID"
// @Success 200 {object} response.APIResponse{data=JournalEntryResponse} "Journal entry posted and locked"
// @Failure 400 {object} response.APIResponse "Already posted or closed period"
// @Failure 404 {object} response.APIResponse "Journal entry not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/journals/{id}/post [post]
func (h *Handler) PostJournalEntryDoc(c *gin.Context) {
	h.PostJournalEntry(c)
}

// VoidJournalEntryDoc voids a posted journal entry by generating an immutable reversal transaction
// @Summary Void Journal Entry (Generate Reversal)
// @Description Voids a posted journal entry by creating an opposite offsetting reversal entry in the ledger
// @Tags Accounting - Journal Entries & Ledger
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Journal Entry UUID"
// @Param request body VoidJournalEntryRequest true "Void Reason"
// @Success 200 {object} response.APIResponse "Journal entry voided and reversal posted"
// @Failure 400 {object} response.APIResponse "Only posted entries can be voided"
// @Failure 404 {object} response.APIResponse "Journal entry not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/journals/{id}/void [post]
func (h *Handler) VoidJournalEntryDoc(c *gin.Context) {
	h.VoidJournalEntry(c)
}
