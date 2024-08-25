package repository

import (
	"carrier-service/domain/model"
	"context"
	"sync"
)

type CarrierRepository interface {
	FetchByID(ctx context.Context, id int) (*model.Carrier, error)
	Create(ctx context.Context, carrier *model.Carrier, err chan<- error, wg *sync.WaitGroup)
	Update(ctx context.Context, carrier *model.Carrier) error
	Delete(ctx context.Context, carrier *model.Carrier) error
}
