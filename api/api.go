package api

import (
	"github.com/liteseed/bungo/database"
	"github.com/liteseed/bungo/store"
)

const MAX_DATA_ITEM_SIZE = 1_073_824

type API struct {
	db    *database.Database
	store *store.Store
}

func New(
	db *database.Database,
	s *store.Store,
) *API {
	return &API{db: db, store: s}
}