package user

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"time"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type userRepo struct {
	db        *gorm.DB
	gcm       cipher.AEAD
	time      uint32
	memory    uint32
	threads   uint8
	keyLen    uint32
	signKey   *rsa.PrivateKey
	accessExp time.Duration
}

func GetRepository(db *gorm.DB, secret string, time uint32, memory uint32, threads uint8, keyLen uint32,
	signKey *rsa.PrivateKey, accessExp time.Duration) (Repository, error) {

	block, err := aes.NewCipher([]byte(secret))

	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {
		return nil, err
	}

	return &userRepo{
		db:        db,
		gcm:       gcm,
		time:      time,
		memory:    memory,
		threads:   threads,
		keyLen:    keyLen,
		signKey:   signKey,
		accessExp: accessExp,
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

func (user *userRepo) VerifyLogin(username, password string, userData model.User) (bool, error) {
	if username != userData.Username {
		return false, nil
	}

	verified, err := user.comparePassword(password, userData.Hash)

	if err != nil {
		return false, err
	}

	return verified, nil

}

func (user *userRepo) GetUserData(username string) (model.User, error) {
	var userData model.User

	if err := user.db.Where(model.User{Username: username}).First(&userData).Error; err != nil {
		return model.User{}, err
	}

	return userData, nil
}
