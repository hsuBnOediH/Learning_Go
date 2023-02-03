package dbrepo

import (
	"database/sql"

	"github.com/hsuBnOediH/Learning_Go/internal/config"
	"github.com/hsuBnOediH/Learning_Go/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
