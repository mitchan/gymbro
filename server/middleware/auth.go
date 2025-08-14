package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/mitchan/gymbro/model"
	"github.com/mitchan/gymbro/repository"
	"github.com/mitchan/gymbro/util"
)

const AuthUser = "middleware.auth.user"

type AuthMiddleware struct {
	userRepo *repository.UserRepository
}

func NewAuthMiddleware(userRepo *repository.UserRepository) *AuthMiddleware {
	return &AuthMiddleware{
		userRepo: userRepo,
	}
}

func (authMiddleware *AuthMiddleware) JwtAuth(next http.Handler) http.Handler {
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

		// TODO: validate userID with userRepo

		ctx := context.WithValue(r.Context(), AuthUser, userID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserIDFromContext(ctx context.Context) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	id, ok := ctx.Value(AuthUser).(string)
	fmt.Println("id is: " + id)
	if !ok {
		log.Printf("UserService.Me - cannot retrieve AuthUser from context")
		return uuid.New(), model.UnauthedError
	}

	parsed, err := uuid.Parse(id)
	if err != nil {
		log.Printf("UserService.Me - cannot convert %s to uuid", id)
		return uuid.New(), model.UnauthedError
	}

	return parsed, nil
}
