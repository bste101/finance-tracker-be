package router

import (
	"net/http"

	"github.com/bste101/finance-tracker/db/sqlc"
	"github.com/bste101/finance-tracker/internal/auth"
	"github.com/bste101/finance-tracker/internal/config"
	"github.com/bste101/finance-tracker/pkg/response"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	r *gin.Engine,
	cfg *config.Config,
	queries *sqlc.Queries,
) {

	// =========================
	// HEALTH
	// =========================

	r.GET("/health", health)

	// =========================
	// AUTH
	// =========================

	authService := auth.NewService(queries)
	authHandler := auth.NewHandler(authService)

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	// =========================
	// PROTECTED ROUTES
	// =========================

	api := r.Group("/api")
	api.Use(auth.Middleware())

	api.GET("/me", func(c *gin.Context) {
		userID, _ := c.Get("userID")

		response.OK(c, gin.H{
			"userID": userID,
		})
	})

	// =========================
	// 404
	// =========================

	r.NoRoute(func(c *gin.Context) {
		response.Error(
			c,
			http.StatusNotFound,
			"route not found",
			"ROUTE_NOT_FOUND",
		)
	})
}

func health(c *gin.Context) {
	response.OK(c, gin.H{
		"status": "ok",
	})
}