package graph

import (
	"github.com/gislihr/gammon/graph/model"
	"github.com/gislihr/gammon/pkg/gammon/db"
)

func gqlReguestToDb(req *model.GameRequest) db.GameRequest {
	limit := 50
	offset := 0
	if req.Limit != nil {
		limit = *req.Limit
	}

	if req.Offset != nil {
		offset = *req.Offset
	}
	return db.GameRequest{
		Limit:    limit,
		Offset:   offset,
		WinnerId: req.WinnerID,
		LoserId:  req.LoserID,
	}
}
