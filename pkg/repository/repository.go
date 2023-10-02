package repository

import (
	"github.com/Murolando/default_rest_arch/ent"
	"github.com/Murolando/default_rest_arch/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Auth interface {
	SignUp(user ent.User) (int, error)
	GetUserByLoginAndPassword(mail *string, password *string) (int, error)
	SetSession(user int, refresh string, expiredAt string) error
	GetByRefreshToken(refresh string) (int, error)
}

type Repository struct {
	Auth
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Auth: postgres.NewAuthPostgres(db),
	}
}
