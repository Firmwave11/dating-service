package controllers

import (
	"dating-service/middleware"
	"dating-service/models/purchase"
	"dating-service/utils/response"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func (c *controller) GetAllPremiumPackagesHandler(e echo.Context) error {
	var (
		req = e.Request()
		ctx = req.Context()
	)

	return c.purchaseService.ListAllPremiumPackages(ctx).Write(e)
}

func (c *controller) PurchasePremiumHandler(e echo.Context) error {
	var (
		req              = e.Request()
		ctx              = req.Context()
		reqPurchase      = new(purchase.PurchaseRequest)
		idPremiumPackage = e.Param("id")
	)

	if userId, ok := ctx.Value(middleware.CtxUserId).(uuid.UUID); ok {
		reqPurchase.UserID = userId
	}

	if idPremiumPackage != "" {
		id, err := strconv.Atoi(idPremiumPackage)
		if err != nil {
			return response.Errors(ctx, http.StatusBadRequest, "000004", "Error", "Invalid Request", err).Write(e)
		}
		reqPurchase.PremiumPackageID = int64(id)
	}

	return c.purchaseService.PurchasePackage(ctx, *reqPurchase).Write(e)
}
