package user

import (
	"context"
	"database/sql"
	"estatehub-api/internal/platform/shared"
	"fmt"
)

type UserRepository interface {
	Create(ctx context.Context, user *User) (User, error)
	List(ctx context.Context, filter ListUserRequest, sort shared.SortRequest) (shared.PageResult[User], error)
	GetByID(ctx context.Context, userId string) (User, error)
	GetAgencies(ctx context.Context, limit int, offset int) ([]Agency, error)
	GetBrokers(ctx context.Context, limit int, offset int) ([]Broker, error)
	Update(ctx context.Context, user *User) (User, error)
	Delete(ctx context.Context, userId string) error
}

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, user *User) (User, error) {
	return User{}, nil
}

func (r *Repository) List(ctx context.Context, filter ListUserRequest, sort shared.SortRequest) (shared.PageResult[User], error) {
	offset, err := shared.DecodePageToken(filter.PageRequest.PageToken)
	if err != nil {
		return shared.PageResult[User]{}, err
	}

	query := fmt.Sprintf("SELECT id, name, email, type, created_at, updated_at FROM users WHERE name LIKE ? AND email LIKE ? ORDER BY %s %s LIMIT ? OFFSET ?", sort.Column, sort.Direction)

	rows, err := r.db.QueryContext(ctx, query, "%"+filter.Name+"%", "%"+filter.Email+"%", filter.PageRequest.PageSize, offset)
	if err != nil {
		return shared.PageResult[User]{}, err
	}
	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Type, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return shared.PageResult[User]{}, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return shared.PageResult[User]{}, err
	}

	nextPageToken := ""

	if len(users) > filter.PageRequest.PageSize {
		users = users[:filter.PageRequest.PageSize]

		next := offset + filter.PageRequest.PageSize

		token, err := shared.EncodePageToken(next)

		if err != nil {
			return shared.PageResult[User]{}, err
		}

		nextPageToken = token

	}

	return shared.PageResult[User]{
		Item:          users,
		NextPageToken: nextPageToken,
	}, nil
}

func (r *Repository) GetByID(ctx context.Context, userId string) (User, error) {
	return User{}, nil
}

func (r *Repository) GetAgencies(ctx context.Context, limit int, offset int) ([]Agency, error) {
	query := "SELECT id, user_id, agency_id, cnpj FROM agencies LIMIT ? OFFSET ?"

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	agencies := make([]Agency, 0)

	for rows.Next() {

		var agency Agency

		if err := rows.Scan(&agency.ID, &agency.UserID, &agency.AgencyID, &agency.Cnpj); err != nil {
			return nil, err
		}
		agencies = append(agencies, agency)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return agencies, nil
}

func (r *Repository) GetBrokers(ctx context.Context, limit int, offset int) ([]Broker, error) {
	query := "SELECT id, user_id, broker_id FROM brokers LIMIT ? OFFSET ?"

	rows, err := r.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	brokers := make([]Broker, 0)

	for rows.Next() {
		var broker Broker
		if err := rows.Scan(&broker.ID, &broker.UserID, &broker.BrokerID); err != nil {
			return nil, err
		}
		brokers = append(brokers, broker)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return brokers, nil
}

func (r *Repository) Update(ctx context.Context, user *User) (User, error) {
	return User{}, nil
}

func (r *Repository) Delete(ctx context.Context, userId string) error {
	return nil
}
