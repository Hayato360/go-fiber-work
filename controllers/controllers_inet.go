package controllers

import (
	"go-fiber-work/database"
	m "go-fiber-work/models"
	"github.com/go-playground/validator/v10"
	"strings"

	"github.com/gofiber/fiber/v2"
)
func AddDog(c *fiber.Ctx) error {

	db := database.DBConn
	var dog m.Dogs

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Create(&dog)
	return c.Status(201).JSON(dog)
}

func GetDogs(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //delelete = null
	return c.Status(200).JSON(dogs)
}

func GetDog(c *fiber.Ctx) error {
	db := database.DBConn
	search := strings.TrimSpace(c.Query("search"))
	var dog []m.Dogs

	result := db.Find(&dog, "dog_id = ?", search)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}
	return c.Status(200).JSON(&dog)
}

func UpdateDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dog m.Dogs
	id := c.Params("id")

	if err := c.BodyParser(&dog); err != nil {
		return c.Status(503).SendString(err.Error())
	}

	db.Where("id = ?", id).Updates(&dog)
	return c.Status(200).JSON(dog)
}

func RemoveDog(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var dog m.Dogs

	result := db.Delete(&dog, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.SendStatus(200)
}

func GetDeletedDog(c *fiber.Ctx) error  {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)

	return  c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted dogs retrieved successfully",
		"data":    dogs,
		"count":   len(dogs),
	})
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs) //10ตัว

	var dataResults []m.DogsRes
	for _, v := range dogs { //1 inet 112 //2 inet1 113
		typeStr := ""
		if v.DogID == 111 {
			typeStr = "red"
		} else if v.DogID == 113 {
			typeStr = "green"
		} else if v.DogID == 999 {
			typeStr = "pink"
		} else {
			typeStr = "no color"
		}

		d := m.DogsRes{
			Name:  v.Name,  //inet
			DogID: v.DogID, //112
			Type:  typeStr, //no color
		}
		dataResults = append(dataResults, d)
		// sumAmount += v.Amount
	}


	r := m.ResultData{
		Data:  dataResults,
		Name:  "golang-test",
		Count: len(dogs), //หาผลรวม,
	}
	return c.Status(200).JSON(r)
}

func GetDogsByRange(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	// ค้นหาสุนัขที่มี dog_id มากกว่า 50 แต่น้อยกว่า 100
	db.Where("dog_id > ? AND dog_id < ?", 50, 100).Find(&dogs)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Dogs with ID between 50-100 retrieved successfully",
		"data":    dogs,
		"count":   len(dogs),
		"filter":  "dog_id > 50 AND dog_id < 100",
	})
}

func AddCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(company); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	if err := db.Create(&company).Error; err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return c.Status(409).JSON(fiber.Map{
				"status":  "error",
				"message": "Company name already exists",
				"error":   err.Error(),
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"status":  "error",
			"message": "Database error occurred",
			"error":   err.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Company created successfully",
		"data":    company,
	})

}

func GetCompanies(c *fiber.Ctx) error {
	db := database.DBConn
	var companies []m.Company
	db.Find(&companies)
	return c.JSON(&companies)
}

func GetCompany(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var company m.Company
	db.Find(&company, id)
	if company.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"status":  "error",
			"message": "Company not found",
		})
	}
	return c.JSON(&company)
}

func UpdateCompany(c *fiber.Ctx) error {
	db := database.DBConn
	var company m.Company
	id := c.Params("id")

	if err := c.BodyParser(&company); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"status":  "error",
			"message": "Cannot parse JSON",
			"error":   err.Error(),
		})
	}

	validate := validator.New()
	if err := validate.Struct(company); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"status":  "error",
			"message": "Validation failed",
			"errors":  err.Error(),
		})
	}

	db.Where("id = ?", id).Updates(&company)
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Company updated successfully",
		"data":    company,
	})
}

func RemoveCompany(c *fiber.Ctx) error {
	db := database.DBConn
	id := c.Params("id")
	var company m.Company

	result := db.Delete(&company, id)

	if result.RowsAffected == 0 {
		return c.SendStatus(404)
	}

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Company deleted successfully",
	})
}


