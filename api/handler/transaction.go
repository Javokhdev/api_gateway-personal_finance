package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	pb "api-gateway/genproto"
)

// CreateTransaction handles creating a new transaction
// @Summary      Create Transaction
// @Description  Create a new transaction
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.CreateTransactionRequest true "Create transaction request"
// @Success      200 {object} pb.Response "Transaction created successfully"
// @Failure      500 {string} string "Error while creating transaction"
// @Router       /transaction/create [post]
func (h *Handler) CreateTransaction(ctx *gin.Context) {
	req := &pb.CreateTransactionRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	res, err := h.Transaction.CreateTransaction(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetTransactions handles retrieving a list of transactions
// @Summary      Get Transactions
// @Description  Retrieve a list of all transactions
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} pb.TransactionsResponse "List of transactions"
// @Failure      500 {string} string "Error while retrieving transactions"
// @Router       /transaction/list [get]
func (h *Handler) GetTransactions(ctx *gin.Context) {
	req := &pb.GetTransactionsRequest{}

	res, err := h.Transaction.GetTransactions(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// GetTransactionById handles retrieving a transaction by its ID
// @Summary      Get Transaction by ID
// @Description  Retrieve a specific transaction's details by its ID
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        transaction_id query string true "Transaction ID"
// @Success      200 {object} pb.TransactionResponse "Transaction details"
// @Failure      400 {string} string "Missing required query parameter: transaction_id"
// @Failure      500 {string} string "Error while fetching transaction"
// @Router       /transaction/get/{transaction_id} [get]
func (h *Handler) GetTransactionById(ctx *gin.Context) {
	transactionId := ctx.Param("transaction_id")
	if transactionId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: transaction_id"})
		return
	}

	req := &pb.GetTransactionByIdRequest{TransactionId: transactionId}

	res, err := h.Transaction.GetTransactionById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// UpdateTransaction handles updating an existing transaction
// @Summary      Update Transaction
// @Description  Update the details of an existing transaction
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body pb.UpdateTransactionRequest true "Update transaction request"
// @Success      200 {object} pb.Response "Transaction updated successfully"
// @Failure      400 {string} string "Invalid request or missing transaction_id"
// @Failure      500 {string} string "Error while updating transaction"
// @Router       /transaction/update [put]
func (h *Handler) UpdateTransaction(ctx *gin.Context) {
	req := &pb.UpdateTransactionRequest{}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if req.TransactionId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing transaction_id"})
		return
	}

	res, err := h.Transaction.UpdateTransaction(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}

// DeleteTransaction handles deleting a transaction by its ID
// @Summary      Delete Transaction
// @Description  Delete an existing transaction by its ID
// @Tags         Transaction
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        transaction_id query string true "Transaction ID"
// @Success      200 {object} pb.TransactionDeleteResponse "Transaction deleted successfully"
// @Failure      400 {string} string "Missing required query parameter: transaction_id"
// @Failure      500 {string} string "Error while deleting transaction"
// @Router       /transaction/delete/{transaction_id} [delete]
func (h *Handler) DeleteTransaction(ctx *gin.Context) {
	transactionId := ctx.Query("transaction_id")
	if transactionId == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Missing required query parameter: transaction_id"})
		return
	}

	req := &pb.DeleteTransactionRequest{TransactionId: transactionId}

	res, err := h.Transaction.DeleteTransaction(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, res)
}
