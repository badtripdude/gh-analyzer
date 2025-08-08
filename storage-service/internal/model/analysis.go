package model

import (
	"time"

	"gorm.io/datatypes"
)

// ! demo/mock struct used for testing the consumer
// ? this model is a placeholder and will later be refactore
type GitHubAnalysis struct {
	TaskID     string `gorm:"primaryKey"`
	UserID     string `gorm:"index"`
	Username   string
	RepoCount  int
	TotalStars int
	Languages  datatypes.JSON
	LastCommit time.Time
	Status     string `gorm:"index"` //? pending, done, failed
	PdfURL     *string
	ErrorCode  *string
	ErrorMsg   *string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
