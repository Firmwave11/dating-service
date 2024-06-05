package user

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func (u *userRepo) UpdatePremiumProfile(tx *sqlx.Tx, profileId int64, isVerified bool) error {
	query := `UPDATE public.profiles SET is_premium=true, is_verified=$1, updated_at=now() WHERE id=$2`

	stmt, err := tx.Preparex(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		isVerified,
		profileId,
	)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected != 1 {
		err = fmt.Errorf("failed update data event_registration. rows affected %d", rowsAffected)
		return err
	}

	return nil
}
