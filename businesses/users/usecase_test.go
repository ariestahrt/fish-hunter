package users_test

import (
	"errors"
	"fish-hunter/businesses/users"
	_mockUser "fish-hunter/businesses/users/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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
		IsActive:  false,
		Name:	  "Hina Setoguchi",
		Phone:	  "081234567890",
		University: "Universitas Sakura",
		Position:   "student", // Student, Lecturer, Staff
		Proposal:   "https://drive.google.com/asdfghjkl",
		Roles:     []string{"guest"},
		CreatedAt: primitive.NewDateTimeFromTime(time.Now()),
		UpdatedAt: primitive.NewDateTimeFromTime(time.Now()),
		DeletedAt: primitive.NewDateTimeFromTime(primitive.NilObjectID.Timestamp()),
	}
	m.Run()
}

// Test Register
func TestRegister(t *testing.T) {
	t.Run("Test Case 1 | Valid Register", func(t *testing.T) {
		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("Not Found")).Once()

		// Mock on GetByUsername
		mockUserRepository.On("GetByUsername", userDomain.Username).Return(users.Domain{}, errors.New("Not Found")).Once()
		
		// Mock on Create
		mockUserRepository.On("Create", &userDomain).Return(userDomain, nil).Once()

		_, err := userService.Register(&userDomain)
		assert.Nil(t, err)
	})

	// Email already exist
	t.Run("Test Case 2 | Email already exist", func(t *testing.T) {
		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(userDomain, nil).Once()

		_, err := userService.Register(&userDomain)
		assert.NotNil(t, err)
	})

	// Username already exist
	t.Run("Test Case 3 | Username already exist", func(t *testing.T) {
		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("Not Found")).Once()

		// Mock on GetByUsername
		mockUserRepository.On("GetByUsername", userDomain.Username).Return(userDomain, nil).Once()

		_, err := userService.Register(&userDomain)
		assert.NotNil(t, err)
	})
}

// Test Login
func TestLogin(t *testing.T) {
	t.Run("Test Case 1 | Valid Login", func(t *testing.T) {
		// Mock on GetByEmail
		mockRes := userDomain
		mockRes.IsActive = true

		userDomain.Password = "nasipadang5000"
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(mockRes, nil).Once()

		_, err := userService.Login(&userDomain)
		assert.Nil(t, err)
	})

	t.Run("Test Case 2 | Invalid Login", func(t *testing.T) {
		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("Not Found")).Once()

		_, err := userService.Login(&userDomain)
		assert.NotNil(t, err)
	})

	// User Inactive
	t.Run("Test Case 3 | User Inactive", func(t *testing.T) {
		// Mock on GetByEmail
		mockRes := userDomain
		mockRes.IsActive = false

		userDomain.Password = "nasipadang5000"
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(mockRes, nil).Once()

		_, err := userService.Login(&userDomain)
		assert.NotNil(t, err)
	})

	// Invalid password
	t.Run("Test Case 4 | Invalid password", func(t *testing.T) {
		// Mock on GetByEmail
		mockRes := userDomain
		mockRes.IsActive = true

		userDomain.Password = "nasipadang5001"
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(mockRes, nil).Once()

		_, err := userService.Login(&userDomain)
		assert.NotNil(t, err)
	})
}

// Test GetProfile
func TestGetProfile(t *testing.T) {
	t.Run("Test Case 1 | Valid Get Profile", func(t *testing.T) {
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		_, err := userService.GetProfile(userDomain.Id.Hex())
		assert.Nil(t, err)
	})
}

// Update Profile
func TestUpdateProfile(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Profile", func(t *testing.T) {
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("Not Found")).Once()

		// Mock on GetByUsername
		mockUserRepository.On("GetByUsername", userDomain.Username).Return(users.Domain{}, errors.New("Not Found")).Once()

		// Mock on Update
		newVal := userDomain
		newVal.Name = "Hina Setoguchi"
		newVal.Phone = "081234567890"
		newVal.University = "Universitas Sakura"
		newVal.Position = "Updated"

		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(newVal, nil).Once()

		_, err := userService.UpdateProfile(&userDomain, &newVal)
		assert.Nil(t, err)
	})

	// Email already exist
	t.Run("Test Case 2 | Email already exist", func(t *testing.T) {
		newVal := userDomain
		newVal.Email = "newemail@nyakit.in"

		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", newVal.Email).Return(userDomain, nil).Once()

		_, err := userService.UpdateProfile(&userDomain, &newVal)
		assert.NotNil(t, err)
	})

	// Username already exist
	t.Run("Test Case 3 | Username already exist", func(t *testing.T) {
		newVal := userDomain
		newVal.Username = "newusername"

		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", userDomain.Email).Return(users.Domain{}, errors.New("Not Found")).Once()

		// Mock on GetByUsername
		mockUserRepository.On("GetByUsername", newVal.Username).Return(userDomain, nil).Once()

		_, err := userService.UpdateProfile(&userDomain, &newVal)
		assert.NotNil(t, err)
	})

	// Failed get by Id
	t.Run("Test Case 4 | Failed get by Id", func(t *testing.T) {
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, errors.New("Not Found")).Once()

		_, err := userService.UpdateProfile(&userDomain, &userDomain)
		assert.NotNil(t, err)
	})
}

