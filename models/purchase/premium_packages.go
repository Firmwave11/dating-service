package purchase

import "time"

// PremiumPackages --
type PremiumPackages struct {
	ID          int64      `db:"id"`
	Name        string     `db:"name"`
	Description string     `db:"description"`
	Price       float64    `db:"price"`
	CreatedAT   *time.Time `db:"created_at"`
	CreatedBY   string     `db:"created_by"`
	UpdatedAT   *time.Time `db:"updated_at"`
	UpdatedBY   string     `db:"updated_by"`
	DeletedAT   *time.Time `db:"deleted_at"`
	DeletedBY   *string    `db:"deleted_by"`
	IsDeleted   bool       `db:"is_deleted"`
}
