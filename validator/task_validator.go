package validator

import (
	"todo-app/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ITaskValidator インターフェース
type ITaskValidator interface {
	ValidateTask(task model.Task) error
}

// taskValidator ITaskValidator を実装する構造体
type taskValidator struct {
}

// NewTaskValidator コンストラクタ
func NewTaskValidator() ITaskValidator {
	return &taskValidator{}
}

// ValidateTask タスクのバリデーションを行う
func (tv *taskValidator) ValidateTask(task model.Task) error {
	return validation.ValidateStruct(
		&task,
		validation.Field(
			&task.Title,
			validation.Required.Error("title is required"),
			validation.RuneLength(1, 10).Error("title must be between 1 and 10 characters"),
		),
	)
}
