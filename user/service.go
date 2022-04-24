package user

import (
	"errors"
	"os"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(Id int, fileLocation string) (User, error)
	GetUserById(Id int) (User, error)
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

func (u *UserServiceImpl) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := u.userRepository.FindByEmail(email)

	if err != nil {
		return false, err
	}

	if user.Id == 0 {
		return true, nil
	}

	return false, nil
}

func (u *UserServiceImpl) SaveAvatar(Id int, fileLocation string)  (User, error) {
	user, err := u.userRepository.FindById(Id)

	if err != nil {
		return user, err
	}

	if user.AvatarFileName != "" {
		os.Remove(user.AvatarFileName)
	}

	user.AvatarFileName = fileLocation
	updatedUser, err := u.userRepository.Update(user)

	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (u *UserServiceImpl) GetUserById(Id int) (User, error) { 
	user, err := u.userRepository.FindById(Id)

	if err != nil {
		return user, err
	}

	if user.Id == 0 {
		return user, errors.New("User not found")
	}

	return user, nil
}
