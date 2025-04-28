package app

import (
	"database/sql"
	"log/slog"

	"github.com/fkrhykal/quickbid-account/api/handler"
	"github.com/fkrhykal/quickbid-account/db"
	"github.com/fkrhykal/quickbid-account/db/persistence"
	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/fkrhykal/quickbid-account/internal/service"
	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/gofiber/fiber/v2"
)

type BootstrapConfig struct {
	Fiber      fiber.Router
	DB         *sql.DB
	Logger     *slog.Logger
	Credential *credential.JwtCredentialManagerConfig
}

func Bootstrap(config *BootstrapConfig) {
	execManager := db.NewSqlExecutorManager(config.DB)
	passwordManager := credential.NewBcryptPasswordManager(config.Logger)
	credentialManager := credential.NewJwtCredentialManager(config.Credential)
	saveUser := persistence.PgSaveUser(config.Logger)
	findByUsername := persistence.PgFindUserByUsername(config.Logger)
	findByID := persistence.PgFindUserByID(config.Logger)

	signUpService := service.SignUpService(
		config.Logger,
		validation.ValidateSignUpRequest,
		execManager,
		saveUser,
		findByUsername,
		passwordManager,
	)
	signInService := service.SignInService(
		config.Logger,
		validation.ValidateSignInRequest,
		execManager,
		findByUsername,
		passwordManager,
		credentialManager,
	)
	userProfileService := service.UserProfileService(execManager, findByID)

	handler.SetupSignUp(config.Logger, config.Fiber, signUpService)
	handler.SetupSignIn(config.Logger, config.Fiber, signInService)
	handler.SetupGetCurrentUserProfile(config.Fiber, credentialManager, userProfileService)
}
