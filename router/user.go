package router

import (
	"mehm8128_study_server/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type SignUpRequest struct {
	Name        string `json:"name"`
	Password    string `json:"password"`
	Description string `json:"description"`
}
type LoginRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}
type LoginResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func postSignUp(c echo.Context) error {
	var req SignUpRequest
	c.Bind(&req)
	if req.Name == "" || req.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name or password is empty")
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "hash error")
	}
	ctx := c.Request().Context()
	ID, err := model.CreateUser(ctx, req.Name, hashedPass, req.Description)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	res := LoginResponse{
		ID:   *ID,
		Name: req.Name,
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func postLogin(c echo.Context) error {
	var req LoginRequest
	c.Bind(&req)
	if req.Name == "" || req.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name or password is empty")
	}
	ctx := c.Request().Context()
	user, err := model.GetUserByName(ctx, req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user not found")
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPass, []byte(req.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "password is invalid")
	}
	res := LoginResponse{
		ID:   user.ID,
		Name: user.Name,
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func getUsers(c echo.Context) error {
	ctx := c.Request().Context()
	users, err := model.GetUsers(ctx)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if users == nil {
		return echo.NewHTTPError(http.StatusOK, []*model.UserResponse{})
	}
	return echo.NewHTTPError(http.StatusOK, users)
}

func getUser(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	user, err := model.GetUser(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if user == nil {
		return echo.NewHTTPError(http.StatusOK, []*model.UserResponse{})
	}
	return echo.NewHTTPError(http.StatusOK, user)
}
