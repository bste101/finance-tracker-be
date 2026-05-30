package router

import (
	"net/http"

	"github.com/bste101/finance-tracker/db/sqlc"
	"github.com/bste101/finance-tracker/internal/auth"
	"github.com/bste101/finance-tracker/internal/config"
	"github.com/bste101/finance-tracker/internal/user"
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
	authRepo := auth.NewRepository(queries)
	authService := auth.NewService(authRepo)
	authHandler := auth.NewHandler(authService)

	r.POST("/auth/register", authHandler.Register)
	r.POST("/auth/login", authHandler.Login)

	// =========================
	// PROTECTED ROUTES
	// =========================

	api := r.Group("")
	api.Use(auth.Middleware())

	// USER
	userRepo := user.NewRepository(queries)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	api.GET("/me", userHandler.Me)
	api.PATCH("/me", userHandler.Update)
	api.DELETE("/me", userHandler.Delete)
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
