package router

import (
	"mehm8128_study_server/model"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Goal struct {
	Title       string    `json:"title" db:"title"`
	Comment     string    `json:"comment" db:"comment"`
	GoalDate    string    `json:"goalDate" db:"goal_date"`
	IsCompleted bool      `json:"isCompleted" db:"is_completed"`
	CreatedBy   uuid.UUID `json:"createdBy" db:"created_by"`
}
type GoalFavorite struct {
	CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
}

func getGoals(c echo.Context) error {
	ctx := c.Request().Context()
	goals, err := model.GetGoals(ctx)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	for _, goal := range goals {
		favorites, err := model.GetGoalFavorites(ctx, goal.ID)
		if err != nil {
			c.Logger().Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		}
		if favorites == nil {
			favorites = []model.GoalFavoriteResponse{}
		}
		goal.Favorites = favorites
	}
	if goals == nil {
		return echo.NewHTTPError(http.StatusOK, []*model.GoalResponse{})
	}
	return echo.NewHTTPError(http.StatusOK, goals)
}

func postGoal(c echo.Context) error {
	var goal Goal
	err := c.Bind(&goal)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ctx := c.Request().Context()
	res, err := model.CreateGoal(ctx, goal.Title, goal.Comment, goal.GoalDate, goal.CreatedBy)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}

func getGoal(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	goal, err := model.GetGoal(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if goal == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "goal not found")
	}
	favorites, err := model.GetGoalFavorites(ctx, goal.ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if favorites == nil {
		favorites = []model.GoalFavoriteResponse{}
	}
	goal.Favorites = favorites
	return echo.NewHTTPError(http.StatusOK, goal)
}

func putGoal(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var goal Goal
	err = c.Bind(&goal)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ctx := c.Request().Context()
	err = model.PutGoal(ctx, ID, goal.Title, goal.Comment, goal.GoalDate, goal.IsCompleted)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

func deleteGoal(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	err = model.DeleteGoalFavorites(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	err = model.DeleteGoal(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK)
}

func getGoalsByUser(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	ctx := c.Request().Context()
	goal, err := model.GetGoalsByUser(ctx, ID)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	if goal == nil {
		return echo.NewHTTPError(http.StatusOK, []*model.GoalResponse{})
	}
	return echo.NewHTTPError(http.StatusOK, goal)
}

func putGoalFavorite(c echo.Context) error {
	ID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var favorite GoalFavorite
	err = c.Bind(&favorite)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request")
	}
	ctx := c.Request().Context()
	res, err := model.PutGoalFavorite(ctx, ID, favorite.CreatedBy)
	if err != nil {
		c.Logger().Error(err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return echo.NewHTTPError(http.StatusOK, res)
}
