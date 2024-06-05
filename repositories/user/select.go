package user

import (
	"dating-service/models/user"
)

func (t *userRepo) GetProfileByUserId(userId string) (profile user.Profiles, err error) {

	query := "SELECT id, first_name, last_name, gender, birth_date, phone_number, nickname, description, photo, is_premium, updated_at, created_at, is_deleted, user_account_id, is_verified FROM public.profiles where user_account_id = $1 "

	err = t.db.QueryRowx(query, userId).StructScan(&profile)

	return profile, err
}

func (t *userRepo) GetProfileSwipe(profileId int, gender string) (profile user.Profiles, err error) {
	query := `SELECT id, first_name, last_name, gender, birth_date, phone_number, nickname, description, photo, is_premium, updated_at, created_at, is_deleted, user_account_id, is_verified FROM public.profiles where id not in (select receiver_profile_id from swipes where DATE(Created_at) >= CURRENT_DATE AND DATE(Created_at) < CURRENT_DATE + INTERVAL '1 DAY'  and sender_profile_id = $1) and id != $1 and gender != $2 ORDER BY random() LIMIT 1`

	err = t.db.QueryRowx(query, profileId, gender).StructScan(&profile)

	return profile, err
}
