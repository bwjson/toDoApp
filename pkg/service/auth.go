package service

import (
	"crypto/sha1"
	"fmt"
	"github.com/bwjson/toDoApp"
	"github.com/bwjson/toDoApp/pkg/repository"
)

const salt = "snnfaif13h813h1ni191n"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user toDoApp.User) (int, error) {
	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
