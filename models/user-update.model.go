package models

type UserUpdate struct {
	FullName *string `json:"fullName" validate:"omitempty"`
	Username *string `json:"username" validate:"omitempty"`
	Email    *string `json:"email" validate:"omitempty"`
	Phone    *string `json:"phone" validate:"omitempty"`
	Password *string `json:"password" validate:"omitempty"`
}

func (u UserUpdate) GetMap() map[string]any {

	updates := make(map[string]any)

	if u.FullName != nil {
		updates["full_name"] = u.FullName
	}
	if u.Username != nil {
		updates["username"] = u.Username
	}
	if u.Email != nil {
		updates["email"] = u.Email
	}
	if u.Phone != nil {
		updates["phone"] = u.Phone
	}
	if u.Password != nil {
		updates["password"] = u.Password
	}

	return updates
}
