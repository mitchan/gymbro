package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/mitchan/gymbro/middleware"
	"github.com/mitchan/gymbro/model"
	"github.com/mitchan/gymbro/repository"
	"github.com/mitchan/gymbro/util"
)

type UserService struct {
	userRepo *repository.UserRepository
}

type JWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (userService *UserService) CreateUser(req model.CreateUser) (*model.ResponseLoginUser, error) {
	log.Printf("UserService.CreateUser - Starting user creation for: %s", req.Email)

	if req.Email == "" || req.Username == "" || req.Password == "" {
		log.Printf("UserService.CreateUser - Validation failed: missing required fields")
		return nil, fmt.Errorf("username, email, and password are required")
	}

	if len(req.Password) < 8 {
		log.Printf("UserService.CreateUser - Validation failed: password too short")
		return nil, fmt.Errorf("password must be at least 6 characters")
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		log.Printf("UserService.CreateUser - Password hashing failed: %v", err)
		return nil, fmt.Errorf("failed to process password")
	}

	u := &repository.User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: &hashedPassword,
	}

	user, err := userService.userRepo.CreateUser(u)
	if err != nil {
		log.Printf("UserService.CreateUser - Database error: %v", err)
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			return nil, fmt.Errorf("username or email already exists")
		}
		return nil, fmt.Errorf("failed to create user: %v", err)
	}

	log.Printf("UserService.CreateUser - User created successfully in database: %s", user.ID.String())

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID:       user.ID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	secretKey := util.GetEnv("JWT_SECRET_KEY", "")
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &model.ResponseLoginUser{
		AccessToken: ss,
		Username:    user.Username,
		ID:          user.ID.String(),
	}, nil
}

func (userService *UserService) Login(req model.LoginUser) (*model.ResponseLoginUser, error) {
	log.Printf("UserService.Login - Starting login for: %s", req.Email)

	if req.Email == "" || req.Password == "" {
		log.Printf("UserService.Login - Validation failed: missing required fields")
		return nil, fmt.Errorf("email and password are required")
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		log.Printf("UserService.CreateUser - Password hashing failed: %v", err)
		return nil, fmt.Errorf("failed to process password")
	}

	u := &repository.User{
		Email:        req.Email,
		PasswordHash: &hashedPassword,
	}

	user, err := userService.userRepo.GetUserByEmail(u)
	if err != nil {
		log.Printf("UserService.Login - Database error: %v", err)
		return nil, fmt.Errorf("failed to authenticate user")
	}

	if user == nil {
		log.Printf("UserService.Login - User not found for email: %s", req.Email)
		return nil, fmt.Errorf("invalid email or password")
	}

	err = util.CheckPassword(req.Password, *user.PasswordHash)
	if err != nil {
		log.Printf("UserService.Login - Password check failed for user: %s", user.ID.String())
		return nil, fmt.Errorf("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID:       user.ID.String(),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID.String(),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	secretKey := util.GetEnv("JWT_SECRET_KEY", "")
	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	return &model.ResponseLoginUser{
		AccessToken: ss,
		Username:    user.Username,
		ID:          user.ID.String(),
	}, nil
}

func (userService *UserService) Me(ctx context.Context) (*model.ResponseMeUser, error) {
	userID, err := middleware.GetUserIDFromContext(ctx)
	if err != nil {
		log.Printf("UserService.Me - cannot retrieve AuthUser from context")
		return nil, model.UnauthedError
	}

	u := &repository.User{
		ID: userID,
	}

	user, err := userService.userRepo.GetUserByID(u)
	if err != nil || user == nil {
		log.Printf("UserService.Me - cannot get user from id: %s", userID.String())
		return nil, model.UnauthedError
	}

	return &model.ResponseMeUser{
		Username: user.Username,
		ID:       user.ID.String(),
	}, nil
}
