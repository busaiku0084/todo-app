package repository

import (
	"todo-app/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

// userRepository は IUserRepository を実装する構造体
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository コンストラクタ
func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

// GetUserByEmail メールアドレスからユーザーを取得する
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email = ?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

// CreateUser ユーザーを作成する
func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
