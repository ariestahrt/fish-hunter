package users

import (
	"errors"
	appjwt "fish-hunter/util/jwt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
	// Check if email already exists
	_, err := u.UserRepository.GetByEmail(domain.Email)
	if err == nil {
		return Domain{}, errors.New("email already exists")
	}

	// Check if username already exists
	_, err = u.UserRepository.GetByUsername(domain.Username)
	if err == nil {
		return Domain{}, errors.New("username already exists")
	}

	// Encrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(domain.Password), bcrypt.MinCost)
	domain.Password = string(password)
	domain.IsActive = false

	domain.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	domain.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())
	domain.DeletedAt = primitive.NewDateTimeFromTime(primitive.NilObjectID.Timestamp())

	// Force change role to user
	domain.Roles = []string{"guest"}

	// Create user
	user, err := u.UserRepository.Create(domain)
	return user, err
}

func (u *UserUseCase) Login(domain *Domain) (string, error) {
	// Get user by email
	user, err := u.UserRepository.GetByEmail(domain.Email)
	if err != nil {
		return "", err
	}

	// Check is user active
	if !user.IsActive {
		return "", errors.New("user is not active")
	}

	// Check password bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(domain.Password))
	if err != nil {
		return "", err
	}

	// Generate JWT
	token, err := appjwt.GenerateToken(user.Id.Hex(), user.Roles)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) GetProfile(id string) (Domain, error) {
	// return u.UserRepository.GetProfile(id)
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return u.UserRepository.GetByID(ObjId)
}

func (u *UserUseCase) UpdateProfile(old *Domain, new *Domain) (Domain, error) {
	// Get user by id
	user, err := u.UserRepository.GetByID(old.Id)
	if err != nil {
		return Domain{}, err
	}

	if new.Username != user.Username {
		// Check if username already exists
		_, err = u.UserRepository.GetByUsername(new.Username)
		if err == nil {
			return Domain{}, errors.New("username already exists")
		}
	}

	if new.Email != user.Email {
		// Check if email already exists
		_, err = u.UserRepository.GetByEmail(new.Email)
		if err == nil {
			return Domain{}, errors.New("email already exists")
		}
	}

	user.Username = new.Username
	user.Email = new.Email
	user.Name = new.Name
	user.Phone = new.Phone
	user.University = new.University
	user.Position = new.Position
	user.Proposal = new.Proposal
	user.UpdatedAt = primitive.NewDateTimeFromTime(time.Now())

	return u.UserRepository.Update(old, &user)
}

func (u *UserUseCase) UpdatePassword(old *Domain, new *Domain) (Domain, error) {
	// Get user by id
	user, err := u.UserRepository.GetByID(old.Id)
	if err != nil {
		return Domain{}, err
	}

	// Check password bcrypt
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old.Password))
	if err != nil {
		return Domain{}, err
	}

	// Encrypt password
	password, _ := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
	user.Password = string(password)

	return u.UserRepository.Update(old, &user)
}

func (u *UserUseCase) GetAllUsers() ([]Domain, error) {
	return u.UserRepository.GetAll()
}

func (u *UserUseCase) GetByID(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return u.UserRepository.GetByID(ObjId)
}

func (u *UserUseCase) UpdateByAdmin(new *Domain) (Domain, error) {
	// Get user by id
	user, err := u.UserRepository.GetByID(new.Id)
	if err != nil {
		return Domain{}, err
	}

	if new.Username != user.Username {
		// Check if username already exists
		_, err = u.UserRepository.GetByUsername(new.Username)
		if err == nil {
			return Domain{}, errors.New("username already exists")
		}
	}

	if new.Email != user.Email {
		// Check if email already exists
		_, err = u.UserRepository.GetByEmail(new.Email)
		if err == nil {
			return Domain{}, errors.New("email already exists")
		}
	}

	if new.Password != "" {
		// Bcrypt password
		password, _ := bcrypt.GenerateFromPassword([]byte(new.Password), bcrypt.MinCost)
		new.Password = string(password)
	}

	return u.UserRepository.Update(&user, new)
}

func (u *UserUseCase) Delete(id string) (Domain, error) {
	ObjId, _ := primitive.ObjectIDFromHex(id)
	return u.UserRepository.Delete(ObjId)
}