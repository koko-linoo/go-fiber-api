package models

type UserCreate struct {
	FullName string  `json:"fullName" validate:"required"`
	Username string  `json:"username" validate:"required"`
	Email    string  `json:"email" validate:"required"`
	Phone    *string `json:"phone"`
	Password string  `json:"password" validate:"required"`
}

func (u UserCreate) ToUser() User {
	return User{
		FullName: u.FullName,
		Username: u.Username,
		Email:    u.Email,
		Phone:    u.Phone,
		Password: u.Password,
	}
}
