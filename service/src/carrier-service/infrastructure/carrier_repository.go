package infrastructure

import (
	"carrier-service/domain/model"
	"carrier-service/domain/repository"
	"context"
	"gorm.io/gorm"
)

type carrierRepository struct {
	Conn *gorm.DB
}

func NewCarrierRepository(Conn *gorm.DB) repository.CarrierRepository {
	return &carrierRepository{Conn}
}

func (r *carrierRepository) FetchByID(ctx context.Context, id int) (*model.Carrier, error) {
	u := &model.Carrier{ID: id}
	result := r.Conn.First(u, id)

	for result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}
