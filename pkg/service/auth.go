package service

import (
	"crypto/sha1"
	"fmt"
	todo "github.com/korinviktor/restWithJWT"
	"github.com/korinviktor/restWithJWT/pkg/repository"
)

const salt = "LYkhOwDpFDfERTmQwBZSWjpmHBIgfLGDUsGperwHjhKpliVI"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = s.GeneratePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GeneratePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
