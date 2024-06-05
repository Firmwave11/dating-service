package swipe

import (
	"context"
	modelSwipe "dating-service/models/swipe"
	"dating-service/repositories/purchase"
	"dating-service/repositories/swipe"
	"dating-service/repositories/user"
	"dating-service/utils/response"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type swipeService struct {
	purchaseRepo purchase.PurchaseRepo
	userRepo     user.UserRepo
	swipeRepo    swipe.SwipeRepo
	db           *sqlx.DB
}

type SwipeService interface {
	GetSwipeProfile(ctx context.Context, userId uuid.UUID) response.Output
	PostActionSwipe(ctx context.Context, req modelSwipe.SwipeRequest) response.Output
}

func NewSwipeService(swipe swipe.SwipeRepo, purchase purchase.PurchaseRepo, user user.UserRepo, db *sqlx.DB) SwipeService {
	return &swipeService{
		purchaseRepo: purchase,
		userRepo:     user,
		swipeRepo:    swipe,
		db:           db,
	}
}
