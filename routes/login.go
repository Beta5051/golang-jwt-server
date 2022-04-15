package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"jwt-server/database"
	"jwt-server/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Login(c *fiber.Ctx) error {
	req := new(LoginRequest)

	if err := c.BodyParser(req); err != nil || (req.Username == "" || req.Password == "") {
		return utils.Error(c, 300, "요청 인자가 비어있습니다.")
	}

	user := new(database.User)
	if has, err := database.DB.Where("username = ?", req.Username).Get(user); !has || err != nil {
		return utils.Error(c, 400, "해당 유저가 존재하지 않습니다.")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return utils.Error(c, 400, "비밀번호가 일치하지 않습니다.")
	}

	accessToken, accessTokenExp, err := utils.GenerateAccessToken(user.Id)
	if err != nil {
		return utils.Error(c, 500, "엑세스 토큰 생성중 오류가 발생했습니다.")
	}

	refreshToken, refreshTokenExp, err := utils.GenerateRefreshToken(user.Id)
	if err != nil {
		return utils.Error(c, 500, "리프레시 토큰 생성중 오류가 발생했습니다.")
	}

	return utils.Result(c, fiber.Map{
		"message":           "로그인 성공!",
		"access_token":      accessToken,
		"access_token_exp":  accessTokenExp,
		"refresh_token":     refreshToken,
		"refresh_token_exp": refreshTokenExp,
	})
}
