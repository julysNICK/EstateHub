package lead

import (
	"context"
	"database/sql"
)

type LeadRepository interface {
	Create(ctx context.Context, lead *Lead) (Lead, error)
	List(ctx context.Context) ([]Lead, error)
	GetByID(ctx context.Context, id string) (Lead, error)
	Update(ctx context.Context, lead *Lead) (Lead, error)
	Delete(ctx context.Context, id string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, lead *Lead) (Lead, error) {
	query := `
		INSERT INTO leads (id, name, email, phone, property_id, agency_id, broker_id, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, name, email, phone, property_id, agency_id, broker_id, created_at, updated_at
	`

	var createdLead Lead
	err := r.db.QueryRowContext(ctx, query, lead.ID, lead.Name, lead.Email, lead.Phone, lead.PropertyID, lead.AgencyID, lead.BrokerID, lead.CreatedAt, lead.UpdatedAt).Scan(
		&createdLead.ID,
		&createdLead.Name,
		&createdLead.Email,
		&createdLead.Phone,
		&createdLead.PropertyID,
		&createdLead.AgencyID,
		&createdLead.BrokerID,
		&createdLead.CreatedAt,
		&createdLead.UpdatedAt,
	)

	if err != nil {
		return Lead{}, err
	}

	return createdLead, nil

}

func (r *Repository) GetByID(ctx context.Context, id string) (Lead, error) {
	query := `
		SELECT id, name, email, phone, property_id, agency_id, broker_id, created_at, updated_at
		FROM leads
		WHERE id = $1
	`
	var lead Lead
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&lead.ID,
		&lead.Name,
		&lead.Email,
		&lead.Phone,
		&lead.PropertyID,
		&lead.AgencyID,
		&lead.BrokerID,
		&lead.CreatedAt,
		&lead.UpdatedAt,
	)

	if err != nil {
		return Lead{}, err
	}

	return lead, nil
}

func (r *Repository) List(ctx context.Context) ([]Lead, error) {
	query := `
		SELECT id, name, email, phone, property_id, agency_id, broker_id, created_at, updated_at
		FROM leads
	`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var leads []Lead

	for rows.Next() {
		var lead Lead
		err := rows.Scan(
			&lead.ID,
			&lead.Name,
			&lead.Email,
			&lead.Phone,
			&lead.PropertyID,
			&lead.AgencyID,
			&lead.BrokerID,
			&lead.CreatedAt,
			&lead.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		leads = append(leads, lead)
	}

	return leads, nil
}

func (r *Repository) Update(ctx context.Context, lead *Lead) (Lead, error) {
	query := `
		UPDATE leads
		SET name = $1, email = $2, phone = $3, property_id = $4, agency_id = $5, broker_id = $6, updated_at = $7
		WHERE id = $8
		RETURNING id, name, email, phone, property_id, agency_id, broker_id, created_at, updated_at
	`

	var updatedLead Lead
	err := r.db.QueryRowContext(ctx, query, lead.Name, lead.Email, lead.Phone, lead.PropertyID, lead.AgencyID, lead.BrokerID, lead.UpdatedAt, lead.ID).Scan(
		&updatedLead.ID,
		&updatedLead.Name,
		&updatedLead.Email,
		&updatedLead.Phone,
		&updatedLead.PropertyID,
		&updatedLead.AgencyID,
		&updatedLead.BrokerID,
		&updatedLead.CreatedAt,
		&updatedLead.UpdatedAt,
	)

	if err != nil {
		return Lead{}, err
	}

	return updatedLead, nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM leads WHERE id = $1
	`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
