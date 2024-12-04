package post

import (
	"github.com/terfo1/postcreator/internal/db"
	"net/http"
)

type Repository struct {
	DB *db.Database
}

func NewRepository(db *db.Database) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) HandlePosts(w http.ResponseWriter, req *http.Request) {
	// Логика обработки запросов к постам
}
