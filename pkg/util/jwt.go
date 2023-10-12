package util

import (
	"errors"
	"fmt"
	"github.com/DarioKnezovic/campaign-service/config"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strings"
	"time"
)

type Claims struct {
	Id    uint   `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func VerifyJWT(tokenString string, secretKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse JWT token: %w", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid JWT token")
	}

	// Check if the token has expired
	if time.Now().Unix() > claims.ExpiresAt {
		return nil, fmt.Errorf("JWT token has expired")
	}

	return claims, nil
}

func GetUserIdFromToken(r *http.Request) (uint, error) {
	token := strings.Split(r.Header.Get("Authorization"), " ")
	cfg := config.LoadConfig()

	if len(token) != 2 {
		return 0, errors.New("Invalid Authorization header")
	}

	// TODO: In the future add GRPC call to User service to check token
	claims, err := VerifyJWT(token[1], []byte(cfg.JWTSecretKey))
	if err != nil {
		log.Printf("Error during verifying JWT: %v", err)
		return 0, errors.New("Error during verifying JWT")
	}

	userId := claims.Id
	if userId == 0 {
		log.Print("User ID is not available from JWT token")
		return 0, errors.New("Undefined User ID from Authorization token")
	}

	return userId, nil
}
