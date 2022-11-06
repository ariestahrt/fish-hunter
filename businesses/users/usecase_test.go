package users_test

import (
	"errors"
	"fish-hunter/businesses/users"
	_mockUser "fish-hunter/businesses/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	mockUserRepository _mockUser.Repository
	userService 	  users.UseCase
	userDomain		users.Domain
)

/*
type Domain struct {
	Id        primitive.ObjectID    `bson:"_id,omitempty" json:"_id"`
	Username  string 			  `bson:"username" json:"username"`
	Email     string			  `bson:"email" json:"email"`
	Password  string 			  `bson:"password" json:"password"`
	NewPassword  string 		  `bson:"new_password" json:"new_password"`
	IsActive  bool 				  `bson:"is_active" json:"is_active"`
	Name	  string 			  `bson:"name" json:"name"`
	Phone	  string 			  `bson:"phone" json:"phone"`
	University string 			  `bson:"university" json:"university"`
	Position   string 				`bson:"position" json:"position"` // Student, Lecturer, Staff 
	Proposal   string			`bson:"proposal" json:"proposal"`
	Roles     []string		  `bson:"roles" json:"roles"`
	CreatedAt primitive.DateTime 		  `bson:"created_at" json:"created_at"`
	UpdatedAt primitive.DateTime 		  `bson:"updated_at" json:"updated_at"`
	DeletedAt primitive.DateTime 		  `bson:"deleted_at" json:"deleted_at"`
}
*/

func TestMain(m *testing.M) {
	userService = users.NewUserUseCase(&mockUserRepository)
	userDomain = users.Domain{
		Id:        primitive.NewObjectID(),
		Username:  "hina_setoguchi",
		Email:     "test@ganteng.com",
		Password:  "nasipadang5000",
		IsActive:  true,
		Name:	  "Hina Setoguchi",
		Phone:	  "081234567890",
		University: "Universitas Sakura",
		Position:   "student", // Student, Lecturer, Staff
		Proposal:   "https://drive.google.com/asdfghjkl",
		Roles:     []string{"guest"},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		DeletedAt: primitive.NewDateTimeFromTime(time.Time{}),
	}
	m.Run()
}

// Test Register
func TestRegister(t *testing.T) {
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		mockUserRepository.On("Register", &userDomain).Return(userDomain, nil).Once()
		result, err := userService.Register(&userDomain)
		assert.Nil(t, err)
		assert.Equal(t, userDomain, result)
	})

	t.Run("Test Case 2 | Invalid Register", func(t *testing.T) {
		mockUserRepository.On("Register", &userDomain).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.Register(&userDomain)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Login
func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		mockUserRepository.On("Login", &userDomain).Return(userDomain, nil).Once()
		result, err := userService.Login(&userDomain)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Login", func(t *testing.T) {
		mockUserRepository.On("Login", &userDomain).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.Login(&userDomain)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Get Profile
func TestGetProfile(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Profile", func(t *testing.T) {
		mockUserRepository.On("GetProfile", userDomain.Id.Hex()).Return(userDomain, nil).Once()
		result, err := userService.GetProfile(userDomain.Id.Hex())
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Get Profile", func(t *testing.T) {
		mockUserRepository.On("GetProfile", userDomain.Id.Hex()).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.GetProfile(userDomain.Id.Hex())
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Update Profile
func TestUpdateProfile(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Profile", func(t *testing.T) {
		mockUserRepository.On("UpdateProfile", &userDomain).Return(userDomain, nil).Once()
		result, err := userService.UpdateProfile(&userDomain)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Update Profile", func(t *testing.T) {
		mockUserRepository.On("UpdateProfile", &userDomain).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.UpdateProfile(&userDomain)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Update Password
func TestUpdatePassword(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Password", func(t *testing.T) {
		mockUserRepository.On("UpdatePassword", &userDomain).Return(userDomain, nil).Once()
		result, err := userService.UpdatePassword(&userDomain)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Update Password", func(t *testing.T) {
		mockUserRepository.On("UpdatePassword", &userDomain).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.UpdatePassword(&userDomain)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Delete User
func TestDeleteUser(t *testing.T) {
	t.Run("Test Case 1 | Valid Delete User", func(t *testing.T) {
		mockUserRepository.On("Delete", userDomain.Id.Hex()).Return(userDomain, nil).Once()
		result, err := userService.Delete(userDomain.Id.Hex())
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Delete User", func(t *testing.T) {
		mockUserRepository.On("Delete", userDomain.Id.Hex()).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.Delete(userDomain.Id.Hex())
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Get All User
func TestGetAllUser(t *testing.T) {
	t.Run("Test Case 1 | Valid Get All User", func(t *testing.T) {
		mockUserRepository.On("GetAllUsers").Return([]users.Domain{userDomain}, nil).Once()
		result, err := userService.GetAllUsers()
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Get All User", func(t *testing.T) {
		mockUserRepository.On("GetAllUsers").Return([]users.Domain{}, errors.New("")).Once()
		result, err := userService.GetAllUsers()
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Get User By ID
func TestGetUserByID(t *testing.T) {
	t.Run("Test Case 1 | Valid Get User By ID", func(t *testing.T) {
		mockUserRepository.On("GetByID", userDomain.Id.Hex()).Return(userDomain, nil).Once()
		result, err := userService.GetByID(userDomain.Id.Hex())
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Get User By ID", func(t *testing.T) {
		mockUserRepository.On("GetByID", userDomain.Id.Hex()).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.GetByID(userDomain.Id.Hex())
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

// Test Update User
func TestUpdateUser(t *testing.T) {
	t.Run("Test Case 1 | Valid Update User", func(t *testing.T) {
		mockUserRepository.On("Update", &userDomain).Return(userDomain, nil).Once()
		result, err := userService.Update(&userDomain)
		assert.Nil(t, err)
		assert.NotEmpty(t, result)
	})

	t.Run("Test Case 2 | Invalid Update User", func(t *testing.T) {
		mockUserRepository.On("Update", &userDomain).Return(users.Domain{}, errors.New("")).Once()
		result, err := userService.Update(&userDomain)
		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}