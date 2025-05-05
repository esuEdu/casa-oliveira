package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/esuEdu/casa-oliveira/internal/dto"
	"github.com/esuEdu/casa-oliveira/internal/service"
	"github.com/esuEdu/casa-oliveira/internal/util"
)

type AuthHandler struct {
	Service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler {
	return &AuthHandler{Service: s}
}

func (h *AuthHandler) SingUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		util.RespondWithJSON(w, http.StatusMethodNotAllowed, util.APIResponse{
			Message: "Method Not Allowed",
			Error:   "Invalid_HTTP_Method",
		})
	}

	var credential dto.SignupRequest

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Request body require",
			Error:   "Body_missing",
		})
		return
	}

	// Validate required fields
	if credential.Email == "" || credential.Password == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Missing or invalid fields",
			Error:   "invalid_product_data",
		})
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
		Message: "User SignUp successfully",
		Data:    credential,
	})

}

func (h *AuthHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		util.RespondWithJSON(w, http.StatusMethodNotAllowed, util.APIResponse{
			Message: "Method Not Allowed",
			Error:   "Invalid_HTTP_Method",
		})
	}

	var credential dto.SignupRequest

	err := json.NewDecoder(r.Body).Decode(&credential)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Request body require",
			Error:   "Body_missing",
		})
		return
	}

	// Validate required fields
	if credential.Email == "" || credential.Password == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Missing or invalid fields",
			Error:   "invalid_product_data",
		})
		return
	}

	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Can't register user",
			Error:   "",
		})
		return
	}

}
