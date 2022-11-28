package models

// LastUserMatch ...
type LastUserMatch struct {
	User  *User
	Match *Match
}

// TestLastUserMatch ...
func TestLastUserMatch(u *User, m *Match) *LastUserMatch {
	return &LastUserMatch{
		User:  u,
		Match: m,
	}
}
