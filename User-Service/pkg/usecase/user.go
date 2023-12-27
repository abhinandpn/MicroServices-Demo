package usecase

import (
	"context"
	"errors"

	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/model"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/repository/interfaces"
	services "github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/usecase/interfaces"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	userRepo interfaces.UserReporitory
}

func NewUserUsecase(repo interfaces.UserReporitory) services.UserUsecase {
	return &UserUsecase{userRepo: repo}
}

func (u *UserUsecase) Login(ctx context.Context, user model.User) (model.User, error) {

	DbUser, DbErr := u.userRepo.FindUser(ctx, user)
	if DbErr != nil {
		return DbUser, DbErr
	} else if DbUser.Id == 0 {
		res := errors.New("user does not exist")
		return DbUser, res
	}

	err := bcrypt.CompareHashAndPassword([]byte(DbUser.Password), []byte(user.Password))
	if err != nil {
		return user, err
	}
	return DbUser, err
}

func (u *UserUsecase) SignUp(ctx context.Context, user model.User) error {

	// find user
	user, err := u.userRepo.GetuserByName(ctx, user.Name)
	if err != nil {
		return err
	}
	if user.Id != 0 {
		user, err = u.userRepo.GetUserByEmail(ctx, user.Email)
		if err != nil {
			return err
		}
		if user.Id != 0 {
			user, err = u.userRepo.GetUserByNumber(ctx, user.Phone)
			if err != nil {
				return err
			}
			if user.Id != 0 {
				res := errors.New("user alredy exist with details")
				return res
			}
		}
	}
	// hash the password
	HashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}
	// save the password
	user.Password = string(HashPass)

	// save the user info
	_, err = u.userRepo.SaveUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserUsecase) GetUserInfo(ctx context.Context, id uint) (model.User, error) {

	user, err := u.userRepo.GetUserById(ctx, id)
	if err != nil {
		return user, err
	}
	return user, nil
}
