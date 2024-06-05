package controllers

import (
	"dating-service/services/purchase"
	"dating-service/services/swipe"
	"dating-service/services/token"
	"dating-service/services/user"

	"github.com/labstack/echo"
)

type controller struct {
	userService     user.UserService
	tokenService    token.TokenService
	purchaseService purchase.PurchaseService
	swipeService    swipe.SwipeService
}

type Controller interface {
	LoginHandler(e echo.Context) error
	RegisterHandler(e echo.Context) error
	GetAllPremiumPackagesHandler(e echo.Context) error
	PurchasePremiumHandler(e echo.Context) error
	GetSwipeProfileHandler(e echo.Context) error
	PostSwipeProfileHandler(e echo.Context) error
}

func NewController(user user.UserService, token token.TokenService, purchase purchase.PurchaseService, swipe swipe.SwipeService) Controller {
	return &controller{
		userService:     user,
		tokenService:    token,
		purchaseService: purchase,
		swipeService:    swipe,
	}
}
