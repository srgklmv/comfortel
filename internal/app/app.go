package app

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/srgklmv/comfortel/internal/api"
	"github.com/srgklmv/comfortel/internal/config"
	"github.com/srgklmv/comfortel/internal/controller"
	"github.com/srgklmv/comfortel/internal/repository"
	"github.com/srgklmv/comfortel/internal/usecase"
	"github.com/srgklmv/comfortel/pkg/database"
)

type app struct {
	engine *gin.Engine
	conn   *sql.DB
}

func New() *app {
	// TODO: Logs as JSON.
	return &app{
		engine: gin.Default(),
	}
}

func (a *app) Run() error {
	cfg, err := config.Init()
	if err != nil {
		return fmt.Errorf("config.Init: %w", err)
	}

	conn, err := database.New(
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
		cfg.Database.User,
		cfg.Database.Password,
	)
	if err != nil {
		return fmt.Errorf("database.New: %w", err)
	}
	a.conn = conn

	err = database.Migrate(conn, "file://migrations", 1)
	if err != nil {
		return fmt.Errorf("database.Migrate: %w", err)
	}

	// cache. any need?

	repo := repository.New(a.conn)
	uc := usecase.New(repo)
	c := controller.New(uc)
	// TODO: Add middlewares.

	// routing
	api.SetRoutes(a.engine, a.conn, c)

	if err = a.engine.Run("0.0.0.0:3000"); err != nil {
		return fmt.Errorf("engine.Run: %w", err)
	}

	return nil
}

func (a *app) Shutdown() error {
	return database.Shutdown(a.conn)
}
