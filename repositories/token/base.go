package token

import (
	model "dating-service/models/token"
	modelUser "dating-service/models/user"

	"github.com/jmoiron/sqlx"
)

type tokenRepo struct {
	db *sqlx.DB
}

type TokenRepo interface {
	GetUserAuthFromTokenRequest(authToken string) (userToken model.UserToken, err error)
	GetUserAccountByEmail(email string) (user modelUser.UserAccount, err error)
	InsertUserToken(userToken *model.UserToken) (*model.UserToken, error)
}

func NewTokenRepository(db *sqlx.DB) TokenRepo {
	return &tokenRepo{
		db: db,
	}
}
