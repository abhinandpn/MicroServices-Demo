package repository

import (
	"context"

	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/model"
	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/repository/interfaces"
	"gorm.io/gorm"
)

type UserDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserReporitory {
	return &UserDatabase{DB: DB}
}

func (u *UserDatabase) SaveUser(ctx context.Context, user model.User) (uint, error) {

	query := `insert into users (name,email,phone)values ($1,$2,$3) retirning id;`
	var body model.User

	err := u.DB.Raw(query, user).Scan(&body).Error
	if err != nil {
		return body.Id, err
	}
	return body.Id, nil
}

func (u *UserDatabase) UpdateUser(ctx context.Context,
	Id uint, Userinfo model.User) (model.User, error) {

	query := `update users set name = $1,
		email = $2, 
		phone = $3,
		where id = $4 ;`
	var body model.User

	err := u.DB.Raw(query, Userinfo, Id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (u *UserDatabase) DeleteUser(ctx context.Context, id uint) error {

	query := `DELETE FROM users WHERE id = $1;`
	err := u.DB.Exec(query, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *UserDatabase) GetuserByName(ctx context.Context, name string) (model.User, error) {

	query := `select * from users where name = $1;`
	var body model.User
	err := u.DB.Raw(query, name).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (u *UserDatabase) GetUserByEmail(ctx context.Context, mail string) (model.User, error) {
	query := `select * from users where email = $1;`
	var body model.User
	err := u.DB.Raw(query, mail).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (u *UserDatabase) GetUserByNumber(ctx context.Context, number string) (model.User, error) {
	query := `select * from users where phone = $1;`
	var body model.User
	err := u.DB.Raw(query, number).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (u *UserDatabase) GetUserById(ctx context.Context, id uint) (model.User, error) {
	query := `select * from users where id = $1;`
	var body model.User
	err := u.DB.Raw(query, id).Scan(&body).Error
	if err != nil {
		return body, err
	}
	return body, nil
}

func (u *UserDatabase) FindUser(ctx context.Context, user model.User) (model.User, error) {

	query := `SELECT * FROM users WHERE id = ? OR email = ? OR phone = ? OR name = ?`
	err := u.DB.Raw(query, user.Id, user.Email, user.Phone, user.Name).Scan(&user).Error
	if err != nil {
		return user, err
	}
	return user, err
}
