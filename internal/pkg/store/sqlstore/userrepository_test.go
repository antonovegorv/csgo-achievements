package sqlstore_test

import (
	"testing"

	"github.com/antonovegorv/csgo-achievements/internal/pkg/models"
	"github.com/antonovegorv/csgo-achievements/internal/pkg/store/sqlstore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	db, teardown := sqlstore.TestDB(t, databaseURL)
	defer teardown("users")

	s := sqlstore.New(db)
	u := models.TestUser()

	u, err := s.User().Create(u)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
