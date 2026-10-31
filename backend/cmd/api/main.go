package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/azrilpramudia/libre-go/config"
	"github.com/azrilpramudia/libre-go/internal/client/openlibrary"
	"github.com/azrilpramudia/libre-go/internal/handler"
	"github.com/azrilpramudia/libre-go/internal/middleware"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres"
	"github.com/azrilpramudia/libre-go/internal/repository/postgres/sqlc"
	"github.com/azrilpramudia/libre-go/internal/repository/redis"
	"github.com/azrilpramudia/libre-go/internal/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	databaseURL, err := config.DatabaseURL()
	if err != nil {
		log.Fatalf("failed to configure database connection: %v", err)
	}

	pool, err := pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	defer pool.Close()

	queries := sqlc.New(pool)
	bookRepo := postgres.NewBookRepository(queries)
	authorRepo := postgres.NewAuthorRepository(queries)
	bookService := service.NewBookService(bookRepo, authorRepo)
	olClient := openlibrary.NewClient(os.Getenv("OPENLIBRARY_BASE_URL"))
	cache := redis.NewCache(os.Getenv("REDIS_ADDR"))
	bookHandler := handler.NewBookHandler(bookService, olClient, cache)
	memberRepo := postgres.NewMemberRepository(queries)
	memberService := service.NewMemberService(memberRepo)
	memberHandler := handler.NewMemberHandler(memberService)
	loanRepo := postgres.NewLoanRepository(pool, queries)
	loanService := service.NewLoanService(loanRepo)
	loanHandler := handler.NewLoanHandler(loanService)
	authService := service.NewAuthService(memberRepo)
	authHandler := handler.NewAuthHandler(authService)

	r := gin.New()
	r.SetTrustedProxies(nil)

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.Timeout(5 * time.Second))
	r.GET("/health", func(c *gin.Context) { c.JSON(200, gin.H{"status": "ok"}) })

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	books := r.Group("/books")
	{
		books.POST("", bookHandler.Create)
		books.GET("", bookHandler.List)
		books.GET("/:id", bookHandler.GetByID)
		books.GET("/search/openlibrary", bookHandler.SearchOpenLibrary)
		books.POST("/import/openlibrary", bookHandler.ImportFromOpenLibrary)
	}

	members := r.Group("/members")
	{
		members.POST("", memberHandler.Create)
		members.GET("", memberHandler.List)
		members.GET("/:id", memberHandler.GetByID)
	}

	loans := r.Group("/loans")
	loans.Use(middleware.AuthRequired())
	{
		loans.POST("/borrow", loanHandler.Borrow)
		loans.PATCH("/:id/return", loanHandler.Return)
		loans.GET("/member/:memberId", loanHandler.ListByMember)
	}

	log.Println("server running on :5000")
	r.Run(":5000")
}
