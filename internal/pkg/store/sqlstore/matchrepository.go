package sqlstore

import (
	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
)

// MatchRepository ...
type MatchRepository struct {
	store *Store
}

// NewMatchRepository ...
func NewMatchRepository(store *Store) *MatchRepository {
	return &MatchRepository{
		store: store,
	}
}

// Create ...
func (r *MatchRepository) Create(m *models.Match) (*models.Match, error) {
	if err := r.store.db.QueryRow(
		`
		INSERT INTO matches (share_code, match_id, outcome_id, token_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
		`,
		m.ShareCode,
		m.MatchID,
		m.OutcomeID,
		m.TokenID,
	).Scan(&m.ID); err != nil {
		return nil, err
	}

	return m, nil
}

// FindBySharingCode ...
func (r *MatchRepository) FindBySharingCode(sharingCode string) (*models.Match, error) {
	m := &models.Match{}

	if err := r.store.db.QueryRow(
		`
		SELECT id, share_code, match_id, outcome_id, token_id
		FROM matches
		WHERE share_code = $1;
		`,
		sharingCode,
	).Scan(
		&m.ID,
		&m.ShareCode,
		&m.MatchID,
		&m.OutcomeID,
		&m.TokenID,
	); err != nil {
		return nil, err
	}

	return m, nil
}
