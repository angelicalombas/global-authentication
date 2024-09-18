package controllers

import (
	"encoding/json"
	"global-authentication/models"
	"global-authentication/services"
	"global-authentication/utils"
	"net/http"
	"sync"
	"time"

	_ "global-authentication/docs"

	"github.com/go-playground/validator/v10"
)

type UserController struct {
	userService services.UserService
	validate    *validator.Validate
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{
		userService: userService,
		validate:    validator.New(),
	}
}

var (
	loginAttempts = make(map[string]int)
	mu            sync.Mutex
	lockDuration  = 5 * time.Minute
)

// @Summary Register a new user
// @Description Register a new user with username and password
// @Tags users
// @Accept json
// @Produce json
// @Success 201 {object} map[string]string "User created successfully message"
// @Failure 400 {object} utils.ValidationErrorResponse
// @Failure 409 {object} utils.ValidationErrorResponse
// @Router /register [post]
// @Param request body models.User true "Request of Creating User"
func (c *UserController) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := decodeJSONBody(w, r, &user); err != nil {
		return
	}

	if err := c.validate.Struct(user); err != nil {
		response := utils.CreateValidationErrorResponse(err)
		w.WriteHeader(response.Code)
		json.NewEncoder(w).Encode(response)
		return
	}

	if err := c.userService.Register(&user); err != nil {
		utils.HandleError(w, err.Error(), http.StatusConflict, nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully."})
}

// @Summary Log in as user and generate token
// @Description Login with username and password
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {string} string "token"
// @Failure 400 {object} utils.ValidationErrorResponse
// @Failure 401 {object} utils.ValidationErrorResponse
// @Router /login [post]
// @Param request body models.User true "Request of Creating User"
func (c *UserController) Login(w http.ResponseWriter, r *http.Request) {
	var inputUser models.User
	if err := decodeJSONBody(w, r, &inputUser); err != nil {
		return
	}

	mu.Lock()
	attempts, exists := loginAttempts[inputUser.Username]
	if exists && attempts >= 3 {
		utils.HandleError(w, "Too many login attempts. Please try again later.", http.StatusTooManyRequests, nil)
		mu.Unlock()
		return
	}
	mu.Unlock()

	if err := c.validate.Struct(inputUser); err != nil {
		mu.Lock()
		loginAttempts[inputUser.Username]++
		go resetAttemptsAfterDelay(inputUser.Username)
		mu.Unlock()
		utils.HandleError(w, "Invalid credentials", http.StatusUnauthorized, nil)
		return
	}

	token, err := c.userService.Login(inputUser.Username, inputUser.Password)
	if err != nil {
		mu.Lock()
		loginAttempts[inputUser.Username]++
		go resetAttemptsAfterDelay(inputUser.Username)
		mu.Unlock()
		utils.HandleError(w, err.Error(), http.StatusUnauthorized, nil)
		return
	}

	mu.Lock()
	delete(loginAttempts, inputUser.Username)
	mu.Unlock()

	w.Write([]byte(token))
}

// @Summary Home endpoint
// @Description Welcome message for authenticated users
// @Tags home
// @Produce json
// @Success 200 {object} map[string]string "Welcome to the token-protected endpoint!"
// @Router /home [get]
// @Param Authorization header string true "Bearer token"
func (c *UserController) Home(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Welcome to the token-protected endpoint!"})
}

func decodeJSONBody(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		utils.HandleError(w, "Invalid request payload", http.StatusBadRequest, nil)
		return err
	}
	return nil
}

func resetAttemptsAfterDelay(username string) {
	time.Sleep(lockDuration)
	mu.Lock()
	delete(loginAttempts, username)
	mu.Unlock()
}
