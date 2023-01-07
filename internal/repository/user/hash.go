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

func (ur *userRepo) GenerateUserHash(password string) (hash string, err error) {
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	argonHash := argon2.IDKey([]byte(password), salt, ur.time, ur.memory, ur.threads, ur.keyLen)

	b64Hash := ur.encrypt(argonHash)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)

	encodeResult := fmt.Sprintf(cryptFormat, argon2.Version, ur.memory, ur.time, ur.threads, b64Salt, b64Hash)

	return encodeResult, nil

}

func (ur *userRepo) encrypt(text []byte) string {
	nonce := make([]byte, ur.gcm.NonceSize())

	ciphertext := ur.gcm.Seal(nonce, nonce, text, nil)

	return base64.StdEncoding.EncodeToString(ciphertext)
}
