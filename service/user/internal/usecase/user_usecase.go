package usecase

import (
	"context"
	"time"

	common_errors "github.com/Nha1410/go-zero-template/common/errors"
	"github.com/Nha1410/go-zero-template/service/user/internal/domain/entity"
	"github.com/Nha1410/go-zero-template/service/user/internal/domain/repository"
)

// UserUsecase handles user business logic
type UserUsecase struct {
	userRepo repository.UserRepository
}

// NewUserUsecase creates a new user usecase
func NewUserUsecase(userRepo repository.UserRepository) *UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (uc *UserUsecase) CreateUser(ctx context.Context, email, name string) (*entity.User, error) {
	// Check if user already exists
	existing, _ := uc.userRepo.GetByEmail(ctx, email)
	if existing != nil {
		return nil, common_errors.ErrConflict.WithDetails("User with this email already exists")
	}

	// Create new user
	user := &entity.User{
		Email:     email,
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(ctx, user); err != nil {
		return nil, common_errors.ErrInternalError.WithDetails(err.Error())
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*entity.User, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, common_errors.ErrNotFound.WithDetails("User not found")
	}
	return user, nil
}

// GetUsers retrieves a list of users with pagination
func (uc *UserUsecase) GetUsers(ctx context.Context, page, pageSize int64) ([]*entity.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}

	users, total, err := uc.userRepo.List(ctx, page, pageSize)
	if err != nil {
		return nil, 0, common_errors.ErrInternalError.WithDetails(err.Error())
	}

	return users, total, nil
}

// UpdateUser updates an existing user
func (uc *UserUsecase) UpdateUser(ctx context.Context, id int64, email, name string) (*entity.User, error) {
	// Get existing user
	user, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, common_errors.ErrNotFound.WithDetails("User not found")
	}

	// Update fields if provided
	if email != "" {
		// Check if email is already taken by another user
		existing, _ := uc.userRepo.GetByEmail(ctx, email)
		if existing != nil && existing.ID != id {
			return nil, common_errors.ErrConflict.WithDetails("Email already taken")
		}
		user.Email = email
	}
	if name != "" {
		user.Name = name
	}
	user.UpdatedAt = time.Now()

	if err := uc.userRepo.Update(ctx, user); err != nil {
		return nil, common_errors.ErrInternalError.WithDetails(err.Error())
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (uc *UserUsecase) DeleteUser(ctx context.Context, id int64) error {
	// Check if user exists
	_, err := uc.userRepo.GetByID(ctx, id)
	if err != nil {
		return common_errors.ErrNotFound.WithDetails("User not found")
	}

	if err := uc.userRepo.Delete(ctx, id); err != nil {
		return common_errors.ErrInternalError.WithDetails(err.Error())
	}

	return nil
}
