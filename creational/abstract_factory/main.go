package main

import (
	"context"
	"fmt"

	"github.com/Wmuga/go-patterns/models"
)

func main() {
	factory := FactoryFromConfig(&StorageConfig{
		Postgres: &PostgresConfig{
			ConnectionString: "postgresql://postgres:postgres@127.0.0.1/postgres?ssl_mode=disable",
			CreateTable:      true,
		},
		Sqlite: &SqliteConfig{
			ConnectionString: "sqlite.db",
			CreateTable:      true,
		},
	})

	putGetItem(factory.MustGetRepo("postgres"))
	putGetItem(factory.MustGetRepo("sqlite"))
}

func putGetItem(st models.ItemRepo) {
	ctx := context.Background()
	item := &models.Item{ID: 1, Name: "test"}

	err := st.Put(ctx, item)
	if err != nil {
		panic(err)
	}

	itemGot, err := st.Get(ctx, item.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println(itemGot)
}
