package interfaces

import (
	"context"

	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/model"
)

type UserReporitory interface {

	//
	SaveUser(ctx context.Context, user model.User) (uint, error)
	UpdateUser(ctx context.Context, Id uint, Userinfo model.User) (model.User, error)
	DeleteUser(ctx context.Context, id uint) error

	//
	GetuserByName(ctx context.Context, name string) (model.User, error)
	GetUserByEmail(ctx context.Context, mail string) (model.User, error)
	GetUserByNumber(ctx context.Context, number string) (model.User, error)
	GetUserById(ctx context.Context, id uint) (model.User, error)
	FindUser(ctx context.Context, user model.User) (model.User, error)
}