// Test Update Password
func TestUpdatePassword(t *testing.T) {
	t.Run("Test Case 1 | Valid Update Password", func(t *testing.T) {
		userDomain.Password = "nasipadang5000"
		mockUser := userDomain
		
		// Bcrypt password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("nasipadang5000"), bcrypt.DefaultCost)
		mockUser.Password = string(hashedPassword)
		
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(mockUser, nil).Once()

		// Mock on Update
		newVal := userDomain
		newVal.Password = "nasipadang5001"

		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(newVal, nil).Once()

		_, err := userService.UpdatePassword(&userDomain, &newVal)
		assert.Nil(t, err)
	})

	// Failed get by Id
	t.Run("Test Case 2 | Failed get by Id", func(t *testing.T) {
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, errors.New("Not Found")).Once()

		_, err := userService.UpdatePassword(&userDomain, &userDomain)
		assert.NotNil(t, err)
	})

	// Old Password is wrong
	t.Run("Test Case 3 | Old Password is wrong", func(t *testing.T) {
		userDomain.Password = "xxxxxxxxxxxx"
		mockUser := userDomain
		
		// Bcrypt password
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("nasipadang5001"), bcrypt.DefaultCost)
		mockUser.Password = string(hashedPassword)
		
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(mockUser, nil).Once()

		_, err := userService.UpdatePassword(&userDomain, &userDomain)
		assert.NotNil(t, err)
	})
}

// Test Get All User
func TestGetAllUser(t *testing.T) {
	t.Run("Test Case 1 | Valid Get All User", func(t *testing.T) {
		// Mock on GetAll
		mockUserRepository.On("GetAll").Return([]users.Domain{userDomain}, nil).Once()

		_, err := userService.GetAllUsers()
		assert.Nil(t, err)
	})
}

// Test GetByID
func TestGetById(t *testing.T) {
	t.Run("Test Case 1 | Valid GetByID", func(t *testing.T) {
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		result, err := userService.GetByID(userDomain.Id.Hex())
		assert.Nil(t, err)
		assert.Equal(t, userDomain, result)
	})
}

// Test UpdateByAdmin
func TestUpdateByAdmin(t *testing.T) {
	t.Run("Test Case 1 | Valid UpdateByAdmin", func(t *testing.T) {
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		// Mock on Update
		newVal := userDomain
		newVal.Name = "Hina Setoguchi"
		newVal.Phone = "081234567890"
		newVal.University = "Universitas Sakura"
		newVal.Position = "Updated"
		newVal.Roles = []string{"ok"}

		mockUserRepository.On("Update", mock.Anything, mock.Anything).Return(newVal, nil).Once()

		_, err := userService.UpdateByAdmin(&newVal)
		assert.Nil(t, err)
	})

	// Failed get by ID
	t.Run("Test Case 2 | Failed get by ID", func(t *testing.T) {
		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(users.Domain{}, errors.New("Not Found")).Once()

		_, err := userService.UpdateByAdmin(&userDomain)
		assert.NotNil(t, err)
	})

	// Email already exist
	t.Run("Test Case 3 | Email already exist", func(t *testing.T) {
		newVal := userDomain
		newVal.Email = "asdasd@nyakit.in"

		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		// Mock on GetByEmail
		mockUserRepository.On("GetByEmail", newVal.Email).Return(userDomain, nil).Once()

		_, err := userService.UpdateByAdmin(&newVal)
		assert.NotNil(t, err)
	})

	// Username already exist
	t.Run("Test Case 4 | Username already exist", func(t *testing.T) {
		newVal := userDomain
		newVal.Username = "asdasd"

		// Mock on GetByID
		mockUserRepository.On("GetByID", userDomain.Id).Return(userDomain, nil).Once()

		// Mock on GetByUsername
		mockUserRepository.On("GetByUsername", newVal.Username).Return(userDomain, nil).Once()

		_, err := userService.UpdateByAdmin(&newVal)
		assert.NotNil(t, err)
	})
}

// Test Delete
func TestDelete(t *testing.T) {
	t.Run("Test Case 1 | Valid Delete", func(t *testing.T) {
		// Mock on Delete
		mockUserRepository.On("Delete", userDomain.Id).Return(userDomain, nil).Once()

		_, err := userService.Delete(userDomain.Id.Hex())
		assert.Nil(t, err)
	})
}