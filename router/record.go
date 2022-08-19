package router

import (
	"mehm8128_study_server/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Record struct {
	Title     string    `json:"title" db:"title"`
	Page      int       `json:"page" db:"page"`
	Time      int       `json:"time" db:"time"`
	Comment   string    `json:"comment" db:"comment"`
	FileID    uuid.UUID `json:"fileId" db:"file_id"`
	CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
}
type RecordFavorite struct {
	CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
}

func getRecords(c echo.Context) error {
	ctx := c.Request().Context()
	records, err := model.GetRecords(ctx)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	for _, record := range records {
		favorites, err := model.GetRecordFavorites(ctx, record.ID)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		if favorites == nil {
			favorites = []model.RecordFavoriteResponse{}
		}
		record.Favorites = favorites
	}
	if records == nil {
		return echo.NewHTTPError(http.StatusOK, []*model.RecordResponse{})
	}
	return echo.NewHTTPError(http.StatusOK, records)
}

func postRecord(c echo.Context) error {
	var record Record
	err := c.Bind(&record)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ctx := c.Request().Context()
	res, err := model.CreateRecord(ctx, record.Title, record.Page, record.Time, record.Comment, record.FileID, record.CreatedBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func getRecord(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	record, err := model.GetRecord(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if record == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "record not found")
	}
	favorites, err := model.GetRecordFavorites(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if favorites == nil {
		favorites = []model.RecordFavoriteResponse{}
	}
	record.Favorites = favorites
	return echo.NewHTTPError(http.StatusOK, record)
}

func putRecord(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var record Record
	err = c.Bind(&record)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ctx := c.Request().Context()
	err = model.PutRecord(ctx, ID, record.Title, record.Page, record.Time, record.Comment, record.FileID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

func deleteRecord(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	err = model.DeleteRecordFavorites(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = model.DeleteRecord(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

func getRecordsByUser(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	record, err := model.GetRecordsByUser(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if record == nil {
		return echo.NewHTTPError(http.StatusOK, []*model.RecordResponse{})
	}
	return echo.NewHTTPError(http.StatusOK, record)
}

func PutRecordFavorite(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var favorite RecordFavorite
	err = c.Bind(&favorite)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ctx := c.Request().Context()
	res, err := model.PutRecordFavorite(ctx, ID, favorite.CreatedBy)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}
