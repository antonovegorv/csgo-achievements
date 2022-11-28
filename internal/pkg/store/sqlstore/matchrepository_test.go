package sqlstore_test

import (
	"testing"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
	"github.com/antonovegorv/csgo-achievements/internal/pkg/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestMatchRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("matches")

	s := sqlstore.New(db)
	m := models.TestMatch()

	m, err := s.Match().Create(m)
	assert.NoError(t, err)
	assert.NotNil(t, m)
}

func TestMatchRepository_FindBySharingCode(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("matches")

	s := sqlstore.New(db)
	m := models.TestMatch()

	_, err := s.Match().FindBySharingCode(m.ShareCode)
	assert.Error(t, err)

	s.Match().Create(m)

	m, err = s.Match().FindBySharingCode(m.ShareCode)
	assert.NoError(t, err)
	assert.NotNil(t, m)
}
