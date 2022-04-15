package main

import (
	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"jwt-server/database"
	"jwt-server/routes"
	"jwt-server/utils"
)

func main() {
	if err := database.InitDB(); err != nil {
		panic(err)
	}

	app := fiber.New()

	setupRoutes(app)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

func setupRoutes(app *fiber.App) {
	app.Post("/signup", routes.Signup)

	app.Post("/login", routes.Login)

	reissued := app.Group("/reissued")
	reissued.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(utils.RefreshTokenSecretKey),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Error(c, 500, "리프레시 토큰이 만료되었거나 정상적인 리프레시 토큰이 아닙니다.")
		},
	}))
	reissued.Post("/", routes.Reissued)

	access := app.Group("/access")
	access.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(utils.AccessTokenSecretKey),
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return utils.Error(c, 500, "엑세스 토큰이 만료되었거나 정상적인 엑세스 토큰이 아닙니다.")
		},
	}))
	access.Get("/", routes.Access)
}
