package purchase

import (
	"dating-service/models/purchase"

	"github.com/jmoiron/sqlx"
)

type purchaseRepo struct {
	db *sqlx.DB
}

type PurchaseRepo interface {
	FindAllPremiumPackages() ([]*purchase.PremiumPackages, error)
	InsertPurchaseHistories(tx *sqlx.Tx, premiumPackages *purchase.PurchaseHistories) (*purchase.PurchaseHistories, error)
	FindPurhaseHistoriesByProfile(profileId int) ([]*purchase.PurchaseHistoriesJoin, error)
	FindPremiumPackagesById(id int) (premiumPackages purchase.PremiumPackages, err error)
}

func NewTokenRepository(db *sqlx.DB) PurchaseRepo {
	return &purchaseRepo{
		db: db,
	}
}
