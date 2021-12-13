package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gimmickless/keat-kit-service/internal/app"
	egdb "github.com/gimmickless/keat-kit-service/internal/transport/egress/db"
	inhttp "github.com/gimmickless/keat-kit-service/internal/transport/ingress/http"
	"github.com/gimmickless/keat-kit-service/pkg/custom"
	applog "github.com/gimmickless/keat-kit-service/pkg/logging"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/etag"
	httplogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/joho/godotenv"
)

var (
	port = os.Getenv("HTTP_PORT")
)

func main() {
	logger := applog.NewLogger()

	// Load .env file (only expected on local workstations)
	if err := godotenv.Load(".env"); err != nil {
		logger.Debugf(".env file could not be loaded (only harmful when running as standalone on local workstations).")
	}

	// Initialize the db connection
	db, cancel, disconnect := initdb()
	defer cancel()
	defer disconnect()

	// Initialize languages
	initI18n()

	// Init and bind projects
	var (
		catgRepo   = egdb.NewCategoryRepository(logger, db)
		ingredRepo = egdb.NewIngredientRepository(logger, db)
		kitsRepo   = egdb.NewKitRepository(logger, db)
		catgSrv    = app.NewCategoryService(logger, catgRepo)
		ingredSrv  = app.NewIngredientService(logger, ingredRepo)
		kitSrv     = app.NewKitService(logger, kitsRepo)
		handler    = inhttp.NewHTTPHandler(logger, catgSrv, ingredSrv, kitSrv)
	)

	// Init Fiber web framework and attach middlewares
	app := fiber.New(
		fiber.Config{
			ErrorHandler: custom.CreateCustomHTTPErrorHandler(),
		},
	)
	app.Use(etag.New())
	app.Use(recover.New())
	app.Use(httplogger.New())
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	inhttp.Register(app, handler)

	// Start server in a separate goroutine
	go func() {
		if err := app.Listen(fmt.Sprintf(":%s", port)); err != nil && err != http.ErrServerClosed {
			logger.Fatalw("shutting down the server:", "error", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	if err := app.Shutdown(); err != nil {
		logger.Fatal(err)
	}
}
