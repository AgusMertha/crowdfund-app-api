package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) (User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) Save(user User) (User, error) {
	err := u.db.Create(&user).Error

	if  err != nil {
		return user, err
	}

	return user, nil
}