package user

import (
	"context"
	"estatehub-api/internal/platform/shared"
	utils "estatehub-api/internal/utils"
	"strings"
)

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, req CreateUserRequest) (UserResponse, error) {
	if err := req.Validate(); err != nil {
		return UserResponse{}, err
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return UserResponse{}, ErrFailedToCreateUser
	}

	now := utils.Now()

	user := User{
		ID:           utils.GenerateID(),
		Name:         strings.TrimSpace(req.Name),
		Email:        strings.TrimSpace(req.Email),
		Type:         strings.TrimSpace(req.Type),
		PasswordHash: hashedPassword,
		CreatedAt:    now,
		UpdatedAt:    now,
	}

	createdUser, err := s.repo.Create(ctx, &user)
	if err != nil {
		return UserResponse{}, ErrFailedToCreateUser
	}

	return ToUserResponse(&createdUser), nil
}

func (s *UserService) List(ctx context.Context, req ListUserRequest, sort shared.SortRequest) (ListUserResponse, *shared.APIError) {
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)

	users, err := s.repo.List(ctx, req, sort)
	if err != nil {
		return ListUserResponse{}, shared.NewInternalServerError("failed to list users")
	}

	return ListUserResponse{
		Users:         make([]UserResponse, len(users.Item)),
		NextPageToken: users.NextPageToken,
	}, nil
}

func (s *UserService) GetByID(ctx context.Context, userID string) (UserResponse, error) {
	if strings.TrimSpace(userID) == "" {
		return UserResponse{}, ErrInvalidID
	}

	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(&user), nil
}

func (s *UserService) Update(ctx context.Context, userID string, req UpdateUserRequest) (UserResponse, error) {
	if strings.TrimSpace(userID) == "" {
		return UserResponse{}, ErrInvalidID
	}

	if err := req.Validate(); err != nil {
		return UserResponse{}, err
	}

	existingUser, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return UserResponse{}, err
	}

	if req.Name != nil {
		existingUser.Name = strings.TrimSpace(*req.Name)
	}

	if req.Email != nil {
		existingUser.Email = strings.TrimSpace(*req.Email)
	}

	if req.Type != nil {
		existingUser.Type = strings.TrimSpace(*req.Type)
	}

	if req.Password != nil {
		hashedPassword, err := utils.HashPassword(*req.Password)
		if err != nil {
			return UserResponse{}, ErrFailedToUpdateUser
		}

		existingUser.PasswordHash = hashedPassword
	}

	existingUser.UpdatedAt = utils.Now()

	updatedUser, err := s.repo.Update(ctx, &existingUser)
	if err != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(&updatedUser), nil
}

func (s *UserService) Delete(ctx context.Context, userID string) error {
	if strings.TrimSpace(userID) == "" {
		return ErrInvalidID
	}

	return s.repo.Delete(ctx, userID)
}
