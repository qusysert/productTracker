package test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"productTracker/internal/app/pkg/config"
	"productTracker/internal/app/service"
	"productTracker/internal/pkg/db"
	"testing"
)

func doTest(t *testing.T, fn func(ctx context.Context)) {
	cfg, err := config.LoadConfig("../")
	if err != nil {
		log.Fatal("cannot load cfg:", err)
	}

	conn := db.New(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	ctx := db.AddToContext(context.Background(), &db.Db{conn})

	_, err = db.FromContext(ctx).Exec("begin;")
	assert.Nil(t, err)

	fn(ctx)

	_, err = db.FromContext(ctx).Exec("rollback;")
	assert.Nil(t, err)
}

func TestUpdateTable(t *testing.T) {
	doTest(t, func(ctx context.Context) {

		//
		f, err := os.Open("fixture/book1.xlsx")
		assert.Nil(t, err)
		products, err := service.TableToSlice(f, 234)

		assert.Nil(t, err)
		assert.Equal(t, len(products), 3)

		//
	})
}

func TestGetProductList(t *testing.T) {
	doTest(t, func(ctx context.Context) {

	})
}
