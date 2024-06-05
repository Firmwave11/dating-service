package main

import (
	"dating-service/adapter"
	"dating-service/config"
	"dating-service/controllers"
	"dating-service/middleware"
	purchaseRepo "dating-service/repositories/purchase"
	swipeRepo "dating-service/repositories/swipe"
	tokenRepo "dating-service/repositories/token"
	userRepo "dating-service/repositories/user"
	"dating-service/routes"
	purchaseService "dating-service/services/purchase"
	swipeService "dating-service/services/swipe"
	tokenService "dating-service/services/token"
	userService "dating-service/services/user"
	"dating-service/utils/server"
	"strings"

	"github.com/spf13/viper"
)

func init() {
	config.Load()
	adapter.LoadMultipleDatabase()
}

func main() {
	listDb := viper.GetString("datingservice.db.name") // remove this if you won't use database
	dbName := strings.Split(listDb, ",")               // remove this if you won't use database
	db := adapter.DBConnection(dbName[0])              // remove this if you won't use database

	log := middleware.Logrus()
	timezone := middleware.Timezone()

	//list repo will be here
	tokenRepo := tokenRepo.NewTokenRepository(db)
	userRepo := userRepo.NewUserRepository(db)
	purchaseRepo := purchaseRepo.NewTokenRepository(db)
	swipeRepo := swipeRepo.NewTokenRepository(db)

	//list service will be here
	tokenService := tokenService.NewTokenService(tokenRepo, db)
	userService := userService.NewUserService(tokenRepo, userRepo, db)
	purchaseService := purchaseService.NewPurchaseService(purchaseRepo, userRepo, db)
	swipeService := swipeService.NewSwipeService(swipeRepo, purchaseRepo, userRepo, db)

	//list controller will be here
	controller := controllers.NewController(userService, tokenService, purchaseService, swipeService)
	mdl := middleware.NewMiddleware(log, timezone, tokenService)

	router := routes.NewRouter(mdl, controller)

	server.Wrapper(router.Route())

}
