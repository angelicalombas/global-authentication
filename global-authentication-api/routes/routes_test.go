package routes

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"global-authentication/utils/mocks"

	"github.com/gorilla/mux"
)

func TestInitializeRoutes(t *testing.T) {
	mockUserController := &mocks.MockUserController{
		RegisterFunc: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully."})
		},
		LoginFunc: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("token"))
		},
		HomeFunc: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the home!"})
		},
	}

	router := mux.NewRouter()
	homeRouter := router.PathPrefix("/home").Subrouter()
	homeRouter.Use(mocks.TestAuthMiddleware)
	homeRouter.HandleFunc("", mockUserController.Home).Methods("GET")

	InitializeRoutes(router, mockUserController)

	t.Run("POST /register", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/register", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code for /register: got %v want %v", status, http.StatusCreated)
		}
	})

	t.Run("POST /login", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/login", nil)
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code for /login: got %v want %v", status, http.StatusOK)
		}
	})

	t.Run("GET /home", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/home", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer token")

		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)

		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code for /home: got %v want %v", status, http.StatusOK)
		}
	})
}
