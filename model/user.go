package model

import (
	time "time"
)

type User struct {
	Id        uint      `Json:"id" gorm:"primary_key"`
	Email     string    `Json:"email" gorm:"unique"`
	Password  string    `Json:"password"`
	CreatedAt time.Time `Json:"created_at"`
	UpdatedAt time.Time `Json:"updated_at"`
}

type UserResponse struct {
	Id    uint   `Json:"id" gorm:"primary_key"`
	Email string `Json:"email" gorm:"unique"`
}
