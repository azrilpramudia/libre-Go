package main

import (
	"context"
	"log"
	"os"

	"github.com/azrilpramudia/libre-go/internal/client/openlibrary"
	"github.com/azrilpramudia/libre-go/internal/handler"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres/sqlc"
	"github.com/azrilpramudia/libre-go/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer pool.Close()

	queries := sqlc.New(pool)
	bookRepo := postgres.NewBookRepository(queries)
	bookService := service.NewBookRepository(bookRepo)
	olClient := openlibrary.NewClient(os.Getenv("OPENLIBRARY_BASE_URL"))
	bookHandler := handler.NewBookHandler(bookService, olClient)

	r := gin.Default()
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	books := r.Group("/books")
	{
		books.POST("", bookHandler.Create)
		books.GET("", bookHandler.List)
		books.GET("/:id", bookHandler.GetByID)
		books.GET("/search/openlibrary", bookHandler.SearchOpenLibrary)
	}

	log.Println("server running on :5000")
	r.Run(":5000")
}
