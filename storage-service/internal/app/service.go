package app

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/badtripdude/gh-analyzer/internal/model"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db: db}
}

// ! demo/mock struct used for testing the consumer
// ? this structure is a placeholder and will later be refactored
type AnalysisResultsMsg struct {
	TaskID     string         `json:"task_id"`
	UserID     string         `json:"user_id"`
	Username   string         `json:"username"`
	RepoCount  int            `json:"repo_count"`
	TotalStars int            `json:"total_stars"`
	Languages  map[string]int `json:"languages"`
	LastCommit time.Time      `json:"last_commit"`
	ComputedAt time.Time      `json:"computed_at"`
}

func (s *Service) HandleAnalysisResults(value []byte) error {
	var msg AnalysisResultsMsg
	if err := json.Unmarshal(value, &msg); err != nil {
		return err
	}

	langsJSON, _ := json.Marshal(msg.Languages)

	analysis := model.GitHubAnalysis{
		TaskID:     msg.TaskID,
		UserID:     msg.UserID,
		Username:   msg.Username,
		RepoCount:  msg.RepoCount,
		TotalStars: msg.TotalStars,
		Languages:  datatypes.JSON(langsJSON),
		LastCommit: msg.LastCommit,
		Status:     "done",
	}
	// TODO: replace hardcoded status strings with constants
	if err := s.db.WithContext(context.Background()).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "task_id"}},
			DoUpdates: clause.AssignmentColumns([]string{"repo_count", "total_stars", "languages", "last_commit", "status", "updated_at"}),
		}).
		Create(&analysis).Error; err != nil {
		return err
	}

	log.Printf("Saved analysis for task_id=%s user=%s", msg.TaskID, msg.Username)
	return nil
}
