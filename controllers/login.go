package controllers

import (
	model "dating-service/models/token"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (c *controller) LoginHandler(e echo.Context) error {
	var (
		req      = e.Request()
		ctx      = req.Context()
		reqLogin = new(model.Login)
	)

	if err := e.Bind(&reqLogin); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "LoginHandler",
			"event":   "Bind",
		}).Error(err)
	}

	return c.tokenService.LoginAction(ctx, *reqLogin).Write(e)
}
