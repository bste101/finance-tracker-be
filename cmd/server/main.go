package main

import (
	"log"

	"github.com/bste101/finance-tracker/db/sqlc"
	"github.com/bste101/finance-tracker/internal/config"
	"github.com/bste101/finance-tracker/internal/database"
	"github.com/bste101/finance-tracker/internal/router"
	"github.com/gin-gonic/gin"
)

func main() {

	cfg := config.Load()

	dbpool, err := database.New()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	defer dbpool.Close()

	queries := sqlc.New(dbpool)

	r := gin.Default()

	router.SetupRoutes(r, cfg, queries)

	if err := r.Run(":" + cfg.AppPort); err != nil {
		log.Fatal(err)
	}
}
