package users

import (
	"errors"
	appjwt "fish-hunter/util/jwt"
)

type UserUseCase struct {
	UserRepository Repository
}

func NewUserUseCase(userRepository Repository) UseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Register(domain *Domain) (Domain, error) {
	return u.UserRepository.Register(domain)
}

func (u *UserUseCase) Login(domain *Domain) (string, error) {
	user, err := u.UserRepository.Login(domain)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	token, err := appjwt.GenerateToken(user.ID, user.Roles)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) GetProfile(id string) (Domain, error) {
	return u.UserRepository.GetProfile(id)
}

func (u *UserUseCase) UpdateProfile(domain *Domain) (Domain, error) {
	return u.UserRepository.UpdateProfile(domain)
}

func (u *UserUseCase) UpdatePassword(domain *Domain) (Domain, error) {
	return u.UserRepository.UpdatePassword(domain)
}

func (u *UserUseCase) GetAllUsers() ([]Domain, error) {
	return u.UserRepository.GetAllUsers()
}

func (u *UserUseCase) GetByID(id string) (Domain, error) {
	return u.UserRepository.GetByID(id)
}

func (u *UserUseCase) Update(domain *Domain) (Domain, error) {
	return u.UserRepository.Update(domain)
}

func (u *UserUseCase) Delete(id string) (Domain, error) {
	return u.UserRepository.Delete(id)
}

func (u *UserUseCase) Logout(token string) error {
	appjwt.RemoveToken(token)
	return nil
}