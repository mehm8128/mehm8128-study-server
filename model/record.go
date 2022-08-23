package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type RecordResponse struct {
	ID          uuid.UUID                `json:"id" db:"id"`
	Title       string                   `json:"title" db:"title"`
	Page        int                      `json:"page" db:"page"`
	Time        int                      `json:"time" db:"time"`
	Comment     string                   `json:"comment" db:"comment"`
	Favorites   []RecordFavoriteResponse `json:"favorites" db:"favorites"`
	FavoriteNum int                      `json:"favoriteNum" db:"favorite_num"`
	FileID      uuid.UUID                `json:"fileId" db:"file_id"`
	CreatedBy   uuid.UUID                `json:"createdBy" db:"created_by"`
	CreatedAt   time.Time                `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time                `json:"updatedAt" db:"updated_at"`
}

func GetRecords(ctx context.Context) ([]*RecordResponse, error) {
	var records []*RecordResponse
	err := db.SelectContext(ctx, &records, "SELECT * FROM records ORDER BY created_at DESC")
	if err != nil {
		return nil, err
	}
	return records, nil
}

func CreateRecord(ctx context.Context, title string, page int, timeRecord int, comment string, fileID uuid.UUID, createdBy uuid.UUID) (*RecordResponse, error) {
	recordID := uuid.New()
	date := time.Now()
	var favorites []RecordFavoriteResponse
	zeroUuid, err := uuid.Parse("00000000-0000-0000-0000-000000000000")
	if err != nil {
		return nil, err
	}
	if fileID == zeroUuid {
		_, err = db.ExecContext(ctx, "INSERT INTO records (id, title, page, time, comment, favorite_num, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", recordID, title, page, timeRecord, comment, 0, createdBy, date, date)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = db.ExecContext(ctx, "INSERT INTO records (id, title, page, time, comment, favorite_num, file_id, created_by, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)", recordID, title, page, timeRecord, comment, 0, fileID, createdBy, date, date)
		if err != nil {
			return nil, err
		}
	}
	record := &RecordResponse{
		ID:          recordID,
		Title:       title,
		Page:        page,
		Time:        timeRecord,
		Comment:     comment,
		Favorites:   favorites,
		FavoriteNum: 0,
		FileID:      fileID,
		CreatedBy:   createdBy,
		CreatedAt:   date,
		UpdatedAt:   date,
	}
	return record, nil
}

func GetRecord(ctx context.Context, id uuid.UUID) (*RecordResponse, error) {
	var record RecordResponse
	err := db.GetContext(ctx, &record, "SELECT * FROM records WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func PutRecord(ctx context.Context, id uuid.UUID, title string, page int, timeRecord int, comment string, fileID uuid.UUID) error {
	date := time.Now()
	_, err := db.ExecContext(ctx, "UPDATE records SET title=$1, page=$2, time=$3, comment=$4, file_id=$5, updated_at=$6 WHERE id=$7", title, page, timeRecord, comment, fileID, date, id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRecord(ctx context.Context, id uuid.UUID) error {
	_, err := db.ExecContext(ctx, "DELETE FROM records WHERE id=$1", id)
	if err != nil {
		return err
	}
	return nil
}

func GetRecordsByUser(ctx context.Context, id uuid.UUID) ([]*RecordResponse, error) {
	var records []*RecordResponse
	err := db.SelectContext(ctx, &records, "SELECT * FROM records WHERE created_by = $1 ORDER BY created_at DESC", id)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func PutRecordFavorite(ctx context.Context, id uuid.UUID, createdBy uuid.UUID) (*RecordFavoriteResponse, error) {
	favoriteID := uuid.New()
	date := time.Now()
	_, err := db.ExecContext(ctx, "INSERT INTO record_favorites (id, record_id, created_by, created_at) VALUES ($1, $2, $3, $4)", favoriteID, id, createdBy, date)
	if err != nil {
		return nil, err
	}
	record, err := GetRecord(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, "UPDATE records SET favorite_num=$1, updated_at=$2 WHERE id=$3", record.FavoriteNum+1, date, id)
	if err != nil {
		return nil, err
	}
	FavoriteResponse := &RecordFavoriteResponse{
		ID:        favoriteID,
		CreatedBy: createdBy,
		CreatedAt: date,
		RecordID:  id,
	}
	return FavoriteResponse, nil
}
