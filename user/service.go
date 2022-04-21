package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
}

type UserServiceImpl struct {
	userRepository UserRepository
}

func NewUserService(userRepository UserRepository) *UserServiceImpl {
	return &UserServiceImpl{userRepository}
}

func (u *UserServiceImpl) RegisterUser(input RegisterUserInput) (User, error) {
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

func (u *UserServiceImpl) Login(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("Invalid email")
	}

	// match password
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))

	if err != nil {
		return user, errors.New("password not correct")
	}

	return user, nil
}
