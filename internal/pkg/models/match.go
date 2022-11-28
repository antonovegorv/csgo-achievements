package models

// Match ...
type Match struct {
	ID        int
	ShareCode string
	MatchID   uint64
	OutcomeID uint64
	TokenID   uint64
}

// TestMatch ...
func TestMatch() *Match {
	return &Match{
		ShareCode: "CSGO-8Le73-rTm9B-4KD79-rhjJU-x62NF",
		MatchID:   3585047646223466579,
		OutcomeID: 3585054404354506867,
		TokenID:   61828,
	}
}
