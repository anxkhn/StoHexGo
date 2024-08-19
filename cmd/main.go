package main

import (
	"fmt"
	"log"
	"net/http"
	"stock-portfolio/internal/adapters/handlers"
	"stock-portfolio/internal/adapters/repositories"
	"stock-portfolio/internal/config"
	"stock-portfolio/internal/core/domain"
	"stock-portfolio/internal/core/services"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()

	db, err := gorm.Open(sqlite.Open(cfg.DatabasePath), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&domain.Transaction{})

	repo := repositories.NewSQLiteRepository(db)
	portfolioService := services.NewPortfolioService(repo)
	handler := handlers.NewHTTPHandler(portfolioService)

	http.HandleFunc("/", handler.HandleRoot)
	http.HandleFunc("/transactions", handler.HandleTransactions)
	http.HandleFunc("/balance", handler.HandleBalance)

	log.Printf("Server starting on http://localhost:%d", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", cfg.ServerPort), nil))
}
