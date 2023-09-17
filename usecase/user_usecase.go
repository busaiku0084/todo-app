package usecase

import (
	"time"
	"todo-app/model"
	"todo-app/repository"
	"todo-app/validator"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// IUserUsecase インターフェース
type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	LogIn(user model.User) (string, error)
}

// userUsecase IUserUsecase を実装する構造体
type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

// NewUserUsecase コンストラクタ
func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

// SignUp サインアップ処理
func (uu *userUsecase) SignUp(user model.User) (model.UserResponse, error) {
	if err := uu.uv.ValidateUser(user); err != nil {
		return model.UserResponse{}, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return model.UserResponse{}, err
	}

	newUser := model.User{Email: user.Email, Password: string(hash)}
	if err := uu.ur.CreateUser(&newUser); err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{Id: newUser.Id, Email: newUser.Email}

	return resUser, nil
}

// LogIn ログイン処理
func (uu *userUsecase) LogIn(user model.User) (string, error) {
	if err := uu.uv.ValidateUser(user); err != nil {
		return "", err
	}

	storedUser := model.User{}

	if err := uu.ur.GetUserByEmail(&storedUser, user.Email); err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  storedUser.Id,
		"exp": time.Now().Add(time.Hour * 12).Unix(),
	})

	tokenString, err := token.SignedString([]byte("SECRET"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
