package token

import (
	"context"
	"database/sql"
	model "dating-service/models/token"
	utils "dating-service/utils/encrypt"
	utilsEncrypt "dating-service/utils/encrypt"
	utilsLog "dating-service/utils/log"
	"dating-service/utils/response"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	"github.com/spf13/viper"
)

var (
	// ErrLoginFail :nodoc:
	ErrLoginFail = errors.New("login fail")
	// ErrUserNotFound :nodoc:
	ErrUserNotFound = errors.New("user not found")
	// ErrAccountDisabled :nodoc:
	ErrTokenNotValid = errors.New("token not valid")
	ErrGenerateToken = errors.New("invalid generate token")
)

func (t *tokenService) GenerateToken(ctx context.Context, req model.Login) (*model.UserToken, error) {

	res, err := t.tokenRepo.GetUserAccountByEmail(req.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Invalid email or password.\r\n")
		}

		return nil, err
	}

	isPasswordVerified := utils.Verify(req.Password, res.Salt, res.Password)

	if !isPasswordVerified {
		return nil, ErrLoginFail
	}

	dt := time.Now()
	expirtyTime := time.Now().Add(time.Minute * 60)
	key := viper.GetString("application.secret")

	payload := model.Payload{
		Type:      "access_token",
		Scope:     "public",
		UserID:    res.ID.String(),
		CreatedAt: dt,
		ExpiredAt: expirtyTime,
	}

	token, err := utilsEncrypt.Encrypt(payload, key)
	if err != nil {
		utilsLog.LogError(err)
		return nil, ErrGenerateToken
	}

	accessToken := base64.StdEncoding.EncodeToString(token)
	modelToken := model.UserToken{}
	modelToken.AuthToken = accessToken
	modelToken.CreatedAt = dt
	modelToken.UserAccountID = res.ID
	modelToken.ExpiredAt = expirtyTime

	resToken, err := t.tokenRepo.InsertUserToken(&modelToken)
	if err != nil {
		return nil, err
	}
	modelToken.TokenID = resToken.TokenID
	return &modelToken, nil

}

func (t *tokenService) ValidateToken(ctx context.Context, auth_token string) (*model.UserToken, error) {
	res, err := t.tokenRepo.GetUserAuthFromTokenRequest(auth_token)
	dt := time.Now()
	if err != nil {

		if err == sql.ErrNoRows {
			return nil, ErrTokenNotValid
		}

		return nil, err
	}

	if res.ExpiredAt.Before(dt) {
		return nil, errors.New("The token is expired.\r\n")
	}

	return &res, nil
}

func (t *tokenService) LoginAction(ctx context.Context, req model.Login) response.Output {
	res, err := t.GenerateToken(ctx, req)
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000002", "Failed Login", "Failed Login", err)
	}

	return response.Success(ctx, http.StatusCreated, "000002", "Success Login", res)
}
