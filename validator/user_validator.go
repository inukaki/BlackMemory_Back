package validator

import (
	"go_rest_api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type IUserValidator interface {
	UserCreateValidate(user model.User) error
	UserLoginValidate(user model.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) UserCreateValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			// validation.RuneLength(1, 30).Error("email must be between 1 and 30 characters"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 0).Error("password must be at least 8 characters"),
		),
		validation.Field(
			&user.Name,
			validation.Required.Error("name is required"),
			validation.RuneLength(3, 0).Error("name must be at least 3 characters"),
		),
	)
}

func (uv *userValidator) UserLoginValidate(user model.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			// validation.RuneLength(1, 30).Error("email must be between 1 and 30 characters"),
			is.Email.Error("invalid email format"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(8, 0).Error("password must be at least 8 characters"),
		),
	)
}
