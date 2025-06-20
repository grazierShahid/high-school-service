package handlers

import (
	"encoding/json"
	"fmt"
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
	h.userService.Ping(r.Context())

	w.WriteHeader(http.StatusAccepted)
	err := json.NewEncoder(w).Encode(map[string]string{
		"status": "success",
	})
	if err != nil {
		fmt.Printf("chill")
	}
}
