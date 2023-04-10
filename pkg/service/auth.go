package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/aidos-dev/toDoApp"
	"github.com/aidos-dev/toDoApp/pkg/repository"
)

const salt = "lk65vm29vkf437fb817hfn3857kdn4nv"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
