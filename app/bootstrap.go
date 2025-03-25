package app

import (
	"database/sql"
	"log/slog"

	"github.com/fkrhykal/quickbid-account/api/handler"
	"github.com/fkrhykal/quickbid-account/api/route"
	"github.com/fkrhykal/quickbid-account/db"
	"github.com/fkrhykal/quickbid-account/db/persistence"
	"github.com/fkrhykal/quickbid-account/internal/service"
	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/gofiber/fiber/v2"
)

type BootstrapConfig struct {
	Fiber  fiber.Router
	DB     *sql.DB
	Logger *slog.Logger
}

func Bootstrap(config *BootstrapConfig) {
	execManager := db.NewSqlExecutorManager(config.DB)

	signUpService := service.SignUpService(
		config.Logger.WithGroup("SignUpService"),
		validation.ValidateSignUpRequest,
		execManager,
		persistence.PgSaveUser(config.Logger.WithGroup("PgSaveUser")),
		persistence.PgFindUserByUsername(config.Logger.WithGroup("PgFindUserByUsername")),
	)

	signUpHandler := handler.SignUpHandler(
		config.Logger.WithGroup("SignUpHandler"),
		signUpService,
	)

	route.SignUpRoute(config.Fiber, signUpHandler)
}
