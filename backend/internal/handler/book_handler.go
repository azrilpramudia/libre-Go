package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/azrilpramudia/libre-go/internal/client/openlibrary"
	"github.com/azrilpramudia/libre-go/internal/domain/book"
	"github.com/azrilpramudia/libre-go/internal/repository/redis"
	"github.com/azrilpramudia/libre-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandler struct {
	svc      *service.BookService
	olClient *openlibrary.Client
	cache    *redis.Cache
}

type importBookRequest struct {
	Title      string `json:"title" binding:"required"`
	AuthorName string `json:"author_name"`
	ISBN       string `json:"isbn"`
	CoverID    int    `json:"cover_i"`
	OLKey      string `json:"key" binding:"required"`
	Stock      int32  `json:"stock"`
}

func NewBookHandler(svc *service.BookService, olClient *openlibrary.Client, cache *redis.Cache) *BookHandler {
	return &BookHandler{svc: svc, olClient: olClient, cache: cache}
}

type createBookRequest struct {
	Title string `json:"title" binding:"required"`
	ISBN  string `json:"isbn"`
	Stock int32  `json:"stock"`
}

func (h *BookHandler) Create(c *gin.Context) {
	var req createBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	b, err := h.svc.CreateBook(c.Request.Context(), &book.Book{
		Title: req.Title,
		ISBN:  req.ISBN,
		Stock: req.Stock,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, b)
}

func (h *BookHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	books, err := h.svc.ListBooks(c.Request.Context(), int32(limit), int32(offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	b, err := h.svc.GetBook(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "book not found"})
		return
	}
	c.JSON(http.StatusOK, b)
}

func (h *BookHandler) SearchOpenLibrary(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "title query param required"})
		return
	}

	cacheKey := "ol:search:" + title

	var cached openlibrary.SearchResult
	hit, err := h.cache.Get(c.Request.Context(), cacheKey, &cached)
	if err == nil && hit {
		c.JSON(http.StatusOK, cached)
		return
	}

	result, err := h.olClient.SearchByTitle(c.Request.Context(), title)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
		return
	}

	_ = h.cache.Set(c.Request.Context(), cacheKey, result, time.Hour)

	c.JSON(http.StatusOK, result)
}

func (h *BookHandler) ImportFromOpenLibrary(c *gin.Context) {
	var req importBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if req.Stock < 1 {
		req.Stock = 1
	}

	b, err := h.svc.ImportBook(c.Request.Context(), service.ImportBookInput{
		Title:         req.Title,
		AuthorName:    req.AuthorName,
		ISBN:          req.ISBN,
		CoverURL:      openlibrary.CoverURL(req.CoverID),
		OpenLibraryID: req.OLKey,
		Stock:         req.Stock,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, b)
}
