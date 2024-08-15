package handler

import (
	"net/http"

	pb "api-gateway/genproto"

	"github.com/gin-gonic/gin"
)

// CreateAccount handles the creation of a new account
// @Summary      Create Account
// @Description  Create a new account for a user
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.CreateAccountRequest true "Create account request"
// @Success      200 {object} pb.CreateAccountRes "Account created successfully"
// @Failure      500 {string} string "Error while creating account"
// @Router       /account/create [post]
func (h *Handler) CreateAccount(ctx *gin.Context) {
	req := &pb.CreateAccountRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Account.CreateAccount(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// ListAccounts handles listing all accounts
// @Summary      List Accounts
// @Description  List all accounts for a user
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.ListAccountsResponse "List of accounts"
// @Failure      500 {string} string "Error while listing accounts"
// @Router       /account/list [get]
func (h *Handler) ListAccounts(ctx *gin.Context) {
	req := &pb.ListAccountsRequest{}

	res, err := h.Account.ListAccounts(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetAccountById handles retrieving a specific account by ID
// @Summary      Get Account by Id
// @Description  Retrieve an account's details by its ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        account_id query string true "Account Id"
// @Success      200 {object} pb.AccountResponse "Account details"
// @Failure      400 {string} string "Missing required query parameter: account_id"
// @Failure      500 {string} string "Error while fetching account"
// @Router       /account/get/{account_id} [get]
func (h *Handler) GetAccountById(ctx *gin.Context) {
	accountID := ctx.Query("account_id")
	if accountID == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: account_id"})
		return
	}

	req := &pb.GetAccountByIdRequest{AccountId: accountID}

	res, err := h.Account.GetAccountById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateAccount handles updating an existing account
// @Summary      Update Account
// @Description  Update the details of an existing account
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateAccountRequest true "Update account request"
// @Success      200 {object} pb.CreateAccountRes "Account updated successfully"
// @Failure      400 {string} string "Invalid request or missing account_id"
// @Failure      500 {string} string "Error while updating account"
// @Router       /account/update [put]
func (h *Handler) UpdateAccount(ctx *gin.Context) {
	req := &pb.UpdateAccountRequest{}
	err := ctx.BindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.AccountId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing account_id"})
		return
	}

	res, err := h.Account.UpdateAccount(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteAccount handles deleting an account by ID
// @Summary      Delete Account
// @Description  Delete an existing account by its ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        account_id query string true "Account ID"
// @Success      200 {object} pb.DeleteResponse "Account deleted successfully"
// @Failure      400 {string} string "Missing required query parameter: account_id"
// @Failure      500 {string} string "Error while deleting account"
// @Router       /account/delete/{account_id} [delete]
func (h *Handler) DeleteAccount(ctx *gin.Context) {
	accountId := ctx.Query("account_id")
	if accountId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: account_id"})
		return
	}

	req := &pb.DeleteAccountRequest{AccountId: accountId}

	res, err := h.Account.DeleteAccount(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
