package user

import "estatehub-api/internal/platform/shared"

var allowedUserFieldMask = map[string]struct{}{
	"id":         {},
	"email":      {},
	"type":       {},
	"created_at": {},
	"updated_at": {},
}

func ApplyUserFieldMask(user *UserResponse, fieldMask shared.FieldMask) (map[string]any, *shared.APIError) {
	if err := fieldMask.Validate(allowedUserFieldMask); err != nil {
		return nil, err
	}

	fullUser := map[string]interface{}{
		"id":         user.ID,
		"name":       user.Name,
		"email":      user.Email,
		"type":       user.Type,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}

	if fieldMask.ShouldReturnAll() {
		return fullUser, nil
	}

	partialUser := make(map[string]interface{})

	for _, field := range fieldMask {
		partialUser[field] = fullUser[field]
	}

	return partialUser, nil
}
