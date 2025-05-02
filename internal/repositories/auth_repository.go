package repositories

import (
	"errors"
	"fmt"

	"github.com/ViniciusTei/viniciustei-blog/internal/database"
	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
	"github.com/ViniciusTei/viniciustei-blog/utils"
)

type AuthRepositoryImpl struct {
	Db *database.DatabaseImpl
}

func (r *AuthRepositoryImpl) SignIn(username, password string) (string, error) {
	var user entities.User
	err := r.Db.SelectOne("SELECT * FROM users WHERE username = $1 AND password = $2", user)
	if err != nil {
		//TODO: handle SQL error and return a more user-friendly error
		return "", err
	}

	if user.Email == username && user.Password == password {
		key := []byte("6091835705053067")
		token, err := utils.Encrypt(key, fmt.Sprintf("%d:%s", user.Id, user.Email))
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", errors.New("invalid credentials")
}
