package sqlstore

import (
	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
)

// LastUserMatchRepository ...
type LastUserMatchRepository struct {
	store *Store
}

// NewLastUserMatchRepository ...
func NewLastUserMatchRepository(store *Store) *LastUserMatchRepository {
	return &LastUserMatchRepository{
		store: store,
	}
}

// Create ...
func (r *LastUserMatchRepository) Create(lum *models.LastUserMatch) error {
	if _, err := r.store.db.Exec(
		`
		INSERT INTO last_users_matches (user_id, match_id)
		VALUES ($1, $2);
		`,
		lum.User.ID,
		lum.Match.ID,
	); err != nil {
		return err
	}

	return nil
}

// FindByUserID ...
func (r *LastUserMatchRepository) FindByUserID(userID int) (*models.LastUserMatch, error) {
	lum := &models.LastUserMatch{
		User:  &models.User{},
		Match: &models.Match{},
	}

	if err := r.store.db.QueryRow(
		`
		SELECT u.*, m.* FROM users u
		JOIN last_users_matches lum
		ON u.id = lum.user_id AND u.id = $1
		JOIN matches m ON
		lum.match_id = m.id;
		`,
		userID,
	).Scan(
		&lum.User.ID,
		&lum.User.SteamID,
		&lum.User.GameAuthenticationCode,
		&lum.Match.ID,
		&lum.Match.ShareCode,
		&lum.Match.MatchID,
		&lum.Match.OutcomeID,
		&lum.Match.TokenID,
	); err != nil {
		return nil, err
	}

	return lum, nil
}

// GetAll ...
func (r *LastUserMatchRepository) GetAll() ([]*models.LastUserMatch, error) {
	rows, err := r.store.db.Query(
		`
		SELECT u.*, m.* 
		FROM last_users_matches lum 
		JOIN users u ON lum.user_id = u.id 
		JOIN matches m ON lum.match_id = m.id;
		`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	lastUsersMatches := []*models.LastUserMatch{}

	for rows.Next() {
		lum := &models.LastUserMatch{
			User:  &models.User{},
			Match: &models.Match{},
		}

		if err := rows.Scan(
			&lum.User.ID,
			&lum.User.SteamID,
			&lum.User.GameAuthenticationCode,
			&lum.Match.ID,
			&lum.Match.ShareCode,
			&lum.Match.MatchID,
			&lum.Match.OutcomeID,
			&lum.Match.TokenID,
		); err != nil {
			return nil, err
		}

		lastUsersMatches = append(lastUsersMatches, lum)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return lastUsersMatches, nil
}

// UpdateByUserID ...
func (r *LastUserMatchRepository) UpdateByUserID(userID, matchID int) error {
	if _, err := r.store.db.Exec(
		`
		UPDATE last_users_matches
		SET match_id = $1
		WHERE user_id = $2;
		`,
		matchID,
		userID,
	); err != nil {
		return err
	}

	return nil
}
