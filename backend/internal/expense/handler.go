package expense

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

type createRequest struct {
	Amount      int64  `json:"amount"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func (h *Handler) CreateExpense(c *gin.Context) {
	id := c.GetHeader("Idempotency-Key")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Idempotency-Key required"})
		return
	}

	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	if req.Amount <= 0 || req.Date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid amount or date"})
		return
	}

	exp := &Expense{
		ID:          id,
		Amount:      req.Amount,
		Category:    req.Category,
		Description: req.Description,
		Date:        req.Date,
	}

	_ = h.repo.Create(exp)

	existing, err := h.repo.GetByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch expense"})
		return
	}

	c.JSON(http.StatusOK, existing)
}

func (h *Handler) ListExpenses(c *gin.Context) {
	category := c.Query("category")

	expenses, err := h.repo.List(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list expenses"})
		return
	}

	c.JSON(http.StatusOK, expenses)
}
