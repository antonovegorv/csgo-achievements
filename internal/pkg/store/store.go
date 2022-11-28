package store

// Store ...
type Store interface {
	User() UserRepository
	Match() MatchRepository
	MatchHistory() MatchHistoryRepository
	LastUserMatch() LastUserMatchRepository
}
