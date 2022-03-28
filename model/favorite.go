package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RecordFavoriteResponse struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	RecordID  uuid.UUID `json:"recordId" db:"record_id"`
}
type GoalFavoriteResponse struct {
	ID        uuid.UUID `json:"id" db:"id"`
	CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	GoalID    uuid.UUID `json:"goalId" db:"goal_id"`
}

func GetRecordFavorites(ctx context.Context, id uuid.UUID) ([]RecordFavoriteResponse, error) {
	var favorites []RecordFavoriteResponse
	err := db.SelectContext(ctx, &favorites, "SELECT * FROM record_favorites WHERE record_id = $1", id)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func GetGoalFavorites(ctx context.Context, id uuid.UUID) ([]GoalFavoriteResponse, error) {
	var favorites []GoalFavoriteResponse
	err := db.SelectContext(ctx, &favorites, "SELECT * FROM goal_favorites WHERE goal_id = $1", id)
	if err != nil {
		return nil, err
	}
	return favorites, nil
}

func DeleteRecordFavorites(ctx context.Context, id uuid.UUID) error {
	_, err := db.ExecContext(ctx, "DELETE FROM record_favorites WHERE record_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteGoalFavorites(ctx context.Context, id uuid.UUID) error {
	_, err := db.ExecContext(ctx, "DELETE FROM goal_favorites WHERE goal_id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
