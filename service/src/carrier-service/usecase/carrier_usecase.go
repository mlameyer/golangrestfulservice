package usecase

import (
	"context"
)

type CarrierUseCase interface {
	GetCarrier(ctx context.Context, id int) (*model.Carrier, error)
}

type carrierUseCase struct {
	repository.CarrierRepository
}

func NewCarrierUseCase(r repository.CarrierRepository) CarrierUseCase {
	return &carrierUseCase{r}
}

func (c *carrierUseCase) GetUser(ctx context.Context, id int) (*model.Carrier, error) {
	return c.CarrierRepository.FetchByID(ctx, id)
}
