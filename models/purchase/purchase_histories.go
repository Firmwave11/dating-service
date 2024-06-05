package purchase

import "time"

// PurchaseHistories --
type PurchaseHistories struct {
	ID               int64      `db:"id"`
	ProfileId        int64      `db:"profile_id"`
	PremiumPackageID int64      `db:"premium_package_id"`
	CreatedAT        *time.Time `db:"created_at"`
	CreatedBY        string     `db:"created_by"`
	UpdatedAT        *time.Time `db:"updated_at"`
	UpdatedBY        string     `db:"updated_by"`
	DeletedAT        *time.Time `db:"deleted_at"`
	DeletedBY        *string    `db:"deleted_by"`
	IsDeleted        bool       `db:"is_deleted"`
}

type PurchaseHistoriesJoin struct {
	ID                 int64      `db:"id"`
	ProfileId          int64      `db:"profile_id"`
	PremiumPackageID   int64      `db:"premium_package_id"`
	CreatedAT          *time.Time `db:"created_at"`
	CreatedBY          string     `db:"created_by"`
	UpdatedAT          *time.Time `db:"updated_at"`
	UpdatedBY          string     `db:"updated_by"`
	DeletedAT          *time.Time `db:"deleted_at"`
	DeletedBY          *string    `db:"deleted_by"`
	IsDeleted          bool       `db:"is_deleted"`
	PremiumPackageName string     `db:"name"`
}
