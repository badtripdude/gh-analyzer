package model

import (
	"time"

	"gorm.io/datatypes"
)

type GitHubAnalysis struct {
	ID         uint   `gorm:"primaryKey;autoIncrement"`
	TaskID     string `gorm:"index"`
	UserID     string `gorm:"index"`
	RepoCount  int
	TotalStars int
	Languages  datatypes.JSON
	LastCommit time.Time

	CreatedAt time.Time
	UpdatedAt time.Time
}
