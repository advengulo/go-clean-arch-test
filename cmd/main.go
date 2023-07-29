package main

import (
	"github.com/advengulo/go-clean-arch-test/internal/middlewares"
	_authHandler "github.com/advengulo/go-clean-arch-test/internal/modules/auth/handler/http"
	_authUsecase "github.com/advengulo/go-clean-arch-test/internal/modules/auth/usecase"
	_userHandler "github.com/advengulo/go-clean-arch-test/internal/modules/user/handler/http"
	_userRepository "github.com/advengulo/go-clean-arch-test/internal/modules/user/repository"
	_userUsecase "github.com/advengulo/go-clean-arch-test/internal/modules/user/usecase"
	"github.com/advengulo/go-clean-arch-test/pkg/database"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/subosito/gotenv"
	"gorm.io/gorm"
	"log"
)

var (
	db  *gorm.DB
	err error
)

func main() {

	// Load the .env file
	err = gotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	// Open a connection to the database
	db, err = database.InitDB()
	if err != nil {
		log.Fatalf("Error connection to database: %s", err.Error())
	}

	// Run Migration
	err = database.MigrateDB(db)
	if err != nil {
		log.Fatalf("Error run migration: %s", err.Error())
	}

	// Init validation
	validate := validator.New()

	e := echo.New()

	// Execute Middleware
	middlewares.Register(e)

	// Create the Repository dependencies
	userRepository := _userRepository.NewUserRepository(db)

	// Create the UseCase dependencies
	userUseCase := _userUsecase.NewUserUseCase(userRepository)
	authUseCase := _authUsecase.NewAuthUseCase(userUseCase)

	// Register the UserHandler with the Echo router
	_userHandler.NewUserHandler(e, validate, userUseCase)
	_authHandler.NewAuthHandler(e, validate, authUseCase)

	// Start the server
	e.Logger.Fatal(e.Start(":8081"))
}
