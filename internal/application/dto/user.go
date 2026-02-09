package dto

type UserInputDto struct {
	Email         string  `json:"email" binding:"required,checkemail"`
	Username      string  `json:"username" binding:"required"`
	Bio           *string `json:"bio"`
	Password      string  `json:"password" binding:"required,checkpassword"`
	Password_Conf string  `json:"password_conf" binding:"required,eqfield=Password"`
}

type LoginInputDto struct {
	Email    string `json:"email" binding:"required,checkemail"`
	Password string `json:"password" binding:"required"`
}

type GoogleLoginInputDto struct {
	IDtoken string `json:"id_token" binding:"required"`
}
