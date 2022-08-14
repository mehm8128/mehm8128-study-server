package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type FileResponse struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FileName  string    `json:"fileName" db:"file_name"`
	CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

func CreateFile(ctx context.Context, ID uuid.UUID, fileName string, createdBy uuid.UUID) (*FileResponse, error) {
	fileID := uuid.New()
	date := time.Now()
	_, err := db.ExecContext(ctx, "INSERT INTO files (id, file_name, created_by, created_at) VALUES ($1, $2, $3, $4)", fileID, fileName, createdBy, date)
	if err != nil {
		return nil, err
	}
	file := &FileResponse{
		ID:        fileID,
		FileName:  fileName,
		CreatedBy: createdBy,
		CreatedAt: date,
	}
	return file, nil
}

func GetFile(ctx context.Context, id uuid.UUID) (*FileResponse, error) {
	var file FileResponse
	err := db.GetContext(ctx, &file, "SELECT * FROM files WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &file, nil
}
