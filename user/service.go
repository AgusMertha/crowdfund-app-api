package user

import (
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input UserInput) (User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository}
}

func (u *UserServiceImpl) RegisterUser(input UserInput) (User, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := u.userRepository.Save(user)

	if err != nil {
		return newUser, err
	}

	return newUser, nil
}
