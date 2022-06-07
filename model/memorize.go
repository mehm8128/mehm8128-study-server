package model

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type MemorizeResponse struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	Name      string         `json:"name" db:"name"`
	Words     []WordResponse `json:"words"`
	CreatedAt time.Time      `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time      `json:"updatedAt" db:"updated_at"`
}
type WordResponse struct {
	ID         uuid.UUID `json:"id" db:"id"`
	MemorizeID uuid.UUID `json:"memorizeId" db:"memorize_id"`
	Word       string    `json:"word" db:"word"`
	WordJp     string    `json:"wordJp" db:"word_jp"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

func GetMemorizes(ctx context.Context) ([]*MemorizeResponse, error) {
	var memorizes []*MemorizeResponse
	err := db.SelectContext(ctx, &memorizes, "SELECT * FROM memorizes")
	if err != nil {
		return nil, err
	}
	return memorizes, nil
}

func CreateMemorize(ctx context.Context, name string) (*MemorizeResponse, error) {
	memorizeID := uuid.New()
	date := time.Now()
	_, err := db.ExecContext(ctx, "INSERT INTO memorizes (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4)", memorizeID, name, date, date)
	if err != nil {
		return nil, err
	}
	memorize := &MemorizeResponse{
		ID:        memorizeID,
		Name:      name,
		Words:     []WordResponse{},
		CreatedAt: date,
		UpdatedAt: date,
	}
	return memorize, nil
}

func GetMemorize(ctx context.Context, id uuid.UUID) (*MemorizeResponse, error) {
	var memorize MemorizeResponse
	err := db.GetContext(ctx, &memorize, "SELECT * FROM memorizes WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &memorize, nil
}

func GetWords(ctx context.Context, id uuid.UUID) ([]WordResponse, error) {
	var words []WordResponse
	err := db.SelectContext(ctx, &words, "SELECT * FROM words WHERE memorize_id = $1", id)
	if err != nil {
		return nil, err
	}
	return words, nil
}

func AddWord(ctx context.Context, memorizeID uuid.UUID, word string, wordJp string) (*WordResponse, error) {
	wordID := uuid.New()
	date := time.Now()
	_, err := db.ExecContext(ctx, "INSERT INTO words (id, memorize_id, word, word_jp, created_at) VALUES ($1, $2, $3, $4, $5)", wordID, memorizeID, word, wordJp, date)
	if err != nil {
		return nil, err
	}
	wordRes := &WordResponse{
		ID:         wordID,
		MemorizeID: memorizeID,
		Word:       word,
		WordJp:     wordJp,
		CreatedAt:  date,
	}
	return wordRes, nil
}
