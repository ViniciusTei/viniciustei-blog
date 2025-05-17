package repositories

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/ViniciusTei/viniciustei-blog/internal/database"
	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
	"github.com/ViniciusTei/viniciustei-blog/utils"
)

type AuthRepositoryImpl struct {
	Db *database.DatabaseImpl
}

func (r *AuthRepositoryImpl) SignIn(username, password string) (string, error) {
	var user entities.User

	log.Printf("Trying connect user: %v\n", username)
	err := r.
		Db.
		Conn.
		QueryRow(context.Background(), "SELECT * FROM usuarios WHERE email = $1", username).
		Scan(&user.Id, &user.Nome, &user.Email, &user.Password, &user.CriadoEm)
	if err != nil {
		//TODO: handle SQL error and return a more user-friendly error
		log.Printf("Database query error: %v\n", err)
		return "", err
	}

	log.Printf("User found! Validate user pass: %v\n", user)
	if user.Email == username && utils.CompareHashText(password, user.Password) {
		//TODO: use a more secure key management system
		key := []byte("6091835705053067")
		token, err := utils.Encrypt(key, fmt.Sprintf("%d:%s", user.Id, user.Email))
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", errors.New("invalid credentials")
}
