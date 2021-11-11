package graph

import "github.com/gislihr/gammon/pkg/gammon/db"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	store *db.Store
}

func NewResolver(s *db.Store) *Resolver {
	return &Resolver{
		store: s,
	}
}
