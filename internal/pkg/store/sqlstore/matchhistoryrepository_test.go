package sqlstore_test

import (
	"testing"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
	"github.com/antonovegorv/csgo-achievements/internal/pkg/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestMatchHistoryRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("matches_history", "matches")

	s := sqlstore.New(db)
	u := models.TestUser()
	m := models.TestMatch()

	u, _ = s.User().Create(u)
	m, _ = s.Match().Create(m)

	mh := models.TestMatchHistory(u, m)

	mh, err := s.MatchHistory().Create(mh)
	assert.NoError(t, err)
	assert.NotNil(t, mh)
}
