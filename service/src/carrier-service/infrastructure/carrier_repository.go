package infrastructure

import (
	"carrier-service/domain/model"
	"carrier-service/domain/repository"
	"context"
	"gorm.io/gorm"
	"sync"
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

func (r *carrierRepository) Create(ctx context.Context, carrier *model.Carrier, err chan<- error, wg *sync.WaitGroup) {
	result := r.Conn.Create(carrier)
	if result.Error != nil {
		err <- result.Error
	}
	wg.Done()
}

func (r *carrierRepository) Update(ctx context.Context, carrier *model.Carrier) error {
	result := r.Conn.Save(carrier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *carrierRepository) Delete(ctx context.Context, carrier *model.Carrier) error {
	result := r.Conn.Delete(carrier)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
