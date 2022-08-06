package router

import (
	"fmt"
	"mehm8128_study_server/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo-contrib/session"
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
type PutMe struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

func postSignUp(c echo.Context) error {
	var req SignUpRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
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
	sess, err := session.Get("sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "something wrong in getting session")
	}
	sess.Values["userID"] = ID
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "something wrong in saving session")
	}
	res := LoginResponse{
		ID:   *ID,
		Name: req.Name,
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func postLogin(c echo.Context) error {
	var req LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if req.Name == "" || req.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name or password is empty")
	}
	ctx := c.Request().Context()
	user, err := model.GetUserByName(ctx, req.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = bcrypt.CompareHashAndPassword(user.HashedPass, []byte(req.Password))
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "password is invalid")
	}
	sess, err := session.Get("sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "something wrong in getting session")
	}
	sess.Values["userID"] = user.ID.String()
	sess.Options.SameSite = http.SameSiteNoneMode
	sess.Options.Secure = true
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "something wrong in saving session")
	}
	res := LoginResponse{
		ID:   user.ID,
		Name: user.Name,
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func postLogout(c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed in getting session")
	}
	sess.Options.MaxAge = -1
	sess.Values["userID"] = ""
	err = sess.Save(c.Request(), c.Response())
	if err != nil {
		return fmt.Errorf("Failed to delete session: %w", err)
	}
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return echo.NewHTTPError(http.StatusOK)
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

func getMe(c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "something wrong in getting session")
	}
	userID := sess.Values["userID"]
	if userID == nil {
		return echo.NewHTTPError(http.StatusForbidden, "not logged in")
	}
	ctx := c.Request().Context()
	ID, err := uuid.Parse(userID.(string))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err)
	}
	user, err := model.GetUser(ctx, ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, user)
}

func putMe(c echo.Context) error {
	sess, err := session.Get("sessions", c)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "something wrong in getting session")
	}
	userID := sess.Values["userID"]
	if err != nil {
		return echo.NewHTTPError(http.StatusForbidden, "please login")
	}
	var me PutMe
	ctx := c.Request().Context()
	err = c.Bind(&me)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	if userID != me.ID {
		return echo.NewHTTPError(http.StatusForbidden, "your id is invalid")
	}
	if me.Name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is empty")
	}
	me2, err := model.PutMe(ctx, me.ID, me.Name, me.Description)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, me2)
}
