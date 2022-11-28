package sqlstore

import (
	"database/sql"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/store"
)

// Store ...
type Store struct {
	db                      *sql.DB
	userRepository          *UserRepository
	matchRepository         *MatchRepository
	matchHistoryRepository  *MatchHistoryRepository
	lastUserMatchRepository *LastUserMatchRepository
}

// New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository == nil {
		s.userRepository = NewUserRepository(s)
	}

	return s.userRepository
}

// Match ...
func (s *Store) Match() store.MatchRepository {
	if s.matchRepository == nil {
		s.matchRepository = NewMatchRepository(s)
	}

	return s.matchRepository
}

// MatchHistory ...
func (s *Store) MatchHistory() store.MatchHistoryRepository {
	if s.matchHistoryRepository == nil {
		s.matchHistoryRepository = NewMatchHistoryRepository(s)
	}

	return s.matchHistoryRepository
}

// LastUserMatch ...
func (s *Store) LastUserMatch() store.LastUserMatchRepository {
	if s.lastUserMatchRepository == nil {
		s.lastUserMatchRepository = NewLastUserMatchRepository(s)
	}

	return s.lastUserMatchRepository
}
