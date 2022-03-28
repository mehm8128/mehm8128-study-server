package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type GoalResponse struct {
	ID          uuid.UUID              `json:"id" db:"id"`
	Title       string                 `json:"title" db:"title"`
	Comment     string                 `json:"comment" db:"comment"`
	GoalDate    string                 `json:"goalDate" db:"goal_date"`
	IsCompleted bool                   `json:"isCompleted" db:"is_completed"`
	Favorites   []GoalFavoriteResponse `json:"favorites" db:"favorites"`
	FavoriteNum int                    `json:"favoriteNum" db:"favorite_num"`
	CreatedBy   uuid.UUID              `json:"createdBy" db:"created_by"`
	CreatedAt   time.Time              `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time              `json:"updatedAt" db:"updated_at"`
}

func GetGoals(ctx context.Context) ([]*GoalResponse, error) {
	var goals []*GoalResponse
	err := db.SelectContext(ctx, &goals, "SELECT * FROM goals ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	return goals, nil
}

func CreateGoal(ctx context.Context, title string, comment string, goalDate string, createdBy uuid.UUID) (*GoalResponse, error) {
	goalID := uuid.New()
	date := time.Now()
	var favorites []GoalFavoriteResponse
	_, err := db.ExecContext(ctx, "INSERT INTO goals (id, title, comment, goal_date, favorite_num, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)", goalID, title, comment, goalDate, 0, createdBy, date, date)
	if err != nil {
		return nil, err
	}
	goal := &GoalResponse{
		ID:          goalID,
		Title:       title,
		Comment:     comment,
		GoalDate:    goalDate,
		IsCompleted: false,
		Favorites:   favorites,
		FavoriteNum: 0,
		CreatedBy:   createdBy,
		CreatedAt:   date,
		UpdatedAt:   date,
	}
	return goal, nil
}

func GetGoal(ctx context.Context, id uuid.UUID) (*GoalResponse, error) {
	var goal GoalResponse
	err := db.GetContext(ctx, &goal, "SELECT * FROM goals WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &goal, nil
}

func PutGoal(ctx context.Context, id uuid.UUID, title string, comment string, goalDate string, isCompleted bool) error {
	date := time.Now()
	_, err := db.ExecContext(ctx, "UPDATE goals SET title=$1, comment=$2, goal_date=$3, is_completed=$4, updated_at=$5 WHERE id=$6", title, comment, goalDate, isCompleted, date, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGoal(ctx context.Context, id uuid.UUID) error {
	_, err := db.ExecContext(ctx, "DELETE FROM goals WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetGoalsByUser(ctx context.Context, id uuid.UUID) ([]*GoalResponse, error) {
	var goals []*GoalResponse
	err := db.SelectContext(ctx, &goals, "SELECT * FROM goals WHERE created_by = $1 ORDER BY created_at DESC", id)
	if err != nil {
		return nil, err
	}
	return goals, nil
}

func PutGoalFavorite(ctx context.Context, id uuid.UUID, createdBy uuid.UUID) (*GoalFavoriteResponse, error) {
	favoriteID := uuid.New()
	date := time.Now()
	_, err := db.ExecContext(ctx, "INSERT INTO goal_favorites (id, goal_id, created_by, created_at) VALUES ($1, $2, $3, $4)", favoriteID, id, createdBy, date)
	if err != nil {
		return nil, err
	}
	goal, err := GetGoal(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, "UPDATE goals SET favorite_num=$1, updated_at=$2 WHERE id=$3", goal.FavoriteNum+1, date, id)
	if err != nil {
		return nil, err
	}
	FavoriteResponse := &GoalFavoriteResponse{
		ID:        favoriteID,
		CreatedBy: createdBy,
		CreatedAt: date,
		GoalID:    id,
	}
	return FavoriteResponse, nil
}

func CompleteGoal(ctx context.Context, id uuid.UUID) error {
	date := time.Now()
	_, err := db.ExecContext(ctx, "UPDATE goals SET is_completed=$1, updated_at=$2 WHERE id=$3", true, date, id)
	if err != nil {
		return err
	}
	return nil
}
