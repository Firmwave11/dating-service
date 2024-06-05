package swipe

import "time"

// Swipes --
type Swipes struct {
	ID                int64      `db:"id"`
	SenderProfileID   int64      `db:"sender_profile_id"`
	ReceiverProfileID int64      `db:"receiver_profile_id"`
	ActionTypeID      int64      `db:"action_type_id"`
	CreatedAT         *time.Time `db:"created_at"`
	CreatedBY         string     `db:"created_by"`
	UpdatedAT         *time.Time `db:"updated_at"`
	UpdatedBY         string     `db:"updated_by"`
	DeletedAT         *time.Time `db:"deleted_at"`
	DeletedBY         string     `db:"deleted_by"`
	ISDeleted         bool       `db:"is_deleted"`
	ProfileName       string     `db:"-"`
}
