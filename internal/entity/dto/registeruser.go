package dto

import "gateway/internal/entity"

type RegisterUser struct {
	entity.User
	Password string `json:"password"`
}
