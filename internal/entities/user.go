package entities

import "time"

type User struct {
	Id       int       `db:"id"`
	Nome     string    `db:"nome"`
	Email    string    `db:"email"`
	Password string    `db:"password"`
	CriadoEm time.Time `db:"criado_em"`
}

type AuthUser struct {
	UserId   string
	UserName string
}
