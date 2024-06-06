package purchase

import (
	"context"
	"database/sql"
	purchaseModel "dating-service/models/purchase"
	"dating-service/utils/response"
	"errors"
	"net/http"
)

func (p *purchaseService) ListAllPremiumPackages(ctx context.Context) response.Output {
	premiumPackages, err := p.purchaseRepo.FindAllPremiumPackages()
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000003", "Failed Get All PremiumPackage", "Failed Get All PremiumPackage", err)
	}

	return response.Success(ctx, http.StatusOK, "000003", "Success Get All PremiumPackage", premiumPackages)
}

func (p *purchaseService) PurchasePackage(ctx context.Context, req purchaseModel.PurchaseRequest) response.Output {
	tx := p.db.MustBegin()
	isVerified := false

	profile, err := p.userRepo.GetProfileByUserId(req.UserID.String())
	if err != nil {
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Purchase Package", "Failed Get UserId", err)
	}

	purchaseHistories, err := p.purchaseRepo.FindPurhaseHistoriesByProfile(int(profile.ID))
	if err != nil && err != sql.ErrNoRows {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Purchase Package", "Failed Insert PremiumPackage", err)
	}

	if len(purchaseHistories) > 0 {
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Purchase Package", "You are aleready have premium packages", errors.New("Already have premium packages"))
	}

	modelPurchaseHistories := purchaseModel.PurchaseHistories{
		ProfileId:        profile.ID,
		PremiumPackageID: req.PremiumPackageID,
	}

	premiumPackages, err := p.purchaseRepo.InsertPurchaseHistories(tx, &modelPurchaseHistories)
	if err != nil {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Purchase Package", "Failed Insert PremiumPackage", err)
	}

	premiumPackage, err := p.purchaseRepo.FindPremiumPackagesById(int(req.PremiumPackageID))
	if err != nil {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Purchase Package", "Failed Get PremiumPackage", err)
	}

	if premiumPackage.Name == "Verified Account" {
		isVerified = true
	}

	err = p.userRepo.UpdatePremiumProfile(tx, profile.ID, isVerified)
	if err != nil {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000004", "Failed Purchase Package", "Failed Update Profile", err)
	}

	err = tx.Commit()
	return response.Success(ctx, http.StatusOK, "000004", "Succesfully Update PremiumPackage", premiumPackages)
}
