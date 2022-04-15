package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"jwt-server/utils"
)

func Reissued(c *fiber.Ctx) error {
	userId := utils.GetUserIdFromToken(c.Locals("user").(*jwt.Token))

	accessToken, accessTokenExp, err := utils.GenerateAccessToken(userId)
	if err != nil {
		return utils.Error(c, 500, "엑세스 토큰 생성중 오류가 발생했습니다.")
	}

	return utils.Result(c, fiber.Map{
		"message":          "엑세스 토큰 재발급 완료!",
		"access_token":     accessToken,
		"access_token_exp": accessTokenExp,
	})
}
