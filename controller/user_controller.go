package controller

import (
	"net/http"
	"os"
	"time"
	"todo-app/model"
	"todo-app/usecase"

	"github.com/labstack/echo/v4"
)

// IUserController インターフェース
type IUserController interface {
	SignUp(c echo.Context) error
	LogIn(c echo.Context) error
	LogOut(c echo.Context) error
	CsrfToken(c echo.Context) error
}

// userController IUserController を実装する構造体
type userController struct {
	uu usecase.IUserUsecase
}

// NewUserController コンストラクタ
func NewUserController(uu usecase.IUserUsecase) IUserController {
	return &userController{uu}
}

// SignUp サインアップ処理
func (uc *userController) SignUp(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userRes, err := uc.uu.SignUp(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, userRes)
}

// LogIn ログイン処理
func (uc *userController) LogIn(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tokenString, err := uc.uu.LogIn(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(time.Hour * 24)
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true // postmanで確認する際はfalseにする
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// LogOut ログアウト処理
func (uc *userController) LogOut(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = ""
	cookie.Expires = time.Now()
	cookie.Path = "/"
	cookie.Domain = os.Getenv("API_DOMAIN")
	cookie.Secure = true // postmanで確認する際はfalseにする
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}

// CsrfToken CSRFトークンを取得する
func (uc *userController) CsrfToken(c echo.Context) error {
	token := c.Get("csrf").(string)
	return c.JSON(http.StatusOK, echo.Map{
		"csrf_token": token,
	})
}
