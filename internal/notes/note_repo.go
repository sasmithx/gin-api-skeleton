package notes

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// Repo -> data access layer

var ErrNoteNotFound = errors.New("note not found")

type Repository interface {
	Create(ctx context.Context, req CreateNoteRequest) (Note, error)
	GetByID(ctx context.Context, id string) (Note, error)
	List(ctx context.Context) ([]Note, error)
	Update(ctx context.Context, id string, req CreateNoteRequest) (Note, error)
	Delete(ctx context.Context, id string) error
}

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(pool *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{pool: pool}
}

func (r *PostgresRepository) Create(ctx context.Context, req CreateNoteRequest) (Note, error) {
	const q = `
		INSERT INTO notes (title, content, pinned)
		VALUES ($1, $2, $3)
		RETURNING id, title, content, pinned, created_at, updated_at
	`

	var note Note
	err := r.pool.QueryRow(ctx, q, req.Title, req.Content, req.Pinned).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Pinned,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		return Note{}, fmt.Errorf("create note: %w", err)
	}

	return note, nil
}

func (r *PostgresRepository) GetByID(ctx context.Context, id string) (Note, error) {
	const q = `
		SELECT id, title, content, pinned, created_at, updated_at
		FROM notes
		WHERE id = $1
	`

	var note Note
	err := r.pool.QueryRow(ctx, q, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Pinned,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Note{}, ErrNoteNotFound
		}
		return Note{}, fmt.Errorf("get note by id: %w", err)
	}

	return note, nil
}

func (r *PostgresRepository) List(ctx context.Context) ([]Note, error) {
	const q = `
		SELECT id, title, content, pinned, created_at, updated_at
		FROM notes
		ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("list notes: %w", err)
	}
	defer rows.Close()

	notes := make([]Note, 0)
	for rows.Next() {
		var note Note
		if err := rows.Scan(
			&note.ID,
			&note.Title,
			&note.Content,
			&note.Pinned,
			&note.CreatedAt,
			&note.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan note row: %w", err)
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate notes rows: %w", err)
	}

	return notes, nil
}

func (r *PostgresRepository) Update(ctx context.Context, id string, req CreateNoteRequest) (Note, error) {
	const q = `
		UPDATE notes
		SET title = $2,
		    content = $3,
		    pinned = $4,
		    updated_at = NOW()
		WHERE id = $1
		RETURNING id, title, content, pinned, created_at, updated_at
	`

	var note Note
	err := r.pool.QueryRow(ctx, q, id, req.Title, req.Content, req.Pinned).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.Pinned,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return Note{}, ErrNoteNotFound
		}
		return Note{}, fmt.Errorf("update note: %w", err)
	}

	return note, nil
}

func (r *PostgresRepository) Delete(ctx context.Context, id string) error {
	const q = `
		DELETE FROM notes
		WHERE id = $1
	`

	result, err := r.pool.Exec(ctx, q, id)
	if err != nil {
		return fmt.Errorf("delete note: %w", err)
	}

	if result.RowsAffected() == 0 {
		return ErrNoteNotFound
	}

	return nil
}

