package property

import (
	"context"
	"database/sql"
)

type IPropertyRepository interface {
	Create(ctx context.Context, property *Property) (*Property, error)
	List(ctx context.Context) ([]Property, error)
	GetByID(ctx context.Context, propertyId string) (*Property, error)
	Update(ctx context.Context, property *Property) (*Property, error)
	Delete(ctx context.Context, propertyId string) error
}

type PropertyRepository struct {
	db *sql.DB
}

func NewPropertyRepository(db *sql.DB) *PropertyRepository {
	return &PropertyRepository{
		db: db,
	}
}

func (r *PropertyRepository) Create(ctx context.Context, property *Property) (*Property, error) {
	return nil, nil
}

func (r *PropertyRepository) List(ctx context.Context) ([]Property, error) {
	return []Property{}, nil
}

func (r *PropertyRepository) GetByID(ctx context.Context, propertyId string) (*Property, error) {
	return nil, nil
}

func (r *PropertyRepository) Update(ctx context.Context, property *Property) (*Property, error) {
	return nil, nil
}

func (r *PropertyRepository) Delete(ctx context.Context, propertyId string) error {
	return nil
}
