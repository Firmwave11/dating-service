package user

import (
	"dating-service/models/user"
	model "dating-service/models/user"

	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

type UserRepo interface {
	InsertUserAccount(tx *sqlx.Tx, user *model.UserAccount) (*model.UserAccount, error)
	InsertProfile(tx *sqlx.Tx, profiles *model.Profiles) (*model.Profiles, error)
	GetProfileByUserId(userId string) (profile user.Profiles, err error)
	UpdatePremiumProfile(tx *sqlx.Tx, profileId int64, isVerified bool) error
	GetProfileSwipe(profileId int, gender string) (profile user.Profiles, err error)
}

func NewUserRepository(db *sqlx.DB) UserRepo {
	return &userRepo{
		db: db,
	}
}
