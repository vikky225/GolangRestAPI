package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "supersecretkey"

func GenerateToken(email string, userId int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

// we can do clean seperation between verification and extration
func ValidateToken(token string) (int64, error) {
	parsedToekn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// we should veify signed with HS256
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, errors.New("invalid token")
	}
	tokenIsValid := parsedToekn.Valid
	if !tokenIsValid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToekn.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("invalid token claims")
	}
	//email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))
	return userId, nil

}
