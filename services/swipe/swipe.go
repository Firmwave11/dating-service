package swipe

import (
	"context"
	modelSwipe "dating-service/models/swipe"
	"dating-service/utils/response"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *swipeService) GetSwipeProfile(ctx context.Context, userId uuid.UUID) response.Output {
	isUnlimited := false
	profile, err := s.userRepo.GetProfileByUserId(userId.String())
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000006", "Failed Get Profile", "Failed Get Profile", err)
	}

	// searching premium profile
	if profile.IsPremium {
		purchaseHistories, err := s.purchaseRepo.FindPurhaseHistoriesByProfile(int(profile.ID))
		if err != nil {
			return response.Errors(ctx, http.StatusBadRequest, "000006", "Failed Get Purchase Histories", "Failed Get Purchase Histories", err)
		}

		for _, purchase := range purchaseHistories {
			if purchase.PremiumPackageName == "Unlimited Swipe" {
				isUnlimited = true
				break
			}
		}
	}

	if !isUnlimited {
		count, err := s.swipeRepo.CountSwipeToday(int(profile.ID))
		if err != nil {
			return response.Errors(ctx, http.StatusBadRequest, "000006", "Failed Get Count Swipe", "Failed Get Count Swipe", err)
		}
		fmt.Println(count)
		if count >= 10 {
			return response.Errors(ctx, http.StatusBadRequest, "000006", "Failed Get Swipe", "Your Accout has reached limit 1 day 10 swipe", errors.New("Swipe Reach Limit"))
		}
	}

	profileSwipe, err := s.userRepo.GetProfileSwipe(int(profile.ID), profile.Gender)
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000006", "Failed Get Profile Swipe", "Failed Get ProfileSwipe", err)
	}

	return response.Success(ctx, http.StatusOK, "000006", "Success Get Profile Swipe", profileSwipe)
}

func (s *swipeService) PostActionSwipe(ctx context.Context, req modelSwipe.SwipeRequest) response.Output {
	isUnlimited := false
	profile, err := s.userRepo.GetProfileByUserId(req.UserId.String())
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000007", "Failed Get Profile", "Failed Get Profile", err)
	}

	// searching premium profile
	if profile.IsPremium {
		purchaseHistories, err := s.purchaseRepo.FindPurhaseHistoriesByProfile(int(profile.ID))
		if err != nil {
			return response.Errors(ctx, http.StatusBadRequest, "000007", "Failed insert Swipe", "Failed Get Purchase Histories", err)
		}

		for _, purchase := range purchaseHistories {
			if purchase.PremiumPackageName == "Unlimited Swipe" {
				isUnlimited = true
				break
			}
		}
	}

	if !isUnlimited {
		count, err := s.swipeRepo.CountSwipeToday(int(profile.ID))
		if err != nil {
			return response.Errors(ctx, http.StatusBadRequest, "000007", "Failed insert Swipe", "Failed Get Count Swipe", err)
		}
		if count >= 10 {
			return response.Errors(ctx, http.StatusBadRequest, "000007", "Failed insert Swipe", "Your Accout has reached limit 1 day 10 swipe", errors.New("Swipe Reach Limit"))
		}
	}

	modelSwipe := modelSwipe.Swipes{
		SenderProfileID:   profile.ID,
		ReceiverProfileID: req.ReceiverProfileID,
		ActionTypeID:      req.ActionTypeID,
		ProfileName:       profile.FirstName + " " + profile.LastName,
	}

	err = s.swipeRepo.InsertSwipe(&modelSwipe)

	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000007", "Failed insert Swipe", "Failed insert Swipe", err)
	}

	return response.Success(ctx, http.StatusCreated, "000007", "Succesfully Swipe Profile", nil)
}
