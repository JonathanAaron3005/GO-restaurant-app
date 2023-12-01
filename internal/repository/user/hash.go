package user

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

func (user *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	argonHash := argon2.IDKey([]byte(password), salt, user.time, user.memory, user.threads, user.keyLen)

	b64Hash := user.encrypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, user.memory, user.time, user.threads, b64Salt, b64Hash)

	return encodedHash, nil
}

func (user *userRepo) encrypt(text []byte) string {
	nonce := make([]byte, user.gcm.NonceSize())

	ciphertext := user.gcm.Seal(nonce, nonce, text, nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}
