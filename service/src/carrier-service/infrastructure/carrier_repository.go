package infrastructure

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
)

type carrierRepository struct {
	Conn *sql.DB
}

func NewCarrierRepository(Conn *sql.DB) repository.carrierRepository {
	return &carrierRepository{Conn}
}

func (r *carrierRepository) FetchByID(ctx fiber.Ctx, id int) (*model.Carrier, error) {
	u := &model.Carrier{ID: id}
	rows, err := r.Conn.Query("SELECT * FROM carrier")
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(u)
		if err != nil {
			return nil, err
		}
	}

	return u, err
}
