package backend

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize   = 16
	keySize    = 32
	iterations = 4096
)

func deriveKey(password []byte, salt []byte) []byte {
	return pbkdf2.Key(password, salt, iterations, keySize, sha256.New)
}

// EncryptAES256 encrypts the given message using AES-256 encryption
func EncryptAES256(key, message string) (string, error) {
	salt := make([]byte, saltSize)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}

	derivedKey := deriveKey([]byte(key), salt)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", err
	}

	plaintext := []byte(message)
	ciphertext := make([]byte, saltSize+aes.BlockSize+len(plaintext))
	copy(ciphertext[:saltSize], salt)
	iv := ciphertext[saltSize : saltSize+aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[saltSize+aes.BlockSize:], plaintext)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAES256 decrypts the given ciphertext using AES-256 decryption
func DecryptAES256(key, ciphertext string) (string, error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(decodedCiphertext) < saltSize+aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}

	salt := decodedCiphertext[:saltSize]
	iv := decodedCiphertext[saltSize : saltSize+aes.BlockSize]
	encryptedMessage := decodedCiphertext[saltSize+aes.BlockSize:]

	derivedKey := deriveKey([]byte(key), salt)

	block, err := aes.NewCipher(derivedKey)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encryptedMessage, encryptedMessage)

	return string(encryptedMessage), nil
}
