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
	SessionService *models.SessionService
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

	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Printf("Error creating session: %v", err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/users/me", http.StatusFound)
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
	session, err := u.SessionService.Create(user.ID)
	if err != nil {
		fmt.Printf("Error creating session: %v", err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    session.Token,
		Path:     "/",
		HttpOnly: true,
	}
	fmt.Fprintf(w, "Welcome back, %s!", user.Email)
	http.SetCookie(w, &cookie)
	http.Redirect(w, r, "/users/me", http.StatusFound)
}

func (u *Users) CurrentUser(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		log.Printf("Error getting session cookie: %v", err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	user, err := u.SessionService.User(cookie.Value)
	if err != nil {
		log.Printf("Error getting user from session: %v", err)
		http.Redirect(w, r, "/signin", http.StatusFound)
		return
	}
	fmt.Fprintf(w, "Current user: %s", user.Email)
}
