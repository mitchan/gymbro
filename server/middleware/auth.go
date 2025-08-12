package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchan/gymbro/util"
)

const AuthUser = "middleware.auth.user"

func JwtAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")
		if err != nil {
			util.WriteUnauthed(w)
			return
		}

		tokenString := cookie.Value
		if tokenString == "" {
			util.WriteUnauthed(w)
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(util.GetEnv("JWT_SECRET_KEY", "")), nil
		})

		if err != nil || !token.Valid {
			util.WriteUnauthed(w)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			util.WriteUnauthed(w)
			return
		}

		userID, ok := claims["id"].(string)
		if !ok {
			util.WriteUnauthed(w)
			return
		}

		ctx := context.WithValue(r.Context(), AuthUser, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
