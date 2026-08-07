package account

import (
	"github.com/gin-gonic/gin"
	"github.com/tamim1715/novaerp/internal/common/response"
)

var (
	_ response.APIResponse
)

// CreateAccountDoc creates a new Chart of Accounts record
// @Summary Create Chart of Accounts entry
// @Description Add a new account to the General Ledger Chart of Accounts (COA)
// @Tags Accounting - Accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body CreateAccountRequest true "Account Payload"
// @Success 201 {object} response.APIResponse{data=AccountResponse} "Account created successfully"
// @Failure 400 {object} response.APIResponse "Bad request"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/accounts [post]
func (h *Handler) CreateAccountDoc(c *gin.Context) {
	h.CreateAccount(c)
}

// FindAllAccountsDoc retrieves paginated accounts
// @Summary List Chart of Accounts
// @Description Retrieve a paginated list of accounts with optional type filter
// @Tags Accounting - Accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param page query int false "Page number" default(1)
// @Param size query int false "Items per page" default(10)
// @Param sortBy query string false "Sort by field" default(code)
// @Param order query string false "Sort order (asc/desc)" default(asc)
// @Param type query string false "Filter by Account Type (ASSET, LIABILITY, EQUITY, REVENUE, EXPENSE)"
// @Success 200 {object} response.APIResponse{data=pagination.PageResponse{data=[]AccountResponse}} "Accounts retrieved successfully"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/accounts [get]
func (h *Handler) FindAllAccountsDoc(c *gin.Context) {
	h.FindAllAccounts(c)
}

// GetAccountTreeDoc retrieves hierarchical Chart of Accounts tree
// @Summary Get COA Hierarchy Tree
// @Description Retrieve parent-child nested tree structure of all accounts
// @Tags Accounting - Accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse{data=[]AccountTreeResponse} "Account tree hierarchy retrieved"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/accounts/tree [get]
func (h *Handler) GetAccountTreeDoc(c *gin.Context) {
	h.GetAccountTree(c)
}

// FindAccountByIDDoc retrieves an account by ID
// @Summary Get Account by ID
// @Description Retrieve a single Chart of Accounts record by its UUID
// @Tags Accounting - Accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account UUID"
// @Success 200 {object} response.APIResponse{data=AccountResponse} "Account retrieved successfully"
// @Failure 404 {object} response.APIResponse "Account not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/accounts/{id} [get]
func (h *Handler) FindAccountByIDDoc(c *gin.Context) {
	h.FindAccountByID(c)
}

// UpdateAccountDoc updates an existing account
// @Summary Update Account
// @Description Update account attributes (name, description, active status)
// @Tags Accounting - Accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account UUID"
// @Param request body UpdateAccountRequest true "Update Payload"
// @Success 200 {object} response.APIResponse{data=AccountResponse} "Account updated successfully"
// @Failure 400 {object} response.APIResponse "Bad request"
// @Failure 404 {object} response.APIResponse "Account not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/accounts/{id} [put]
func (h *Handler) UpdateAccountDoc(c *gin.Context) {
	h.UpdateAccount(c)
}

// DeleteAccountDoc deletes an account
// @Summary Delete Account
// @Description Delete an account (non-system accounts with no children only)
// @Tags Accounting - Accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Account UUID"
// @Success 200 {object} response.APIResponse "Account deleted successfully"
// @Failure 400 {object} response.APIResponse "Bad request / System account"
// @Failure 404 {object} response.APIResponse "Account not found"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/accounts/{id} [delete]
func (h *Handler) DeleteAccountDoc(c *gin.Context) {
	h.DeleteAccount(c)
}

// SeedAccountsDoc seeds standard GAAP/IFRS Chart of Accounts
// @Summary Seed Standard Chart of Accounts
// @Description Pre-populates standard enterprise GAAP / IFRS accounts if empty
// @Tags Accounting - Accounts
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.APIResponse "Standard Chart of Accounts seeded successfully"
// @Failure 401 {object} response.APIResponse "Unauthorized"
// @Failure 500 {object} response.APIResponse "Internal server error"
// @Router /accounting/accounts/seed [post]
func (h *Handler) SeedAccountsDoc(c *gin.Context) {
	h.SeedAccounts(c)
}
