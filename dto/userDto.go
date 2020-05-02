package dto // Data Transfer Object

import (
	"WebFull/model"
)

type UserDTO struct {
	Name     string `json:"name"`
	Telphone string `json:"telphone"`
}

func ToUserDTO(user model.User) UserDTO {
	return UserDTO{
		Name:     user.Name,
		Telphone: user.Tel,
	}
}
