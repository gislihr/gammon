package graph

import "github.com/gislihr/gammon/pkg/gammon/db/store"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store *store.Store
}

func NewResolver(s *store.Store) *Resolver {
	return &Resolver{
		store: s,
	}
}
