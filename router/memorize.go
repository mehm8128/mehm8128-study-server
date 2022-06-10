package router

import (
	"mehm8128_study_server/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Memorize struct {
	Name string `json:"name" db:"name"`
}
type Word struct {
	Word   string `json:"word" db:"word"`
	WordJp string `json:"wordJp" db:"word_jp"`
}

func getMemorizes(c echo.Context) error {
	ctx := c.Request().Context()
	memorizes, err := model.GetMemorizes(ctx)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if memorizes == nil {
		return echo.NewHTTPError(http.StatusOK, []*model.MemorizeResponse{})
	}
	return echo.NewHTTPError(http.StatusOK, memorizes)
}

func postMemorize(c echo.Context) error {
	var memorize Memorize
	c.Bind(&memorize)
	ctx := c.Request().Context()
	//todo:権限チェック
	res, err := model.CreateMemorize(ctx, memorize.Name)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func getMemorize(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	memorize, err := model.GetMemorize(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if memorize == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "memorize not found")
	}
	words, err := model.GetWords(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if words == nil {
		words = []model.WordResponse{}
	}
	memorize.Words = words
	return echo.NewHTTPError(http.StatusOK, memorize)
}

func postWord(c echo.Context) error {
	var word Word
	c.Bind(&word)
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	//todo:権限チェック
	res, err := model.AddWord(ctx, ID, word.Word, word.WordJp)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}
