package validator

import (
	"todo-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ITaskValidator インターフェース
type IUserValidator interface {
	ValidateUser(user model.User) error
}

// userValidator IUserValidator を実装する構造体
type userValidator struct {
}

// NewUserValidator コンストラクタ
func NewUserValidator() IUserValidator {
	return &userValidator{}
}

// ValidateUser ユーザーのバリデーションを行う
func (uv *userValidator) ValidateUser(user model.User) error {
	return validation.ValidateStruct(
		&user,
		validation.Field(
			&user.Email,
			validation.Required.Error("email is required"),
			validation.RuneLength(1, 30).Error("email must be between 1 and 10 characters"),
		),
		validation.Field(
			&user.Password,
			validation.Required.Error("password is required"),
			validation.RuneLength(6, 30).Error("password must be between 1 and 10 characters"),
		),
	)
}
