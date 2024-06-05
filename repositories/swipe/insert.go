package swipe

import (
	"dating-service/models/swipe"
)

func (s *swipeRepo) InsertSwipe(swipe *swipe.Swipes) error {
	query := `
	INSERT INTO public.swipes (sender_profile_id, receiver_profile_id, action_type_id, created_at, created_by, updated_at, updated_by) VALUES($1, $2, $3, now(), $4, now(),$4)
	returning id
	`

	err := s.db.DB.QueryRow(
		query,
		swipe.SenderProfileID,
		swipe.ReceiverProfileID,
		swipe.ActionTypeID,
		swipe.ProfileName,
	).Scan(&swipe.ID)

	return err
}
