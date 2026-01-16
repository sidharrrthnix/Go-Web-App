package routes

import (
	"web-app.com/simple/controllers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(static *controllers.Static, users *controllers.Users) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Static pages
	r.Get("/", static.Home)
	r.Get("/contact", static.Contact)
	r.Get("/faq", static.Faq)

	// User routes
	r.Get("/signup", users.Signup)
	r.Post("/users", users.CreateUser)
	r.Get("/signin", users.Signin)
	r.Post("/signin", users.ProcessSignin)
	r.Get("/users/me", users.CurrentUser)

	r.NotFound(controllers.NotFound)

	return r
}
