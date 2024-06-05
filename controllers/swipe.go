package controllers

import (
	"dating-service/middleware"
	"dating-service/models/swipe"

	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func (c *controller) GetSwipeProfileHandler(e echo.Context) error {
	var (
		req  = e.Request()
		ctx  = req.Context()
		user uuid.UUID
	)

	if userId, ok := ctx.Value(middleware.CtxUserId).(uuid.UUID); ok {
		user = userId
	}
	return c.swipeService.GetSwipeProfile(ctx, user).Write(e)
}

func (c *controller) PostSwipeProfileHandler(e echo.Context) error {
	var (
		req           = e.Request()
		ctx           = req.Context()
		reqPostAction = new(swipe.SwipeRequest)
	)

	if userId, ok := ctx.Value(middleware.CtxUserId).(uuid.UUID); ok {
		reqPostAction.UserId = userId
	}

	if err := e.Bind(&reqPostAction); err != nil {
		logrus.WithFields(logrus.Fields{
			"service": "CreateRegistrationHandler",
			"event":   "Bind",
		}).Error(err)
	}
	return c.swipeService.PostActionSwipe(ctx, *reqPostAction).Write(e)
}
