package routes

import (
	c "go-fiber-work/controllers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		middleware := basicauth.New(basicauth.Config{
			Users: map[string]string{
				"gofiber": "21022566",
			},
		})
		v1 := api.Group("/v1")
		{
			v1.Post("/mock", middleware , c.CreateMockData)
			dog := v1.Group("/dog")
			{
				dog.Get("/", middleware, c.GetDogs)
				dog.Get("/filter", middleware, c.GetDog)
				dog.Get("/json", middleware, c.GetDogsJson)
				dog.Get("/deleted", middleware, c.GetDeletedDog)
				dog.Get("/get_by_range", middleware , c.GetDogsByRange)
				dog.Post("/", middleware, c.AddDog)
				dog.Put("/:id", middleware, c.UpdateDog)
				dog.Delete("/:id", middleware, c.RemoveDog)
			}

			company := v1.Group("/company")
			{
				company.Get("/", middleware , c.GetCompanies)
				company.Get("/:id", middleware , c.GetCompany)
				company.Post("/", middleware , c.AddCompany)
				company.Put("/:id", middleware , c.UpdateCompany)
				company.Delete(":id",middleware , c.RemoveCompany)

			}

		}
	}
}
