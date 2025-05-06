package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/esuEdu/casa-oliveira/internal/dto"
	"github.com/esuEdu/casa-oliveira/internal/service"
	"github.com/esuEdu/casa-oliveira/internal/util"
)

type UserHandler struct {
	Service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.RespondWithJSON(w, http.StatusMethodNotAllowed, util.APIResponse{
			Message: "Method Not Allowed",
			Error:   "Invalid_HTTP_Method",
		})
	}

	var input struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Request body require",
			Error:   "Body_missing",
		})
		return
	}

	if input.Name == "" || input.Email == "" || input.Phone == "" || input.Password == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Missing or invalid fields",
			Error:   "invalid_product_data",
		})
		return
	}

	user := dto.UserDTO{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}

	userCreated, err := h.Service.SignUp(&user, input.Password)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Failed to SignUp",
			Error:   err.Error(),
		})
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
		Message: "User signup successfully",
		Data:    userCreated,
	})
}

func (h *UserHandler) SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		util.RespondWithJSON(w, http.StatusMethodNotAllowed, util.APIResponse{
			Message: "Method Not Allowed",
			Error:   "Invalid_HTTP_Method",
		})
	}

	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Request body require",
			Error:   "Body_missing",
		})
		return
	}

	if input.Email == "" || input.Password == "" {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Missing or invalid fields",
			Error:   "invalid_product_data",
		})
		return
	}

	auth, err := h.Service.SignIn(input.Email, input.Password)
	if err != nil {
		util.RespondWithJSON(w, http.StatusBadRequest, util.APIResponse{
			Message: "Failed to SignIn",
			Error:   err.Error(),
		})
		return
	}

	util.RespondWithJSON(w, http.StatusCreated, util.APIResponse{
		Message: "User signin successfully",
		Data:    auth,
	})
}
