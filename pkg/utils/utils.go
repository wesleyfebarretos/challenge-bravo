package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"io"
	"math"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashPassword), nil
}

func IsValidPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func Encrypt(plaintext, secret string) (string, error) {
	// Hash the API token to create a 32-byte key
	hash := sha256.Sum256([]byte(secret))
	key := hash[:]

	// Create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", nil
	}

	// Create a byte slice to hold the IV + ciphertext
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	// Generate a random IV and store it at the beginning of the ciphertext slice
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Create the CFB encrypter and XOR the plaintext with it to create the ciphertext
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	// Encode the ciphertext to base64 to make it easily transferable
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func Decrypt(ciphertextBase64, secret string) (string, error) {
	// Decode the base64-encoded ciphertext
	ciphertext, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		return "", err
	}

	// Hash the API token to create a 32-byte key
	hash := sha256.Sum256([]byte(secret))
	key := hash[:]

	// Create a new AES cipher with the key
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Ensure the ciphertext is long enough to contain the IV
	if len(ciphertext) < aes.BlockSize {
		return "", err
	}

	// Extract the IV from the beginning of the ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Create the CFB decrypter and XOR the ciphertext with it to recover the plaintext
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	// Return the decrypted plaintext as a string
	return string(ciphertext), nil
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
