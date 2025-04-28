package usecases

type AuthRepository interface {
	SignIn(username, password string) (string, error)
}

type AuthUseCase struct {
	Repo AuthRepository
}

func (uc *AuthUseCase) SignIn(username, password string) (string, error) {
	return uc.Repo.SignIn(username, password)
}
