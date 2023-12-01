package user

import (
	"crypto/aes"
	"crypto/cipher"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db      *gorm.DB
	gcm     cipher.AEAD
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func GetRepository(db *gorm.DB, secret string, time uint32, memory uint32, threads uint8, keyLen uint32) (Repository, error) {
	block, err := aes.NewCipher([]byte(secret))

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	return &userRepo{
		db:      db,
		gcm:     gcm,
		time:    time,
		memory:  memory,
		threads: threads,
		keyLen:  keyLen,
	}, nil
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
