package sqlstore

import (
	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
)

// MatchHistoryRepository ...
type MatchHistoryRepository struct {
	store *Store
}

// NewMatchHistoryRepository ...
func NewMatchHistoryRepository(store *Store) *MatchHistoryRepository {
	return &MatchHistoryRepository{
		store: store,
	}
}

// Create ...
func (r *MatchHistoryRepository) Create(mh *models.MatchHistory) (*models.MatchHistory, error) {
	if err := r.store.db.QueryRow(
		`
		INSERT INTO matches_history (user_id, match_id)
		VALUES ($1, $2)
		RETURNING id;
		`,
		mh.User.ID,
		mh.Match.ID,
	).Scan(&mh.ID); err != nil {
		return nil, err
	}

	return mh, nil
}
