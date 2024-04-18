package routes

import (
	"student-management/database"
	"student-management/models"
	"github.com/gofiber/fiber/v2"

)



// // func GetallMarks(c *fiber.Ctx, id uint) error {
// // 	StudentId := id;
// // 	marks := []models.Mark{}
// // 	err := database.Database.Db.Where("studentId = ?", StudentId).Error
// // 	if err != nil {
// // 		return c.Status(500).JSON("Error fetching users")
// // 	}
// // 	return c.Status(200).JSON(marks)
// // }

// func GetStudent(c *fiber.Ctx) error {
// 	StudentId, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.Status(400).JSON("Error parsing JSON: " + err.Error())
// 	}
// 	student := models.Student{}
// 	err = database.Database.Db.Preload("marks").Where("id = ?", StudentId).First(&student).Error
// 	if err != nil {
// 		return c.Status(404).JSON("product does not exist")
// 	}
// 	// marks := []models.Mark{}
// 	// err = database.Database.Db.Where("studentId = ?", StudentId).Error
// 	// if err != nil {
// 	// 	return c.Status(500).JSON("Error fetching marks"+err.Error())
// 	// }
// 	studentDetails := fiber.Map{
// 		"id":           student.ID,
//         "regno":        student.Regno,
//         "password":     student.Password,
//         "student_name": student.StudentName,
//         "marks":        student.Marks,
// 	}

// 	return c.Status(200).JSON(studentDetails)
// }


func GetStudent(c *fiber.Ctx) error {
    var student models.Student
    regno := c.Params("regno")

    if err := database.Database.Db.Preload("Marks").Where("regno = ?", regno).First(&student).Error; err != nil {
        return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
            "error": "Student not found",
        })
    }

    return c.JSON(student)
}