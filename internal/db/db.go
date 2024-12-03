package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/terfo1/postcreator/internal/config"
)

type Database struct {
	*sql.DB
}

func NewDatabase(cfg *config.Config) *Database {
	connStr := "user=" + cfg.DBConfig.User + " password=" + cfg.DBConfig.Password + " dbname=" + cfg.DBConfig.DBName + " host=" + cfg.DBConfig.Host + " port=" + string(cfg.DBConfig.Port) + " sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return &Database{db}
}
