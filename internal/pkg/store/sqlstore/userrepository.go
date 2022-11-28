package sqlstore

import (
	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
)

// UserRepository ...
type UserRepository struct {
	store *Store
}

// NewUserRepository ...
func NewUserRepository(store *Store) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

// Create ...
func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	if err := r.store.db.QueryRow(
		`
		INSERT INTO users (steam_id, game_authentication_code)
		VALUES ($1, $2)
		RETURNING id;
		`,
		u.SteamID,
		u.GameAuthenticationCode,
	).Scan(&u.ID); err != nil {
		return nil, err
	}

	return u, nil
}
