package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/mitchan/gymbro/model"
	"github.com/mitchan/gymbro/repository"
	"github.com/mitchan/gymbro/util"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (userService *UserService) CreateUser(req model.CreateUser) error {
	log.Printf("UserService.CreateUser - Starting user creation for: %s", req.Email)

	if req.Email == "" || req.Username == "" || req.Password == "" {
		log.Printf("UserService.CreateUser - Validation failed: missing required fields")
		return fmt.Errorf("username, email, and password are required")
	}

	if len(req.Password) < 6 {
		log.Printf("UserService.CreateUser - Validation failed: password too short")
		return fmt.Errorf("password must be at least 6 characters")
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		log.Printf("UserService.CreateUser - Password hashing failed: %v", err)
		return fmt.Errorf("failed to process password")
	}

	u := &repository.User{
		Email:        req.Email,
		Username:     req.Username,
		PasswordHash: &hashedPassword,
	}

	_, err = userService.userRepo.CreateUser(u)
	if err != nil {
		log.Printf("UserService.CreateUser - Database error: %v", err)
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique") {
			return fmt.Errorf("username or email already exists")
		}
		return fmt.Errorf("failed to create user: %v", err)
	}

	return nil
}
