package handler

import (
	"errors"
	"net/http"

	"github.com/azrilpramudia/libre-go/internal/domain/loan"
	"github.com/azrilpramudia/libre-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoanHandler struct {
	svc *service.LoanService
}

func NewLoanHandler(svc *service.LoanService) *LoanHandler {
	return &LoanHandler{svc: svc}
}

type borrowRequest struct {
	BookID   string `json:"book_id" binding:"required"`
	MemberID string `json:"member_id" binding:"required"`
}

func (h *LoanHandler) Borrow(c *gin.Context) {
	var req borrowRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	bookID, err := uuid.Parse(req.BookID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid book_id"})
		return
	}
	memberID, err := uuid.Parse(req.MemberID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid member_id"})
		return
	}

	l, err := h.svc.BorrowBook(c.Request.Context(), bookID, memberID)
	if errors.Is(err, loan.ErrBookNotAvailable) {
		c.JSON(http.StatusConflict, gin.H{"error": "book not available"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, l)
}

func (h *LoanHandler) Return(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	l, err := h.svc.ReturnBook(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, l)
}

func (h *LoanHandler) ListByMember(c *gin.Context) {
	memberID, err := uuid.Parse(c.Param("memberId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid member id"})
		return
	}

	loans, err := h.svc.ListMemberLoans(c.Request.Context(), memberID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, loans)
}
