package helper

import (
	"errors"
	"strconv"

	"github.com/abhinandpn/MicroServices-Demo/User-Service/pkg/model"
	"github.com/gin-gonic/gin"
)

func CheckUser(user, CurUser model.User) error {
	var err error

	if CurUser.Email == user.Email {
		err = errors.New("user already exists with this email")
	}

	if CurUser.Name == user.Name {
		err = errors.New("user already exists with this name")
	}

	if CurUser.Phone == user.Phone {
		err = errors.New("user already exists with this phone number")
	}

	return err
}

func StringToUInt(str string) (uint, error) {
	if str == "" {
		return 0, errors.New("empty string ")
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	// fmt.Println("xxxxx", uint(val))
	return uint(val), nil
}

func GetUserId(ctx *gin.Context) uint {

	UserId := ctx.GetString("userId") // string Type
	Id, err := strconv.Atoi(UserId)   // Int type
	if err != nil {
		return 0
	}
	return uint(Id) // current User Id

}
