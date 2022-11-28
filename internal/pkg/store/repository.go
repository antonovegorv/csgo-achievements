package store

import "github.com/antonovegorv/csgo-achievements/internal/pkg/models"

// UserRepository ...
type UserRepository interface{} //nolint:gofmt // will add methods later

// MatchRepository ...
type MatchRepository interface {
	Create(*models.Match) (*models.Match, error)
	FindBySharingCode(string) (*models.Match, error)
}

// MatchHistoryRepository ...
type MatchHistoryRepository interface {
	Create(*models.MatchHistory) (*models.MatchHistory, error)
}

// LastUserMatchRepository ...
type LastUserMatchRepository interface {
	GetAll() ([]*models.LastUserMatch, error)
	UpdateByUserID(int, int) error
}
