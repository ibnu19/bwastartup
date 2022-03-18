package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"bwastartup/user"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func TransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

// Get list of transaction by campaign
func (h *transactionHandler) GetCampaignTransaction(c *gin.Context) {
	campaignId, _ := strconv.Atoi(c.Param("id"))

	// Get logged in user
	user := c.MustGet("currentUser").(user.User)

	transactions, err := h.service.GetTransactionsByCampaignId(campaignId, user)
	if err != nil {
		response := helper.APIResponse("Failed to get campaign's transactions", http.StatusBadRequest, "error", transaction.FormatCampaignTransactions(transactions))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of campaign's transactions", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

// Get list of transaction by user
func (h *transactionHandler) GetUserTransactions(c *gin.Context) {

	// Get logged in user
	user := c.MustGet("currentUser").(user.User)

	// Get transactions of current logged in user
	transactions, err := h.service.GetTransactionByUserId(user.ID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transactions", http.StatusBadRequest, "error", transactions)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("List of user's transactions", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

// Create a transaction
func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	var input transaction.CreateTransactionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	user := c.MustGet("currentUser").(user.User)
	input.User = user

	createTransaction, err := h.service.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Succes to create transaction", http.StatusOK, "success", transaction.FormatTransaction(createTransaction))
	c.JSON(http.StatusOK, response)
}
