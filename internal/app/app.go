package app

import (
	"github.com/terfo1/postcreator/internal/config"
	"github.com/terfo1/postcreator/internal/db"
	"github.com/terfo1/postcreator/internal/post"
	"github.com/terfo1/postcreator/internal/user"
	"log"
	"net/http"
)

type App struct {
	Config   *config.Config
	DB       *db.Database
	PostRepo *post.Repository
	UserRepo *user.Repository
	Router   *http.ServeMux
}

func New() *App {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error in loading config", err)
	}
	database := db.NewDatabase(cfg)
	postRepo := post.NewRepository(database)
	userRepo := user.NewRepository(database)
	router := http.NewServeMux()
	return &App{
		Config:   cfg,
		DB:       database,
		PostRepo: postRepo,
		UserRepo: userRepo,
		Router:   router,
	}
}
func (a *App) Run() error {
	// Инициализация маршрутов
	a.initializeRoutes()
	// Запуск сервера
	return http.ListenAndServe(":"+a.Config.Server.Port, a.Router)
}

func (a *App) initializeRoutes() {
	// Инициализация маршрутов для постов и пользователей
	a.Router.HandleFunc("/posts", a.PostRepo.HandlePosts)
	a.Router.HandleFunc("/users", a.UserRepo.HandleUsers)
}
