package server

import (
	"backend/internal/handler"
	"backend/internal/infrastructure/database"
	"backend/internal/router"
	"backend/internal/service"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Server struct {
	Router *echo.Echo
}

func Inject(db *sqlx.DB) *Server {
	userRepo := database.NewUserRepository(db)

	userService := service.NewUserService(userRepo)

	userHandler := handler.NewUserHandler(userService)

	echoRouter := router.NewRouter(userHandler)

	return &Server{
		Router: echoRouter,
	}
}
