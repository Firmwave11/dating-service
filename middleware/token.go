package middleware

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

type Key int

const (
	bearerToken = "Bearer"
	CtxUserId   = Key(10)
)

func (c *client) CheckToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		var (
			ctx           = r.Context()
			authorization = r.Header.Get("Authorization")
			err           error
		)

		if reflect.ValueOf(authorization).IsZero() {
			err = fmt.Errorf("authorization header not found")
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidToken), message[InvalidToken], reason[InvalidToken], err)
			return
		}

		token := strings.Split(authorization, " ")
		if token[0] != bearerToken {
			err = fmt.Errorf("bearer token salah. %s", authorization)
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidBearer), message[InvalidBearer], reason[InvalidBearer], err)
			return
		}

		res, err := c.token.ValidateToken(ctx, token[1])
		if err != nil {
			c.errorResponse(ctx, rw, http.StatusUnauthorized, string(InvalidToken), message[InvalidToken], reason[InvalidToken], err)
			return
		}

		// store to context
		ctx = context.WithValue(
			ctx,
			CtxUserId,
			res.UserAccountID,
		)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}
