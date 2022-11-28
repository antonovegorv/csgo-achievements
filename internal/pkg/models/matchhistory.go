package models

// MatchHistory ...
type MatchHistory struct {
	ID    int
	User  *User
	Match *Match
}

// TestMatchHistory ...
func TestMatchHistory(u *User, m *Match) *MatchHistory {
	return &MatchHistory{
		User:  u,
		Match: m,
	}
}
