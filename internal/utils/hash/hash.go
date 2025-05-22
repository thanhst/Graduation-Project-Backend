package CustomHash

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"hash/fnv"

	"golang.org/x/crypto/bcrypt"
)

func Hash32(s string) uint32 {
	h := fnv.New32a() // FNV-1a 32-bit
	h.Write([]byte(s))
	return h.Sum32()
}
func HashMD5(s string) string {
	hash := md5.Sum([]byte(s))
	return hex.EncodeToString(hash[:])
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func GenerateSecretKey(plainText string) []byte {
	hash := sha256.Sum256([]byte(plainText))
	return hash[:]
}
