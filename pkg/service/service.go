package service

import (
	"github.com/Murolando/default_rest_arch/ent"
	"github.com/Murolando/default_rest_arch/pkg/repository"
)

type Auth interface {
	SignIn(mail *string, password *string) (int, error)
	SignUp(user ent.User) (map[string]interface{}, error)
	GenerateToken(id int) (string, error)
	ParseToken(accesstoken string) (int, error)
	NewRefreshToken(id int) (string, error)
	GetByRefreshToken(refresh string) (int, error)
}

type Service struct {
	Auth
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Auth: NewAuthService(repo),
	}
}
