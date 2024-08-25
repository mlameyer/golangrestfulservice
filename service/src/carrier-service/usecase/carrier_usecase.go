package usecase

import (
	"carrier-service/domain/model"
	"carrier-service/domain/repository"
	"context"
	"sync"
)

type CarrierUseCase interface {
	GetCarrier(ctx context.Context, id int) (*model.Carrier, error)
	CreateCarrier(ctx context.Context, carrier *model.Carrier, err chan<- error, wg *sync.WaitGroup)
	UpdateCarrier(ctx context.Context, carrier *model.Carrier) error
	DeleteCarrier(ctx context.Context, carrier *model.Carrier) error
}

type carrierUseCase struct {
	repository.CarrierRepository
}

func NewCarrierUseCase(r repository.CarrierRepository) *carrierUseCase {
	return &carrierUseCase{r}
}

func (c *carrierUseCase) GetCarrier(ctx context.Context, id int) (*model.Carrier, error) {
	return c.CarrierRepository.FetchByID(ctx, id)
}

func (c *carrierUseCase) CreateCarrier(ctx context.Context, carrier *model.Carrier, err chan<- error, wg *sync.WaitGroup) {
	c.CarrierRepository.Create(ctx, carrier, err, wg)
}

func (c *carrierUseCase) UpdateCarrier(ctx context.Context, carrier *model.Carrier) error {
	return c.CarrierRepository.Update(ctx, carrier)
}

func (c *carrierUseCase) DeleteCarrier(ctx context.Context, carrier *model.Carrier) error {
	return c.CarrierRepository.Delete(ctx, carrier)
}
