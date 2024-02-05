package user

import (
	"context"
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"github.com/JonathanAaron3005/go-restaurant-app/internal/tracing"
	"golang.org/x/crypto/argon2"
)

const (
	cryptFormat = "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
)

func (user *userRepo) GenerateUserHash(ctx context.Context, password string) (hash string, err error) {
	ctx, span := tracing.CreateSpan(ctx, "GenerateUserHash")
	defer span.End()

	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	argonHash := argon2.IDKey([]byte(password), salt, user.time, user.memory, user.threads, user.keyLen)

	b64Hash := user.encrypt(ctx, argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodedHash := fmt.Sprintf(cryptFormat, argon2.Version, user.memory, user.time, user.threads, b64Salt, b64Hash)

	return encodedHash, nil
}

func (user *userRepo) encrypt(ctx context.Context, text []byte) string {
	_, span := tracing.CreateSpan(ctx, "encrypt")
	defer span.End()

	nonce := make([]byte, user.gcm.NonceSize())

	ciphertext := user.gcm.Seal(nonce, nonce, text, nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}

func (user *userRepo) decrypt(ctx context.Context, ciphertext string) ([]byte, error) {
	_, span := tracing.CreateSpan(ctx, "decrypt")
	defer span.End()

	decoded, err := base64.StdEncoding.DecodeString(ciphertext)

	if err != nil {
		return nil, err
	}

	if len(decoded) < user.gcm.NonceSize() {
		return nil, errors.New("invalid nonce size")
	}

	return user.gcm.Open(nil, decoded[:user.gcm.NonceSize()], decoded[user.gcm.NonceSize():], nil)
}

func (user *userRepo) comparePassword(ctx context.Context, password, hash string) (bool, error) {
	_, span := tracing.CreateSpan(ctx, "comparePassword")
	defer span.End()

	parts := strings.Split(hash, "$")

	var memory, time uint32
	var pararellism uint8

	switch parts[1] {
	case "argon2id":
		_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &pararellism)

		if err != nil {
			return false, err
		}

		salt, err := base64.RawStdEncoding.DecodeString(parts[4])

		if err != nil {
			return false, err
		}

		hash := parts[5]

		decrpytedHash, err := user.decrypt(ctx, hash)

		if err != nil {
			return false, err
		}

		var keyLen = uint32(len(decrpytedHash))

		comparisonHash := argon2.IDKey([]byte(password), salt, time, memory, pararellism, keyLen)

		return subtle.ConstantTimeCompare(comparisonHash, decrpytedHash) == 1, nil
	}

	return false, nil
}
