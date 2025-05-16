package auth

import (
	"encoding/json"
	"net/http"

	"github.com/sudarshanmg/gotask/pkg/response"
)

type Handler struct {
	service AuthService
}

func NewHandler(service AuthService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	user, err := h.service.Register(req)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusCreated, user)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	accesstoken, refreshToken, err := h.service.Login(req)
	if err != nil {
		response.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]string{
		"access_token":  accesstoken,
		"refresh_token": refreshToken,
	})
}

func (h *Handler) Refresh(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req RefreshTokenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	newToken, err := h.service.Refresh(req.RefreshToken)
	if err != nil {
		response.WriteError(w, http.StatusUnauthorized, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]string{
		"access_token": newToken,
	})
}

func (h *Handler) Logout(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req RefreshTokenRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		response.WriteError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	err := h.service.Logout(req.RefreshToken)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "logged out successfully",
	})
}

func (h *Handler) LogoutAll(w http.ResponseWriter, r *http.Request) {
	userID := GetUserID(r)
	if userID == 0 {
		response.WriteError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	err := h.service.LogoutAll(userID)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "failed to log logout")
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]string{
		"message": "logged out successfully",
	})
}
