package auth

import (
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(userID string) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

var secret_key = os.Getenv("JWT_SECRET")

func (s *jwtService) GenerateToken(userID string) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signToken, err := token.SignedString([]byte(secret_key))
	if err != nil {
		return signToken, err
	}

	return signToken, nil
}
func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		return []byte(secret_key), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
