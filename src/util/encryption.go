package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"io/ioutil"
	"mime/multipart"
)

// Encrypt encrypts the given plaintext using the given password.
func EncryptData(file multipart.File) ([]byte, error) {
	// Encrypt the file
 key := make([]byte, 32)
 if _, err := rand.Read(key); err != nil {
	 return nil, err
 }
 block, err := aes.NewCipher(key)
 if err != nil {
	 return nil, err
 }
 plaintext, err := ioutil.ReadAll(file)
 if err != nil {
	 return nil, err
 }
 ciphertext := make([]byte, aes.BlockSize+len(plaintext))
 iv := ciphertext[:aes.BlockSize]
 if _, err := rand.Read(iv); err != nil {
	 return nil, err
 }
 stream := cipher.NewCFBEncrypter(block, iv)
 stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
 return ciphertext, nil
}

// Decrypt decrypts the given ciphertext using the given password.
func DecryptData(ciphertext []byte, password string) ([]byte, error) {
	// Generate a key from the password using SHA256.
	key := sha256.Sum256([]byte(password))

	// Extract the initialization vector from the ciphertext.
	iv := ciphertext[:aes.BlockSize]

	// Generate a new cipher block from the key.
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return nil, err
	}

	// Decrypt the ciphertext using AES-CBC mode.
	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, ciphertext[aes.BlockSize:])

	// Verify that the decrypted plaintext is not empty.
	if len(plaintext) == 0 {
		return nil, errors.New("failed to decrypt ciphertext")
	}

	return plaintext, nil
}

// Hash hashes the given data using SHA256.
func Hash(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}