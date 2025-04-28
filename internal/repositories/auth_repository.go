package repositories

import (
	"errors"
	"fmt"

	"github.com/ViniciusTei/viniciustei-blog/internal/entities"
	"github.com/ViniciusTei/viniciustei-blog/utils"
)

type AuthRepositoryImpl struct{}

func (r *AuthRepositoryImpl) SignIn(username, password string) (string, error) {
	if username == "viniciust" && password == "123456" {
		user := entities.User{
			UserId:   "viniciust",
			UserName: "viniciust",
		}

		key := []byte("6091835705053067")
		token, err := utils.Encrypt(key, fmt.Sprintf("%s:%s", user.UserId, user.UserName))
		if err != nil {
			return "", err
		}

		return token, nil
	}

	return "", errors.New("invalid credentials")
}
