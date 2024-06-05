package swipe

import (
	swipeModel "dating-service/models/swipe"
)

func (t *swipeRepo) CountSwipeToday(profileId int) (count int, err error) {
	query := `SELECT Count(*) FROM public.swipes where sender_profile_id = $1 and DATE(Created_at) >= CURRENT_DATE AND DATE(Created_at) < CURRENT_DATE + INTERVAL '1 DAY'`

	err = t.db.QueryRowx(query, profileId).Scan(&count)
	return count, err
}

func (t *swipeRepo) ActionTypeSwipe(action string) (ActionTypeSwipe swipeModel.ActionTypeSwipe, err error) {
	query := `SELECT id, "action", created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted, description FROM public.action_type_swipe `

	err = t.db.QueryRowx(query, action).StructScan(&ActionTypeSwipe)

	return ActionTypeSwipe, err
}
