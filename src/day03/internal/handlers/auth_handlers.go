package handlers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

var signingKey = []byte("your_secret_key")

func GetToken(w http.ResponseWriter, r *http.Request) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
		Subject:   "umaradri",
	})

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Token generation error")
		return
	}

	res := struct {
		Token string `json:"token"`
	}{
		Token: tokenString,
	}

	respondWithJSON(w, http.StatusCreated, res)
}

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "Missing Authorization header")
			return
		}

		parts := strings.Split(authHeader, " ")

		if len(parts) != 2 || parts[0] != "Bearer" {
			respondWithError(w, http.StatusUnauthorized, "Invalid Authorization header format")
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return signingKey, nil
		})

		if !token.Valid {
			respondWithError(w, http.StatusUnauthorized, "Token is not valid")
			return
		}

		if err != nil {
			respondWithError(w, http.StatusUnauthorized, fmt.Sprintf("Error parsing token: %s", err.Error()))
			return
		}

		next.ServeHTTP(w, r)
	}
}
