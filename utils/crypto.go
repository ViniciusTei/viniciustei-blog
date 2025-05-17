package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"

	"golang.org/x/crypto/bcrypt"
)

func Encrypt(key []byte, plaintext string) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := aesGCM.Seal(nonce, nonce, []byte(plaintext), nil)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(key []byte, encrypted string) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := aesGCM.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", err
	}

	nonce, ciphertextClean := ciphertext[:nonceSize], ciphertext[nonceSize:]

	plaintextByte, err := aesGCM.Open(nil, nonce, ciphertextClean, nil)
	if err != nil {
		return "", err
	}

	return string(plaintextByte), nil
}

func HashText(text string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(text), 14)
	return string(b), err
}

func CompareHashText(text, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(text))
	return err == nil
}
