package user

import (
	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &userRepo{
		db: db,
	}
}

func (user *userRepo) RegisterUser(userData model.User) (model.User, error) {
	if err := user.db.Create(&userData).Error; err != nil {
		return model.User{}, err
	}

	return userData, nil
}

func (user *userRepo) CheckRegistered(username string) (bool, error) {
	var userData model.User

	if err := user.db.Where(model.User{Username: username}).First(&userData).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		} else {
			return false, err
		}
	}

	return userData.ID != "", nil
}

func (user *userRepo) GenerateUserHash(password string) (hash string, err error) {
	return "", nil
}
