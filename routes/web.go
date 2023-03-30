package routes

import (
	"github.com/chenyuIT/framework/contracts/http"
	"github.com/chenyuIT/framework/facades"

	"chenyuIT/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(http.StatusOK, http.Json{
			"Hello": "Goravel",
		})
	})

	userController := controllers.NewUserController()
	facades.Route.Get("/users/{id}", userController.Show)
}
