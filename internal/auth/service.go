package auth

import (
	"errors"
	jwtutil "robot-server/internal/jwt"
	"robot-server/internal/models"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repository
}

func (s *Service) Register(u models.User) (string, string, error) {

	hash, _ := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)

	u.UUID = uuid.New().String()
	u.Password = string(hash)

	err := s.Repo.CreateUser(u)
	if err != nil {
		return "", "", err
	}

	token, _ := jwtutil.GenerateToken(u.UUID)

	return token, u.UUID, nil
}

func (s *Service) Login(email, password string) (string, string, error) {

	u, err := s.Repo.GetByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	token, _ := jwtutil.GenerateToken(u.UUID)

	return token, u.UUID, nil
}
