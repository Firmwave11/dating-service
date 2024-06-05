package purchase

import (
	"dating-service/models/purchase"
)

func (p *purchaseRepo) FindAllPremiumPackages() ([]*purchase.PremiumPackages, error) {
	var (
		premiumPackages = make([]*purchase.PremiumPackages, 0)
		err             error
	)
	query := `SELECT id, "Name", description, price, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted FROM public.premium_packages;
	`

	err = p.db.Select(&premiumPackages, query)

	return premiumPackages, err
}

func (p *purchaseRepo) FindPremiumPackagesById(id int) (premiumPackages purchase.PremiumPackages, err error) {
	query := `SELECT id, "Name", description, price, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted FROM public.premium_packages where id = $1
	`

	err = p.db.QueryRowx(query, id).StructScan(&premiumPackages)

	return premiumPackages, err
}

func (p *purchaseRepo) FindPurhaseHistoriesByProfile(profileId int) ([]*purchase.PurchaseHistoriesJoin, error) {
	var (
		purchaseHistories = make([]*purchase.PurchaseHistoriesJoin, 0)
		err               error
	)
	query := `SELECT ph.id, ph.profile_id, ph.premium_package_id, ph.created_at, ph.created_by, ph.updated_at, ph.updated_by, ph.deleted_at, ph.deleted_by, ph.is_deleted, pp."name"
	FROM public.purchase_histories ph
	Left join public.premium_packages pp on ph.premium_package_id  = pp.id   
	where ph.profile_id = $1`

	err = p.db.Select(&purchaseHistories, query, profileId)

	return purchaseHistories, err
}
