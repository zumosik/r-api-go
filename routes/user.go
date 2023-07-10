package routes

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/zumosik/r-api-go/database"
	"github.com/zumosik/r-api-go/models"
)

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DB.Create(&user)

	return c.Status(201).JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.DB.Find(&users)

	return c.Status(200).JSON(users)
}

func FindUser(id int, user *models.User) error {
	database.DB.Find(&user, "id = ?", id)
	if user.ID == 0 {
		return errors.New("User does not exist")

	}

	return nil
}

func GetUserById(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is integer")
	}

	var user models.User

	err = FindUser(id, &user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is integer")
	}

	var user models.User

	err = FindUser(id, &user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateUser struct { // we cant update id or timestamps so we need this structure
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	var updateUser UpdateUser

	if err := c.BodyParser(&updateUser); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	user.Name = updateUser.Name
	user.Age = updateUser.Age

	database.DB.Save(&user)

	return c.Status(200).JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	if err != nil {
		return c.Status(400).JSON("Please ensure that :id is integer")
	}

	var user models.User

	err = FindUser(id, &user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.DB.Delete(&user).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).JSON(user)
}
