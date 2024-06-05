package swipe

import "github.com/google/uuid"

type SwipeRequest struct {
	UserId            uuid.UUID
	ReceiverProfileID int64 `json:"receiver_profile_id"`
	ActionTypeID      int64 `json:"action_type_id"`
}
