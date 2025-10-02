package controllers

import (
	"go-fiber-work/database"
	m "go-fiber-work/models"
	"strings"

	"github.com/go-playground/validator/v10"

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

func GetDeletedDog(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Unscoped().Where("deleted_at IS NOT NULL").Find(&dogs)

	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Deleted dogs retrieved successfully",
		"data":    dogs,
		"count":   len(dogs),
	})
}

func GetDogsJson(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

	db.Find(&dogs)

	var dataResults []m.DogsRes
	sumRed := 0
	sumGreen := 0
	sumPink := 0
	sumNoColor := 0

	for _, v := range dogs {
		typeStr := ""
		if v.DogID >= 10 && v.DogID <= 50 {
			typeStr = "red"
			sumRed++
		} else if v.DogID >= 100 && v.DogID <= 150 {
			typeStr = "green"
			sumGreen++
		} else if v.DogID >= 200 && v.DogID <= 250 {
			typeStr = "pink"
			sumPink++
		} else {
			typeStr = "no color"
			sumNoColor++
		}

		d := m.DogsRes{
			Name:  v.Name,
			DogID: v.DogID,
			Type:  typeStr,
		}
		dataResults = append(dataResults, d)
	}

	return c.Status(200).JSON(fiber.Map{
		"count":       len(dogs),
		"data":        dataResults,
		"name":        "golang-test",
		"sum_red":     sumRed,
		"sum_green":   sumGreen,
		"sum_pink":    sumPink,
		"sum_nocolor": sumNoColor,
	})
}

func GetDogsByRange(c *fiber.Ctx) error {
	db := database.DBConn
	var dogs []m.Dogs

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


func CreateMockData(c *fiber.Ctx) error {
	db := database.DBConn

	mockDogs := []m.Dogs{
		{Name: "Buddy", DogID: 14},
		{Name: "Max", DogID: 27},
		{Name: "Charlie", DogID: 113},
		{Name: "Cooper", DogID: 114},
		{Name: "Rocky", DogID: 214},
		{Name: "Bear", DogID: 222},
		{Name: "Tucker", DogID: 500},
		{Name: "Duke", DogID: 19},
		{Name: "Jack", DogID: 20},
		{Name: "Toby", DogID: 243},
	}

	isActiveTrue := true
	isActiveFalse := false

	mockCompanies := []m.Company{
		{
			Name:        "Tech Solutions Thailand",
			Address:     "123 Silom Road, Bangrak, Bangkok 10500",
			Phone:       "02-234-5678",
			Email:       "info@techsolutions.th",
			Website:     "https://www.techsolutions.th",
			CompanyType: "Technology",
			IsActive:    &isActiveTrue,
		},
		{
			Name:        "Bangkok Manufacturing Co.",
			Address:     "456 Rama IV Road, Klong Toey, Bangkok 10110",
			Phone:       "02-345-6789",
			Email:       "contact@bangkokmfg.com",
			Website:     "https://www.bangkokmfg.com",
			CompanyType: "Manufacturing",
			IsActive:    &isActiveTrue,
		},
		{
			Name:        "Golden Trading Ltd",
			Address:     "789 Sukhumvit Road, Watthana, Bangkok 10110",
			Phone:       "02-456-7890",
			Email:       "sales@goldentrading.co.th",
			Website:     "https://www.goldentrading.co.th",
			CompanyType: "Trading",
			IsActive:    &isActiveTrue,
		},
		{
			Name:        "Green Energy Systems",
			Address:     "321 Phahonyothin Road, Chatuchak, Bangkok 10900",
			Phone:       "02-567-8901",
			Email:       "info@greenenergy.th",
			Website:     "https://www.greenenergy.th",
			CompanyType: "Energy",
			IsActive:    &isActiveTrue,
		},
		{
			Name:        "Creative Design Studio",
			Address:     "654 Ratchadamri Road, Pathumwan, Bangkok 10330",
			Phone:       "02-678-9012",
			Email:       "hello@creativedesign.co.th",
			Website:     "https://www.creativedesign.co.th",
			CompanyType: "Design",
			IsActive:    &isActiveFalse,
		},
		{
			Name:        "Food & Beverage Co.",
			Address:     "987 Lat Phrao Road, Wang Thonglang, Bangkok 10310",
			Phone:       "02-789-0123",
			Email:       "contact@foodbev.th",
			Website:     "https://www.foodbev.th",
			CompanyType: "Food & Beverage",
			IsActive:    &isActiveTrue,
		},
		{
			Name:        "Digital Marketing Agency",
			Address:     "147 Asoke Road, Watthana, Bangkok 10110",
			Phone:       "02-890-1234",
			Email:       "info@digitalmarketing.co.th",
			Website:     "https://www.digitalmarketing.co.th",
			CompanyType: "Marketing",
			IsActive:    &isActiveTrue,
		},
		{
			Name:        "Construction Plus Ltd",
			Address:     "258 Petchburi Road, Ratchathewi, Bangkok 10400",
			Phone:       "02-901-2345",
			Email:       "projects@constructionplus.th",
			Website:     "https://www.constructionplus.th",
			CompanyType: "Construction",
			IsActive:    &isActiveTrue,
		},
		{
			Name:        "Healthcare Solutions",
			Address:     "369 Vibhavadi Road, Chatuchak, Bangkok 10900",
			Phone:       "02-012-3456",
			Email:       "info@healthcaresolutions.th",
			Website:     "https://www.healthcaresolutions.th",
			CompanyType: "Healthcare",
			IsActive:    &isActiveFalse,
		},
		{
			Name:        "Education Excellence Center",
			Address:     "741 Ramkhamhaeng Road, Huai Khwang, Bangkok 10310",
			Phone:       "02-123-4567",
			Email:       "admin@educationexcellence.th",
			Website:     "https://www.educationexcellence.th",
			CompanyType: "Education",
			IsActive:    &isActiveTrue,
		},
	}

	for _, dog := range mockDogs {
		var existingDog m.Dogs
		result := db.Where("dog_id = ?", dog.DogID).First(&existingDog)
		if result.Error != nil {
			db.Create(&dog)
		}
	}

	for _, company := range mockCompanies {
		var existingCompany m.Company
		result := db.Where("name = ?", company.Name).First(&existingCompany)
		if result.Error != nil {
			db.Create(&company)
		}
	}

	return c.Status(201).JSON(fiber.Map{
		"status":  "success",
		"message": "Mock data created successfully",
		"data": fiber.Map{
			"dogs_created":      len(mockDogs),
			"companies_created": len(mockCompanies),
		},
	})
}