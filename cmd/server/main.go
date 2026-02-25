package main

import (
	"log"
	"net/http"

	"robot-server/internal/auth"
	"robot-server/internal/database"
	"robot-server/internal/middleware"
	"robot-server/internal/support"
)

func main() {

	db, err := database.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	repo := &auth.Repository{DB: db}
	service := &auth.Service{Repo: repo}
	handler := &auth.Handler{Service: service}

	supportService := &support.Service{}
	supportHandler := &support.Handler{Service: supportService}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/auth/register", handler.Register)
	mux.HandleFunc("/api/auth/login", handler.Login)
	mux.HandleFunc("/api/auth/forgot-password", handler.ForgotPassword)
	mux.Handle("/api/auth/logout",
		middleware.AuthMiddleware(http.HandlerFunc(handler.Logout)))
	
	mux.Handle("/api/support/send_email",
		middleware.AuthMiddleware(
			http.HandlerFunc(supportHandler.SendEmail),
		))

	log.Println("Server started :8080")
	http.ListenAndServe(":8080", mux)
}
