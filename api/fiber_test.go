package api_test

import (
	"bytes"
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"testing"

	"github.com/fkrhykal/quickbid-account/api"
	"github.com/fkrhykal/quickbid-account/api/handler"
	"github.com/fkrhykal/quickbid-account/api/route"
	"github.com/fkrhykal/quickbid-account/db"
	"github.com/fkrhykal/quickbid-account/db/persistence"
	"github.com/fkrhykal/quickbid-account/internal/credential"
	"github.com/fkrhykal/quickbid-account/internal/entity"
	"github.com/fkrhykal/quickbid-account/internal/service"
	"github.com/fkrhykal/quickbid-account/internal/validation"
	"github.com/go-faker/faker/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSignUp(t *testing.T) {
	log := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))

	app := api.NewFiberApp(log)

	pgDB, err := db.SetupPostgresDB(db.PostgresTestConfig)
	assert.NoError(t, err)

	defer func() {
		pgDB.Exec("TRUNCATE TABLE users")
		pgDB.Close()
	}()

	execManager := db.NewSqlExecutorManager(pgDB)
	saveUser := persistence.PgSaveUser(log)
	findUserByUsername := persistence.PgFindUserByUsername(log)
	passwordManager := credential.NewBcryptPasswordManager(log)

	route.SignUpRoute(app, handler.SignUpHandler(
		log,
		service.SignUpService(
			log,
			validation.ValidateSignUpRequest,
			execManager,
			saveUser,
			findUserByUsername,
			passwordManager,
		),
	))

	t.Run("request success", func(t *testing.T) {
		requestBody, err := json.Marshal(fiber.Map{
			"username": "sdnsid",
			"password": "ncsdfnc&8_A1",
		})
		assert.NoError(t, err)

		httpRequest, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		httpRequest.Header.Add("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)

		assert.Equal(t, httpResponse.StatusCode, fiber.StatusCreated)
	})

	t.Run("request body type mismatch", func(t *testing.T) {
		requestBody, err := json.Marshal(fiber.Map{
			"username": "sdnsidne1",
			"password": 4,
		})
		assert.NoError(t, err)

		httpRequest, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		httpRequest.Header.Add("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)

		assert.Equal(t, httpResponse.StatusCode, fiber.StatusUnprocessableEntity)
	})

	t.Run("request body invalid", func(t *testing.T) {
		requestBody, err := json.Marshal(fiber.Map{
			"username": "sdnsidne1",
			"password": "sds",
		})
		assert.NoError(t, err)

		httpRequest, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		httpRequest.Header.Add("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)

		assert.Equal(t, httpResponse.StatusCode, fiber.StatusBadRequest)
	})

	t.Run("username used", func(t *testing.T) {
		user := &entity.User{
			ID:       uuid.New(),
			Username: faker.Username(),
			Password: faker.Password(),
		}
		saveUser(context.Background(), execManager.Executor(), user)

		requestBody, err := json.Marshal(fiber.Map{
			"username": user.Username,
			"password": "ncsdfnc&8_A1",
		})
		assert.NoError(t, err)

		httpRequest, err := http.NewRequest(http.MethodPost, "/sign-up", bytes.NewBuffer(requestBody))
		assert.NoError(t, err)

		httpRequest.Header.Add("Content-Type", "application/json")

		httpResponse, err := app.Test(httpRequest)
		assert.NoError(t, err)

		assert.Equal(t, httpResponse.StatusCode, fiber.StatusConflict)
	})
}
