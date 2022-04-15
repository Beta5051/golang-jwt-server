package routes

import (
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"jwt-server/database"
	"jwt-server/utils"
)

type SignupRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Signup(c *fiber.Ctx) error {
	req := new(SignupRequest)

	if err := c.BodyParser(req); err != nil || (req.Username == "" || req.Email == "" || req.Password == "") {
		return utils.Error(c, 300, "요청 인자가 비어있습니다.")
	}

	if has, err := database.DB.Where("username = ?", req.Username).Get(new(database.User)); has || err != nil {
		return utils.Error(c, 400, "이미 해당 닉네임이 존재합니다.")
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return utils.Error(c, 500, "비밀번호 해쉬 생성중 오류가 발생했습니다.")
	}

	user := &database.User{
		Username:     req.Username,
		Email:        req.Email,
		PasswordHash: string(passwordHash),
	}
	if _, err := database.DB.Insert(user); err != nil {
		return utils.Error(c, 600, "DB 에서 알수없는 에러가 발생했습니다.")
	}

	return utils.Result(c, fiber.Map{"message": "회원 가입 성공!"})
}
