package user

import "gorm.io/gorm"

type UserRepository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	Update(user User) (User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (u *UserRepositoryImpl) Save(user User) (User, error) {
	err := u.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByEmail(email string) (User, error) {
	var user User
	err := u.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindById(id int) (User, error) {
	var user User
	err := u.db.Where("id = ?", id).Find(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepositoryImpl) Update(user User) (User, error) {
	err := u.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
