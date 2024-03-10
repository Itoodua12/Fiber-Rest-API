package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/itoodua12/Fiber-Rest-API/database"
	"github.com/itoodua12/Fiber-Rest-API/model"
)

// Serializer
type UserSerializer struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user model.User) UserSerializer {
	return UserSerializer{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}
}

func CreateUser(c *fiber.Ctx) error {
	var user model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&user)
	responseUser := CreateResponseUser(user)

	return c.Status(200).JSON(responseUser)
}

func GetUsers(c *fiber.Ctx) error {
	users := []model.User{}

	database.Database.DB.Find(&users)

	responseUsers := []UserSerializer{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}

func findUser(id int, user *model.User) error {
	database.Database.DB.Find(&user, "id=?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}

func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user model.User

	if err != nil {
		return c.Status(400).JSON("Bad Request")
	}
	if err := findUser(id, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseUser := CreateResponseUser(user)
	
	return c.Status(200).JSON(responseUser)

}
