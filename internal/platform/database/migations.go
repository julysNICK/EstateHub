package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func RunMigration(ctx context.Context, db *sql.DB, migrationsDir string) error {

	if err := ensureMigrationsTable(ctx, db); err != nil {
		return err
	}

	files, err := filepath.Glob(filepath.Join(migrationsDir, "*.sql"))
	if err != nil {
		return err
	}

	sort.Strings(files)
	for _, file := range files {
		version := filepath.Base(file)

		applied, err := wasMigrationApplied(ctx, db, version)
		if err != nil {
			return err
		}

		if applied {
			continue
		}

		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		if strings.TrimSpace(string(content)) == "" {
			continue
		}

		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}

		if _, err := tx.ExecContext(ctx, string(content)); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("migration %s failed: %w", version, err)
		}

		if _, err := tx.ExecContext(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, version); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("could not register migration %s: %w", version, err)
		}

		if err := tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func ensureMigrationsTable(ctx context.Context, db *sql.DB) error {
	_, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version TEXT PRIMARY KEY,
			applied_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)

	return err
}

func wasMigrationApplied(ctx context.Context, db *sql.DB, version string) (bool, error) {
	var exists bool

	err := db.QueryRowContext(
		ctx,
		`SELECT EXISTS(SELECT 1 FROM schema_migrations WHERE version = $1)`,
		version,
	).Scan(&exists)

	if err != nil {
		return false, err
	}

	return exists, nil
}
