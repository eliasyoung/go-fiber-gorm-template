package main

import (
	"fmt"
	"time"

	"github.com/eliasyoung/fiber-flavor/config"
	"github.com/eliasyoung/fiber-flavor/internal/db"
	"github.com/eliasyoung/fiber-flavor/pkg"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go.uber.org/zap"
)

type application struct {
	config config.Config
	logger *zap.SugaredLogger
	store  *db.Store
}

func (app *application) mount() *fiber.App {
	f := fiber.New()

	f.Use(recover.New())
	f.Use(cors.New())
	f.Use(requestid.New())
	f.Use(logger.New())
	f.Use(limiter.New(limiter.Config{
		KeyGenerator: func(c *fiber.Ctx) string {
			return c.IP()
		},
		Max:        20,
		Expiration: 1 * time.Second,
	}))
	f.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, token")
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		} else {
			return c.Next()
		}
	})

	userApi := f.Group("/user")
	userApi.Post("/signup", app.createUserHandler)
	userApi.Get("/", app.getAllUsersHandler)
	userApi.Get("/:userId", app.getUserByIdHandler)

	f.Get("/", func(c *fiber.Ctx) error {
		app.logger.Info(app.config)

		return c.JSON(pkg.SuccessResponse(app.config))
	})

	return f
}

func (app *application) run(f *fiber.App) error {
	port := fmt.Sprintf(":%s", app.config.Addr)

	return f.Listen(port)
}
