package pkg

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

func hashAndSaltPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	hashedString := string(hashed)
	return hashedString
}

func comparePasswordHashes(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		// Sleep for a random amount of time to prevent password timing attacks
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		time.Sleep(time.Duration(n) * time.Millisecond)
		return false
	}
	return true
}

func generateAPIKey() string {
	return strings.Replace(uuid.New().String()+uuid.New().String(), "-", "", -1)
}
