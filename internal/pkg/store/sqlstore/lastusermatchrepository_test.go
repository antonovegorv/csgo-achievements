package sqlstore_test

import (
	"testing"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
	"github.com/antonovegorv/csgo-achievements/internal/pkg/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestLastUserMatchRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users", "matches", "last_users_matches")

	s := sqlstore.New(db)
	u := models.TestUser()
	m := models.TestMatch()

	u, _ = s.User().Create(u)
	m, _ = s.Match().Create(m)

	lum := models.TestLastUserMatch(u, m)
	err := s.LastUserMatch().Create(lum)
	assert.NoError(t, err)
}

func TestLastUserMatchRepository_FindByUserID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users", "matches", "last_users_matches")

	s := sqlstore.New(db)

	u := models.TestUser()
	m := models.TestMatch()

	_, err := s.LastUserMatch().FindByUserID(u.ID)
	assert.Error(t, err)

	u, _ = s.User().Create(u)
	m, _ = s.Match().Create(m)

	lum := models.TestLastUserMatch(u, m)
	s.LastUserMatch().Create(lum)

	lum, err = s.LastUserMatch().FindByUserID(u.ID)
	assert.NoError(t, err)
	assert.NotNil(t, lum)
}

func TestLastUserMatchRepository_GetAll(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users", "matches", "last_users_matches")

	s := sqlstore.New(db)
	u := models.TestUser()
	m := models.TestMatch()

	u, _ = s.User().Create(u)
	m, _ = s.Match().Create(m)

	lum := models.TestLastUserMatch(u, m)
	s.LastUserMatch().Create(lum)

	lums, err := s.LastUserMatch().GetAll()
	assert.NoError(t, err)
	assert.NotNil(t, lums)
	assert.Len(t, lums, 1)
}

func TestLastUserMatchRepository_UpdateByUserID(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users", "matches", "last_users_matches")

	s := sqlstore.New(db)
	u := models.TestUser()
	m1 := models.TestMatch()
	m2 := models.TestMatch()
	m2.ShareCode = "ABCDEFG..."

	u, _ = s.User().Create(u)
	m1, _ = s.Match().Create(m1)
	m2, _ = s.Match().Create(m2)

	lum := models.TestLastUserMatch(u, m1)
	s.LastUserMatch().Create(lum)

	err := s.LastUserMatch().UpdateByUserID(u.ID, m2.ID)
	assert.NoError(t, err)

	lum, _ = s.LastUserMatch().FindByUserID(u.ID)
	assert.Equal(t, lum.Match.ShareCode, m2.ShareCode)
}
