package repository

import (
	"carrier-service/domain/model"
	"context"
)

type CarrierRepository interface {
	FetchByID(ctx context.Context, id int) (*model.Carrier, error)
}
