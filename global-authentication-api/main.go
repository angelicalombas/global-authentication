package main

import (
	"log"
	"net/http"

	"global-authentication/config"
	"global-authentication/controllers"
	"global-authentication/models"
	"global-authentication/repositories"
	"global-authentication/routes"
	"global-authentication/services"

	_ "global-authentication/docs"

	"github.com/gorilla/mux"
	"github.com/swaggo/http-swagger"
)

func main() {
	config.Connect()
	config.DB.AutoMigrate(&models.User{})

	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	router := mux.NewRouter()
	routes.InitializeRoutes(router, userController)

	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server running on port :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
