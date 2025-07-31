package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// EncryptAESGCM encrypts plaintext (as []byte) with AES-GCM using key and IV.
// Returns raw ciphertext ([]byte).
func EncryptAESGCM(plaintext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(iv) != gcm.NonceSize() {
		return nil, errors.New("invalid IV size")
	}

	ciphertext := gcm.Seal(nil, iv, plaintext, nil)
	return ciphertext, nil
}

// DecryptAESGCM decrypts raw AES-GCM ciphertext using key and IV.
// Returns raw plaintext ([]byte).
func DecryptAESGCM(ciphertext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(iv) != gcm.NonceSize() {
		return nil, errors.New("invalid IV size")
	}

	plaintext, err := gcm.Open(nil, iv, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// EncryptAESGCMToBase64 encrypts plaintext (string) and returns base64-encoded ciphertext (string).
func EncryptAESGCMToBase64(plaintext string, key, iv []byte) (string, error) {
	ciphertext, err := EncryptAESGCM([]byte(plaintext), key, iv)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// DecryptAESGCMFromBase64 decrypts base64-encoded ciphertext (string) and returns plaintext (string).
func DecryptAESGCMFromBase64(ciphertextB64 string, key, iv []byte) (string, error) {
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return "", err
	}

	plaintext, err := DecryptAESGCM(ciphertext, key, iv)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
