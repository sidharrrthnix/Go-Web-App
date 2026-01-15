package models

import "database/sql"

type Session struct {
	ID        int
	UserId    int
	Token     string
	TokenHash string
}

type SessionService struct {
	DB *sql.DB
}

func (ss *SessionService) Create(UserId int) (*Session, error) {}
