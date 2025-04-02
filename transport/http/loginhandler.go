package loginhttp

import (
	"encoding/json"
	"log"
	"net/http"

	"nexttalenta-backend/services"

	"github.com/go-playground/validator/v10"
)

type LoginHandler struct {
	loginService services.LoginService
}

// Struct for JSON request with validation
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Struct for JSON response
type LoginResponse struct {
	Username  string `json:"username"`
	IsSuccess bool   `json:"isSuccess"`
	Message   string `json:"message"`
}

func NewLoginHandler(loginService services.LoginService) *LoginHandler {
	return &LoginHandler{loginService: loginService}
}

func (h *LoginHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request LoginRequest
	validate := validator.New()

	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Error decoding request: %v", err)
		http.Error(w, `{"username": "", "isSuccess": false, "message": "Invalid JSON format"}`, http.StatusBadRequest)
		return
	}

	// Validate request data
	if err := validate.Struct(request); err != nil {
		log.Printf("Validation error: %v", err)
		http.Error(w, `{"username": "", "isSuccess": false, "message": "Username and password are required"}`, http.StatusBadRequest)
		return
	}

	log.Printf("Received Login Request: username=%s", request.Username)

	// Call usecase to authenticate user
	response, err := h.loginService.Authenticate(request.Username, request.Password)
	if err != nil {
		log.Printf("Authentication failed: %v", err)

		// Return custom JSON response for authentication failure
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized) // 401 Unauthorized
		json.NewEncoder(w).Encode(LoginResponse{
			Username:  request.Username,
			IsSuccess: false,
			Message:   "Invalid username or password",
		})
		return
	}

	// Successful login response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{
		Username:  request.Username,
		IsSuccess: true,
		Message:   response.Message,
	})
}
