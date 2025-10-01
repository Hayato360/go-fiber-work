package routes

import (
	c "go-fiber-work/controllers"
	"go/doc"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	{
		middleware := basicauth.New(basicauth.Config{
			Users: map[string]string{
				"admin": "123456",
			},
		})
		v1 := api.Group("/v1")
		{
			dog := v1.Group("/dog")
			{
				dog.Get("/", middleware, c.GetDogs)
				dog.Get("/filter", middleware, c.GetDog)
				dog.Get("/json", middleware, c.GetDogsJson)
				dog.Get("/deleted", middleware, c.GetDeletedDog)
				dog.Post("/", middleware, c.AddDog)
				dog.Put("/:id", middleware, c.UpdateDog)
				dog.Delete("/:id", middleware, c.RemoveDog)
			}

			company := v1.Group("/company")
			{
				company.Get("")
			}

		}
	}
}
