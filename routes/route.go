package routes

import (
	"dating-service/controllers"
	"dating-service/middleware"

	"github.com/labstack/echo"
)

type route struct {
	middleware middleware.Clients
	controller controllers.Controller
}

type Route interface {
	Route() *echo.Echo
}

func NewRouter(mdl middleware.Clients, controller controllers.Controller) Route {
	return &route{
		middleware: mdl,
		controller: controller,
	}
}

func (r *route) Route() *echo.Echo {
	router := echo.New()

	baseRoute := router.Group("", echo.WrapMiddleware(r.middleware.Logger))

	purchase := baseRoute.Group("/purchase", echo.WrapMiddleware(r.middleware.CheckToken))
	swipe := baseRoute.Group("/swipe", echo.WrapMiddleware(r.middleware.CheckToken))

	baseRoute.POST("/login", r.controller.LoginHandler)
	baseRoute.POST("/register", r.controller.RegisterHandler)

	// list endpoint purchase will be here
	purchase.GET("/premium-packages", r.controller.GetAllPremiumPackagesHandler)
	purchase.PUT("/upgrade-profile/:id", r.controller.PurchasePremiumHandler)

	// list endpoint swipe will be here
	swipe.GET("/profile", r.controller.GetSwipeProfileHandler)
	swipe.POST("/profile", r.controller.PostSwipeProfileHandler)

	return router
}
