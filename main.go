package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"web-app.com/simple/controllers"
	"web-app.com/simple/routes"
	models "web-app.com/simple/sql"
	"web-app.com/simple/templates"
	"web-app.com/simple/views"

	"github.com/gorilla/csrf"
)

func main() {
	// Config
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Database
	db, err := models.Open(models.DefaultPostgresConfig())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
	log.Println("Connected to database")

	// Services
	userService := &models.UserService{DB: db}

	// CSRF protection
	csrfKey := []byte("32-byte-long-auth-key-here-now!!")
	csrfProtect := csrf.Protect(
		csrfKey,
		csrf.Secure(false),
		csrf.Path("/"),
		csrf.TrustedOrigins([]string{"localhost:8080"}),
	)

	// Controllers
	staticController := &controllers.Static{
		HomeTemplate:    views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml")),
		ContactTemplate: views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml")),
		FaqTemplate:     views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml")),
	}

	usersController := &controllers.Users{
		SignupTemplate: views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml")),
		SigninTemplate: views.Must(views.ParseFS(templates.FS, "signin.gohtml", "tailwind.gohtml")),
		UserService:    userService,
	}

	// Router with CSRF
	router := routes.NewRouter(staticController, usersController)
	protectedRouter := csrfProtect(router)

	// Server
	server := &http.Server{
		Addr:    ":" + port,
		Handler: protectedRouter,
	}

	go func() {
		log.Printf("Starting server on http://localhost:%s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %v", err)
	}
	log.Println("Server stopped")
}
