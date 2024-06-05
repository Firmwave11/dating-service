package token

import (
	modelToken "dating-service/models/token"
	modelUser "dating-service/models/user"
)

func (t *tokenRepo) GetUserAuthFromTokenRequest(authToken string) (userToken modelToken.UserToken, err error) {

	query := "SELECT token_id, user_account_id, auth_token, generated_at, expires_at FROM public.authentication_tokens WHERE auth_token = $1"

	err = t.db.QueryRowx(query, authToken).StructScan(&userToken)

	return userToken, err
}

func (t *tokenRepo) GetUserAccountByEmail(email string) (user modelUser.UserAccount, err error) {
	query := `SELECT id, email, "password", salt, "name", phone, created_at, updated_at, is_deleted FROM public.user_account where email = $1`

	err = t.db.QueryRowx(query, email).StructScan(&user)

	return user, err
}
