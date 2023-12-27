package req

type UserLoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserSignUpReq struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
