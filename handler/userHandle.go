package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/jperezviloria/go-jwt-gorm/database"
	"github.com/jperezviloria/go-jwt-gorm/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func MigrateUser(sql *gorm.DB) {
	sql.AutoMigrate(&model.User{})
	fmt.Println("Todo Entity migrated")
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func GetAllUsers(ctx *fiber.Ctx) error {
	db := database.ConnectSqlServer()
	MigrateUser(db)
	var users []model.User
	db.Find(&users)
	return ctx.Status(fiber.StatusOK).JSON(users)

}

func GetUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	db := database.ConnectSqlServer()
	MigrateUser(db)
	var user model.User
	db.Find(&user, id)
	if user.Username == "" {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "No user found with ID",
			"data":    "nil",
		})
	}
	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Product found",
		"data":    user,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	type NewUser struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}

	db := database.ConnectSqlServer()
	MigrateUser(db)
	user := new(model.User)
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Review you input",
			"data":    err,
		})
	}

	hash, err := hashPassword(user.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't hash password",
			"data":    err,
		})
	}

	user.Password = hash
	if err := db.Create(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Couldn't create user",
			"data":    err,
		})
	}
	newUser := NewUser{
		Email:    user.Email,
		Username: user.Username,
	}

	return ctx.JSON(fiber.Map{
		"status":  "success",
		"message": "Created user",
		"data":    newUser,
	})
}
