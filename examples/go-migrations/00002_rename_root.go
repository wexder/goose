package main

import (
	"context"
	"database/sql"

	"github.com/wexder/goose/v3"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE users SET username='admin' WHERE username='root';")
	if err != nil {
		return err
	}
	return nil
}

func Down00002(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.Exec("UPDATE users SET username='root' WHERE username='admin';")
	if err != nil {
		return err
	}
	return nil
}
