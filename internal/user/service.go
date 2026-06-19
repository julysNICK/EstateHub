package user

import (
	"context"
	"estatehub-api/internal/platform/shared"
	utils "estatehub-api/internal/utils"
	"net/http"
	nethttp "net/http"
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

func (s *UserService) Update(ctx context.Context, userID string, req UpdateUserRequest, fieldMask shared.FieldMask) (UserResponse, *shared.APIError) {

	if fieldMask.IsEmpty() {
		return UserResponse{}, &shared.APIError{
			StatusCode: nethttp.StatusBadRequest,
			ErrorBody: shared.ErrorBody{
				Code:    "invalid_field_mask",
				Message: "field_mask cannot be empty",
				Field:   "field_mask",
			},
		}

	}

	patch := UpdateUserRequest{}

	for _, field := range fieldMask {
		switch field {
		case "name":
			if req.Name == nil {
				return UserResponse{}, &shared.APIError{
					StatusCode: http.StatusBadRequest,
					ErrorBody: shared.ErrorBody{
						Code:    "missing_patch_field",
						Message: "Field mask contains field missing from body",
						Field:   "name",
					},
				}
			}

			if *req.Name == "" {
				return UserResponse{}, &shared.APIError{
					StatusCode: http.StatusBadRequest,
					ErrorBody: shared.ErrorBody{
						Code:    "invalid_name",
						Message: "Name cannot be empty",
						Field:   "name",
					},
				}
			}

			patch.Name = req.Name

		case "email":
			if req.Email == nil {
				return UserResponse{}, &shared.APIError{
					StatusCode: http.StatusBadRequest,
					ErrorBody: shared.ErrorBody{
						Code:    "missing_patch_field",
						Message: "Field mask contains field missing from body",
						Field:   "email",
					},
				}
			}

			if *req.Email == "" {
				return UserResponse{}, &shared.APIError{
					StatusCode: http.StatusBadRequest,
					ErrorBody: shared.ErrorBody{
						Code:    "invalid_email",
						Message: "Email cannot be empty",
						Field:   "email",
					},
				}
			}

			patch.Email = req.Email
		}

	}

	updatedUser, err := s.repo.Update(ctx, userID, patch, patch)
	if err != nil {
		return UserResponse{}, &shared.APIError{
			StatusCode: http.StatusBadRequest,
			ErrorBody: shared.ErrorBody{
				Code:    "failed_to_update_user",
				Message: "Failed to update user",
			},
		}
	}

	return ToUserResponse(&updatedUser), nil

}

func (s *UserService) Delete(ctx context.Context, userID string) error {
	if strings.TrimSpace(userID) == "" {
		return ErrInvalidID
	}

	return s.repo.Delete(ctx, userID)
}
