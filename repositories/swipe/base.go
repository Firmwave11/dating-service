package swipe

import (
	"dating-service/models/swipe"
	swipeModel "dating-service/models/swipe"

	"github.com/jmoiron/sqlx"
)

type swipeRepo struct {
	db *sqlx.DB
}

type SwipeRepo interface {
	CountSwipeToday(profileId int) (count int, err error)
	ActionTypeSwipe(action string) (ActionTypeSwipe swipeModel.ActionTypeSwipe, err error)
	InsertSwipe(swipe *swipe.Swipes) error
}

func NewTokenRepository(db *sqlx.DB) SwipeRepo {
	return &swipeRepo{
		db: db,
	}
}
