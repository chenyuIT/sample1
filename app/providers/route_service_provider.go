package providers

import (
	"github.com/chenyuIT/framework/facades"

	"chenyuIT/app/http"
	"chenyuIT/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register() {
	//Add HTTP middlewares
	facades.Route.GlobalMiddleware(http.Kernel{}.Middleware()...)
}

func (receiver *RouteServiceProvider) Boot() {
	receiver.configureRateLimiting()

	routes.Web()
}

func (receiver *RouteServiceProvider) configureRateLimiting() {

}
