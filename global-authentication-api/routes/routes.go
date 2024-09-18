package routes

import (
	"global-authentication/controllers"
	"global-authentication/utils"

	"github.com/gorilla/mux"
)

// func InitializeRoutes(router *mux.Router, userController *controllers.UserController) {
// 	router.HandleFunc("/register", userController.Register).Methods("POST")
// 	router.HandleFunc("/login", userController.Login).Methods("POST")

// 	homeRouter := router.PathPrefix("/home").Subrouter()
// 	homeRouter.Use(utils.AuthMiddleware)
// 	homeRouter.HandleFunc("", userController.Home).Methods("GET")

// }

func InitializeRoutes(router *mux.Router, userController controllers.UserControllerInterface) {
	router.HandleFunc("/register", userController.Register).Methods("POST")
	router.HandleFunc("/login", userController.Login).Methods("POST")

	homeRouter := router.PathPrefix("/home").Subrouter()
	homeRouter.Use(utils.AuthMiddleware)
	homeRouter.HandleFunc("", userController.Home).Methods("GET")
}
