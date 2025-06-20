package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/grazierShahid/high-school-management-system/internal/config"
	"github.com/grazierShahid/high-school-management-system/internal/http/handlers"
	"github.com/grazierShahid/high-school-management-system/internal/http/routes"
	"github.com/grazierShahid/high-school-management-system/internal/services"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(fmt.Sprintf("failed to lead config: %v", err))
	}

	// logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	Level: slog.LevelDebug,
	// }))

	// TODO write logger

	dbConnStr := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DB.USERNAME,
		cfg.DB.PASSWORD,
		cfg.DB.HOST,
		cfg.DB.PORT,
		cfg.DB.NAME,
	)

	dbPool, err := pgxpool.New(context.Background(), dbConnStr)
	if err != nil {
		// TODO write logger
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	// TODO write logger : db connected

	userService := services.NewUserService(dbPool)

	// TODO write logger

	mux := http.NewServeMux()
	userHandler := handlers.NewUserHandler(cfg, userService) // TODO pass logger here
	router := routes.NewRouter(userHandler, mux)
	router.RegisterRoutes()

	http.ListenAndServe(fmt.Sprintf(":%s", cfg.App.AppPort), mux)

}
