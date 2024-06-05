package purchase

import (
	"context"
	purchaseModel "dating-service/models/purchase"
	"dating-service/repositories/purchase"
	"dating-service/repositories/user"
	"dating-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type purchaseService struct {
	purchaseRepo purchase.PurchaseRepo
	userRepo     user.UserRepo
	db           *sqlx.DB
}

type PurchaseService interface {
	ListAllPremiumPackages(ctx context.Context) response.Output
	PurchasePackage(ctx context.Context, req purchaseModel.PurchaseRequest) response.Output
}

func NewPurchaseService(purchase purchase.PurchaseRepo, user user.UserRepo, db *sqlx.DB) PurchaseService {
	return &purchaseService{
		purchaseRepo: purchase,
		userRepo:     user,
		db:           db,
	}
}
