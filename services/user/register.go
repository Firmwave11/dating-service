package user

import (
	"context"
	model "dating-service/models/user"
	utils "dating-service/utils/encrypt"
	"dating-service/utils/response"
	"errors"
	"net/http"
	"time"
)

var (
	ErrAccountExist = errors.New("account already exist")
)

func (u *userService) RegisterAction(ctx context.Context, req model.Register) response.Output {
	tx := u.db.MustBegin()
	userByEmail, err := u.tokenRepo.GetUserAccountByEmail(req.Email)
	if userByEmail.Email == req.Email {
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Account already registered", "Failed Register", ErrAccountExist)
	}

	dt := time.Now()

	salt := utils.GenerateSalt()
	userModel := model.UserAccount{
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  utils.Hash(req.Password, salt),
		Salt:      salt,
		Name:      req.FirstName + " " + req.LastName,
		IsDeleted: false,
		CreatedAT: &dt,
		UpdatedAT: &dt,
	}
	user, err := u.userRepo.InsertUserAccount(tx, &userModel)
	if err != nil {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Account", "Failed Register", err)
	}
	profileModel := model.Profiles{
		FirstName:     req.FirstName,
		LastName:      req.LastName,
		Gender:        req.Gender,
		PhoneNumber:   req.Phone,
		Photo:         &req.Photo,
		BirthDate:     req.BirthDate,
		Nickname:      req.Nickname,
		Description:   &req.Description,
		IsDeleted:     false,
		IsPremium:     false,
		UserAccountID: user.ID,
	}

	profile, err := u.userRepo.InsertProfile(tx, &profileModel)
	if err != nil {
		tx.Rollback()
		return response.Errors(ctx, http.StatusBadRequest, "000001", "Failed Create Profile", "Failed Register", err)
	}
	err = tx.Commit()

	return response.Success(ctx, http.StatusCreated, "000001", "Success Register", profile)
}
