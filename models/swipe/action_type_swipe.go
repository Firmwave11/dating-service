package swipe

import "time"

// ActionTypeSwipe --
type ActionTypeSwipe struct {
	ID          int64      `db:"id"`
	Action      string     `db:"action"`
	CreatedAT   *time.Time `db:"created_at"`
	CreatedBY   string     `db:"created_by"`
	UpdatedAT   *time.Time `db:"updated_at"`
	UpdatedBY   string     `db:"updated_by"`
	DeletedAT   *time.Time `db:"deleted_at"`
	DeletedBY   string     `db:"deleted_by"`
	IsDeleted   bool       `db:"is_deleted"`
	Description string     `db:"description"`
}
