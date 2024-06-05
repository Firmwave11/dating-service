package user

import (
	model "dating-service/models/user"

	"github.com/jmoiron/sqlx"
)

func (u *userRepo) InsertUserAccount(tx *sqlx.Tx, user *model.UserAccount) (*model.UserAccount, error) {
	query := `
	INSERT INTO public.user_account (id, email, "password", salt, "name", phone, created_at, updated_at, is_deleted) VALUES(gen_random_uuid(), $1, $2, $3, $4, $5, now(), now(), $6)
	returning id
	`

	err := u.db.DB.QueryRow(
		query,
		user.Email,
		user.Password,
		user.Salt,
		user.Name,
		user.Phone,
		user.IsDeleted,
	).Scan(&user.ID)

	return user, err
}

func (u *userRepo) InsertProfile(tx *sqlx.Tx, profiles *model.Profiles) (*model.Profiles, error) {
	query := `
	INSERT INTO public.profiles (first_name, last_name, gender, birth_date, phone_number, nickname, description, photo, is_premium, updated_at, created_at, is_deleted, user_account_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, now(), now(), $10, $11)
	returning id
	`

	err := u.db.DB.QueryRow(
		query,
		profiles.FirstName,
		profiles.LastName,
		profiles.Gender,
		profiles.BirthDate,
		profiles.PhoneNumber,
		profiles.Nickname,
		profiles.Description,
		profiles.Photo,
		profiles.IsPremium,
		profiles.IsDeleted,
		profiles.UserAccountID,
	).Scan(&profiles.ID)

	return profiles, err
}
