package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"jwt-server/database"
	"jwt-server/utils"
)

func Access(c *fiber.Ctx) error {
	userId := utils.GetUserIdFromToken(c.Locals("user").(*jwt.Token))

	user := new(database.User)
	if has, err := database.DB.Where("id = ?", userId).Get(user); !has || err != nil {
		return utils.Error(c, 400, "해당 유저가 존재하지 않습니다.")
	}

	return utils.Result(c, fiber.Map{
		"user": user,
	})
}
