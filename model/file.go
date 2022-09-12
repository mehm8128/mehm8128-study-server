package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

//	type FileResponse struct {
//		ID        uuid.UUID `json:"id" db:"id"`
//		FileName  string    `json:"fileName" db:"file_name"`
//		CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
//		CreatedAt time.Time `json:"createdAt" db:"created_at"`
//	}
type FileResponse struct {
	ID        uuid.UUID `json:"id" db:"id"`
	FileName  string    `json:"fileName" db:"file_name"`
	File      string    `json:"file" db:"file"`
	CreatedBy uuid.UUID `json:"createdBy" db:"created_by"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

// func CreateFile(ctx context.Context, fileID uuid.UUID, fileName string, createdBy uuid.UUID) (*FileResponse, error) {
// 	date := time.Now()
// 	_, err := db.ExecContext(ctx, "INSERT INTO files (id, file_name, created_by, created_at) VALUES (?, ?, ?, ?)", fileID, fileName, createdBy, date)
// 	if err != nil {
// 		return nil, err
// 	}
// 	file := &FileResponse{
// 		ID:        fileID,
// 		FileName:  fileName,
// 		CreatedBy: createdBy,
// 		CreatedAt: date,
// 	}
// 	return file, nil
// }

func CreateFile(ctx context.Context, fileName string, file string, createdBy uuid.UUID) (*FileResponse, error) {
	fileID := uuid.New()
	date := time.Now()
	_, err := db.ExecContext(ctx, "INSERT INTO files (id, file_name, file, created_by, created_at) VALUES (?, ?, ?, ?, ?)", fileID, fileName, file, createdBy, date)
	if err != nil {
		return nil, err
	}
	fileResponse := &FileResponse{
		ID:        fileID,
		FileName:  fileName,
		File:      file,
		CreatedBy: createdBy,
		CreatedAt: date,
	}
	return fileResponse, nil
}

func GetFile(ctx context.Context, id uuid.UUID) (*FileResponse, error) {
	var file FileResponse
	err := db.GetContext(ctx, &file, "SELECT * FROM files WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	return &file, nil
}
