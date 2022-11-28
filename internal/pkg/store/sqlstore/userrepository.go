package sqlstore

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
