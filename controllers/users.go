package controllers

import (
	"fmt"
	"log"
	"net/http"

	models "web-app.com/simple/sql"
	"web-app.com/simple/views"
)

type Users struct {
	SignupTemplate views.Template
	SigninTemplate views.Template
	UserService    *models.UserService
}

func (u *Users) Signup(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.SignupTemplate.Execute(w, r, data)
}

func (u *Users) CreateUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserService.Create(email, password)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	log.Printf("Created user: %s (id: %d)", user.Email, user.ID)
	fmt.Fprintf(w, "User created! Email: %s, ID: %d", user.Email, user.ID)
}

func (u *Users) Signin(w http.ResponseWriter, r *http.Request) {
	var data struct{ Email string }
	data.Email = r.FormValue("email")
	u.SigninTemplate.Execute(w, r, data)
}

func (u *Users) ProcessSignin(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")

	user, err := u.UserService.Authenticate(email, password)
	if err != nil {
		log.Printf("Auth error: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome back, %s!", user.Email)
}
