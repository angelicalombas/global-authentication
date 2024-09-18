package controllers

import "net/http"

type UserControllerInterface interface {
	Register(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	Home(w http.ResponseWriter, r *http.Request)
}
