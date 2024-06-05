package purchase

import (
	"dating-service/models/purchase"

	"github.com/jmoiron/sqlx"
)

func (p *purchaseRepo) InsertPurchaseHistories(tx *sqlx.Tx, premiumPackages *purchase.PurchaseHistories) (*purchase.PurchaseHistories, error) {
	query := `
	INSERT INTO public.purchase_histories (profile_id, premium_package_id, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted) VALUES($1, $2, now(), $3, now(), $4, null, null, false)
	returning id
	`

	err := tx.QueryRow(
		query,
		premiumPackages.ProfileId,
		premiumPackages.PremiumPackageID,
		"System",
		"System",
	).Scan(&premiumPackages.ID)

	return premiumPackages, err
}
