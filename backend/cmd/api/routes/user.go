package routes

import (
	"backend/internal/modules/user/handler"
	"backend/internal/modules/user/repository"
	"backend/internal/modules/user/services"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(db *sql.DB, router *gin.Engine) {
	// Inicializar las dependencias
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserController(userService)

	// Grupo de rutas para /users
	userGroup := router.Group("/users")
	{
		userGroup.POST("", userHandler.CreateUser)
		userGroup.GET("", userHandler.GetUsers)
	}
}
