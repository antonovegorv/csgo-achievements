package models

// MatchHistory ...
type MatchHistory struct {
	ID    int
	User  *User
	Match *Match
}
