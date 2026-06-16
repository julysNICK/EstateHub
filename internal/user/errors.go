package user

import "errors"

var (
	ErrFailedToCreateUser   = errors.New("failed to create user")
	ErrFailedToGetUsers     = errors.New("failed to get users")
	ErrFailedToGetAgencies  = errors.New("failed to get agencies")
	ErrFailedToGetBrokers   = errors.New("failed to get brokers")
	ErrUserNotFound         = errors.New("user not found")
	ErrFailedToHashPassword = errors.New("failed to hash password")
	ErrInvalidID            = errors.New("invalid user ID")
	ErrInvalidName          = errors.New("invalid name")
	ErrInvalidEmail         = errors.New("invalid email")
	ErrInvalidType          = errors.New("invalid user type")
	ErrInvalidPassword      = errors.New("invalid password")
	ErrInvalidUserType      = errors.New("invalid user type")
	ErrInvalidCNPJ          = errors.New("invalid CNPJ")
	ErrInvalidCPF           = errors.New("invalid CPF")
	ErrInvalidCRECI         = errors.New("invalid CRECI")
	ErrEmailAlreadyUsed     = errors.New("email already used")
	ErrFailedToUpdateUser   = errors.New("failed to update user")
	ErrFailedToDeleteUser   = errors.New("failed to delete user")
)
