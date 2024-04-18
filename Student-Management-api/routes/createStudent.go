package routes

import (
	"student-management/database"
	"student-management/models"
	"github.com/gofiber/fiber/v2"
	"log"
)

func CreateStudent(c *fiber.Ctx) error {
	var student models.Student
	err := c.BodyParser(&student)
	if err != nil {
		return c.Status(400).JSON("Error parsing JSON: " + err.Error())

	}

	existingStudent := models.Student{}
	if database.Database.Db.Where("regno = ?", student.Regno).First(&existingStudent).Error == nil {
		return c.Status(200).JSON("student reg number exists")
	}

	result := database.Database.Db.Create(&student)
	if result.Error != nil {
		return c.Status(500).SendString("Error creating student")

	}

	log.Printf("student successfully created. ID: \n")
	return c.SendString("Created student")

}