package server

import (
	"fantasy-bball-schedule-grid/internal/config"
	"fantasy-bball-schedule-grid/internal/handlers"
	"log"
	"net/http"
)

func Run() error {
	cfg := config.Load()
	mux := setupRoutes()
	handler := setupMiddleware(mux, cfg)
	log.Printf("Server starting on :%s (env: %s)", cfg.Port, cfg.Environment)

	return http.ListenAndServe(":"+cfg.Port, handler)
}

func setupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/schedule", handlers.GetWeeklyScheduleHandler)
	mux.HandleFunc("/api/gameDates", handlers.GetGameDatesHandler)
	mux.HandleFunc("/api/fantasyTeamRosters", handlers.GetFantasyTeamRostersHandler)
	mux.HandleFunc("/api/freeAgents", handlers.GetFreeAgentsHandler)

	return mux
}

func setupMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return corsMiddleware(loggingMiddleware(next), cfg)
}

func corsMiddleware(next http.Handler, cfg *config.Config) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		for _, allowedOrigin := range cfg.CORSOrigins {
			if origin == allowedOrigin || allowedOrigin == "*" {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}

		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
