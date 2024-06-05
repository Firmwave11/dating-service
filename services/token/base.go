package token

import (
	"context"
	model "dating-service/models/token"
	tokenRepo "dating-service/repositories/token"
	"dating-service/utils/response"

	"github.com/jmoiron/sqlx"
)

type tokenService struct {
	tokenRepo tokenRepo.TokenRepo
	db        *sqlx.DB
}

type TokenService interface {
	ValidateToken(ctx context.Context, auth_token string) (*model.UserToken, error)
	GenerateToken(ctx context.Context, req model.Login) (*model.UserToken, error)
	LoginAction(ctx context.Context, req model.Login) response.Output
}

func NewTokenService(e tokenRepo.TokenRepo, db *sqlx.DB) TokenService {
	return &tokenService{
		tokenRepo: e,
		db:        db,
	}
}
