package users

type UserUseCase struct {
	UserRepository Repository
}

func NewUserUseCase(userRepository Repository) UseCase {
	return &UserUseCase{
		UserRepository: userRepository,
	}
}

func (u *UserUseCase) Register(domain *Domain) Domain {
	return u.UserRepository.Register(domain)
}

func (u *UserUseCase) Login(domain *Domain) (Domain, error) {
	return u.UserRepository.Login(domain)
}