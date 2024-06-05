package purchase

import "github.com/google/uuid"

type PurchaseRequest struct {
	PremiumPackageID int64     `json:"premium_package_id"`
	UserID           uuid.UUID `json:"user_id"`
}
