package repository

import (
	"context"
	"productTracker/internal/pkg/db"
)

func CreateSellerIfNotExists(ctx context.Context, sellerId int) error {
	_, err := db.FromContext(ctx).Exec(`INSERT INTO seller (id) VALUES ($1) ON CONFLICT (ID) DO NOTHING`, sellerId)
	return err
}

func LockSeller(ctx context.Context, sellerId int) error {
	_, err := db.FromContext(ctx).Query(`SELECT id FROM seller WHERE id = $1 FOR UPDATE`, sellerId)
	return err
}
