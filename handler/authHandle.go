package handler

import (
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/jperezviloria/go-jwt-gorm/database"
	"github.com/jperezviloria/go-jwt-gorm/middleware"
	"github.com/jperezviloria/go-jwt-gorm/model"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(e string) (*model.User, error) {
	db := database.ConnectSqlServer()
	MigrateUser(db)
	var user model.User
	err := db.Where(&model.User{Email: e}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func GetUserByUsername(e string) (*model.User, error) {
	db := database.ConnectSqlServer()
	MigrateUser(db)
	var user model.User
	err := db.Where(&model.User{Username: e}).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil

}

func Login(ctx *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"`
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	var input LoginInput
	var ud UserData

	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}
	identity := input.Identity
	pass := input.Password
	fmt.Println(identity + " este es el fmt")
	//email, err := getUserByEmail(identity)
	//if err != nil {
	//	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on email", "data": err})
	//}

	user, err := GetUserByUsername(identity)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "data": err})
	}
	fmt.Println(user.Username)
	fmt.Println(user.Password)
	fmt.Println(user.Email)
	fmt.Println(user.ID)

	newPass, _ := HashPassword(pass)
	fmt.Printf("new pass encripted: %s\n", newPass)
	newPass2, _ := HashPassword(pass)
	fmt.Printf("new pass encripted: %s\n", newPass2)

	//if email == nil && user == nil {
	//	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	//}

	//if user == nil {
	//	return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	//}

	if user == nil {
		ud = UserData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}
	}
	//else {
	//	ud = UserData{
	//		ID:       email.ID,
	//		Username: email.Username,
	//		Email:    email.Email,
	//		Password: email.Password,
	//	}
	//}

	fmt.Println(CheckPasswordHash(newPass, ud.Password))
	fmt.Println(!CheckPasswordHash(newPass, ud.Password))
	fmt.Println(ud)
	fmt.Println(pass)
	if CheckPasswordHash(pass, ud.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = ud.Username
	claims["user_id"] = ud.ID
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	t, err := token.SignedString([]byte(middleware.JwtSecret))
	if err != nil {
		return ctx.SendStatus(fiber.StatusInternalServerError)
	}

	return ctx.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
