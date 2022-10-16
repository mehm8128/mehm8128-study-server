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

	if fileID == uuid.Nil {
		_, err := db.ExecContext(ctx, "INSERT INTO records (id, title, page, time, comment, favorite_num, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", recordID, title, page, timeRecord, comment, 0, createdBy, date, date)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := db.ExecContext(ctx, "INSERT INTO records (id, title, page, time, comment, favorite_num, file_id, created_by, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", recordID, title, page, timeRecord, comment, 0, fileID, createdBy, date, date)
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
	err := db.GetContext(ctx, &record, "SELECT * FROM records WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func PutRecord(ctx context.Context, id uuid.UUID, title string, page int, timeRecord int, comment string, fileID uuid.UUID) error {
	date := time.Now()
	if fileID == uuid.Nil {
		_, err := db.ExecContext(ctx, "UPDATE records SET title=?, page=?, time=?, comment=?, updated_at=? WHERE id=?", title, page, timeRecord, comment, date, id)
		if err != nil {
			return err
		}
	} else {
		_, err := db.ExecContext(ctx, "UPDATE records SET title=?, page=?, time=?, comment=?, file_id=?, updated_at=? WHERE id=?", title, page, timeRecord, comment, fileID, date, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func DeleteRecord(ctx context.Context, id uuid.UUID) error {
	_, err := db.ExecContext(ctx, "DELETE FROM records WHERE id=?", id)
	if err != nil {
		return err
	}
	return nil
}

func GetRecordsByUser(ctx context.Context, id uuid.UUID) ([]*RecordResponse, error) {
	var records []*RecordResponse
	err := db.SelectContext(ctx, &records, "SELECT * FROM records WHERE created_by = ? ORDER BY created_at DESC", id)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func PutRecordFavorite(ctx context.Context, id uuid.UUID, createdBy uuid.UUID) (*RecordFavoriteResponse, error) {
	favoriteID := uuid.New()
	date := time.Now()
	_, err := db.ExecContext(ctx, "INSERT INTO record_favorites (id, record_id, created_by, created_at) VALUES (?, ?, ?, ?)", favoriteID, id, createdBy, date)
	if err != nil {
		return nil, err
	}
	record, err := GetRecord(ctx, id)
	if err != nil {
		return nil, err
	}
	_, err = db.ExecContext(ctx, "UPDATE records SET favorite_num=?, updated_at=? WHERE id=?", record.FavoriteNum+1, date, id)
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
