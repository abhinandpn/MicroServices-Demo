package interfaces

import (
	"context"

	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/model"
)

type UserUsecase interface {
	Login(ctx context.Context, user model.User) (model.User, error)
	SignUp(ctx context.Context, user model.User) error
	GetUserInfo(ctx context.Context,id uint) (model.User, error)
	// Update()()
}
