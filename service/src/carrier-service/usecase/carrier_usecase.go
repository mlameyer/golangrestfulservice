package usecase

import (
	"carrier-service/domain/model"
	"carrier-service/domain/repository"
	"context"
)

type CarrierUseCase interface {
	GetCarrier(ctx context.Context, id int) (*model.Carrier, error)
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
