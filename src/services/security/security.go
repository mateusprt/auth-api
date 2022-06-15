package security

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/mateusprt/auth-api/src/models"
	"github.com/subosito/gotenv"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) string {
	passwordBytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordBytes)
}

func PasswordMatch(passwordUser string, passwordReceived string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordUser), []byte(passwordReceived))
	return err == nil
}

func GenerateToken() string {
	return uuid.NewString()
}

func GenerateJWT(user models.User) string {

	gotenv.Load()

	permissions := jwt.MapClaims{
		"authorized": true,
		"exp":        time.Now().Add(time.Hour * 6).Unix(),
		"user_id":    user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	mySecret := os.Getenv("SECRET_KEY")
	tokenGenerated, _ := token.SignedString([]byte(mySecret))
	return tokenGenerated
}

func ValidateToken(r *http.Request) error {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		fmt.Println("Deu erro")
		return err
	}

	// retorna todos os claims
	// _, ok := token.Claims.(jwt.MapClaims)
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Invalid token.")
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func returnVerificationKey(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)

	if !ok {
		return nil, fmt.Errorf("Unexpected sign method. %v", token.Header["alg"])
	}

	gotenv.Load()
	return []byte(os.Getenv("SECRET_KEY")), nil
}

func ExtractUserId(r *http.Request) (int, error) {
	tokenString := extractToken(r)
	token, err := jwt.Parse(tokenString, returnVerificationKey)

	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["user_id"]), 10, 64)

		if err != nil {
			return 0, err
		}
		return int(userId), nil
	}
	return 0, errors.New("Invalid token")
}
