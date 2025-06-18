package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/grazierShahid/high-school-management-system/internal/config"
	"github.com/grazierShahid/high-school-management-system/internal/services"
)

type UserHandler struct {
	// TODO add logger
	cfg         *config.Config
	userService *services.UserService
}

func NewUserHandler(cfg *config.Config, userService *services.UserService) *UserHandler { // TODO add logger as parameter and use
	return &UserHandler{
		cfg:         cfg,
		userService: userService,
	}
}

func (h *UserHandler) CheckHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusAccepted)
	json.NewDecoder(r.Body).Decode(map[string]string{
		"status": "success",
	})
}
