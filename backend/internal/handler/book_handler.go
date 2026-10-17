package handler

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/azrilpramudia/libre-go/internal/client/openlibrary"
	"github.com/azrilpramudia/libre-go/internal/domain/book"
	"github.com/azrilpramudia/libre-go/internal/repository/redis"
	"github.com/azrilpramudia/libre-go/internal/service"
	"github.com/azrilpramudia/libre-go/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BookHandler struct {
	svc      *service.BookService
	olClient *openlibrary.Client
	cache    *redis.Cache
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
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
		return
	}

	b, err := h.svc.CreateBook(c.Request.Context(), &book.Book{
		Title: req.Title,
		ISBN:  req.ISBN,
		Stock: req.Stock,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.Success(b))
}

func (h *BookHandler) List(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	books, err := h.svc.ListBooks(c.Request.Context(), int32(limit), int32(offset))
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.Success(books))
}

func (h *BookHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Error("invalid id"))
		return
	}
	b, err := h.svc.GetBook(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.Error("book not found"))
		return
	}
	c.JSON(http.StatusOK, response.Success(b))
}

func (h *BookHandler) SearchOpenLibrary(c *gin.Context) {
	title := c.Query("title")
	if title == "" {
		c.JSON(http.StatusBadRequest, response.Error("title query param required"))
		return
	}

	cacheKey := "ol:search:" + strings.ToLower(title)

	var cached openlibrary.SearchResult
	hit, err := h.cache.Get(c.Request.Context(), cacheKey, &cached)
	if err == nil && hit {
		c.JSON(http.StatusOK, response.Success(cached))
		return
	}

	result, err := h.olClient.SearchByTitle(c.Request.Context(), title)
	if err != nil {
		c.JSON(http.StatusBadGateway, response.Error(err.Error()))
		return
	}

	_ = h.cache.Set(c.Request.Context(), cacheKey, result, time.Hour)

	c.JSON(http.StatusOK, response.Success(result))
}

type importBookRequest struct {
	Title      string `json:"title" binding:"required"`
	AuthorName string `json:"author_name"`
	ISBN       string `json:"isbn"`
	CoverID    int    `json:"cover_i"`
	OLKey      string `json:"key" binding:"required"`
	Stock      int32  `json:"stock"`
}

func (h *BookHandler) ImportFromOpenLibrary(c *gin.Context) {
	var req importBookRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Error(err.Error()))
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
		c.JSON(http.StatusInternalServerError, response.Error(err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.Success(b))
}
