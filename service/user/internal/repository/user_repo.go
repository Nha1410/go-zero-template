package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Nha1410/go-zero-template/service/user/internal/domain/entity"
	"github.com/Nha1410/go-zero-template/service/user/internal/domain/repository"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ repository.UserRepository = (*userRepo)(nil)

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) repository.UserRepository {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *entity.User) error {
	query := `
		INSERT INTO users (email, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, user.Email, user.Name, user.CreatedAt, user.UpdatedAt).Scan(&user.ID)
	if err != nil {
		logx.Errorf("Failed to create user: %v", err)
		return err
	}

	return nil
}

func (r *userRepo) GetByID(ctx context.Context, id int64) (*entity.User, error) {
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE id = $1
	`

	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		logx.Errorf("Failed to get user by ID: %v", err)
		return nil, err
	}

	return user, nil
}

func (r *userRepo) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		WHERE email = $1
	`

	user := &entity.User{}
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Name,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		logx.Errorf("Failed to get user by email: %v", err)
		return nil, err
	}

	return user, nil
}

func (r *userRepo) List(ctx context.Context, page, pageSize int64) ([]*entity.User, int64, error) {
	var total int64
	countQuery := `SELECT COUNT(*) FROM users`
	err := r.db.QueryRowContext(ctx, countQuery).Scan(&total)
	if err != nil {
		logx.Errorf("Failed to count users: %v", err)
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	query := `
		SELECT id, email, name, created_at, updated_at
		FROM users
		ORDER BY created_at DESC
		LIMIT $1 OFFSET $2
	`

	rows, err := r.db.QueryContext(ctx, query, pageSize, offset)
	if err != nil {
		logx.Errorf("Failed to list users: %v", err)
		return nil, 0, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		user := &entity.User{}
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			logx.Errorf("Failed to scan user: %v", err)
			continue
		}
		users = append(users, user)
	}

	return users, total, nil
}

func (r *userRepo) Update(ctx context.Context, user *entity.User) error {
	query := `
		UPDATE users
		SET email = $1, name = $2, updated_at = $3
		WHERE id = $4
	`

	result, err := r.db.ExecContext(ctx, query, user.Email, user.Name, user.UpdatedAt, user.ID)
	if err != nil {
		logx.Errorf("Failed to update user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		logx.Errorf("Failed to delete user: %v", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("user not found")
	}

	return nil
}
