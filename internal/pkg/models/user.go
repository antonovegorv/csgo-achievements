package models

// User ...
type User struct {
	ID                     int
	SteamID                int
	GameAuthenticationCode string
}

// TestUser ...
func TestUser() *User {
	return &User{
		SteamID:                1234567890,
		GameAuthenticationCode: "GAME-AUTH-CODE",
	}
}
