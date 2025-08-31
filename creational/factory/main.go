package main

import (
	"context"
	"fmt"

	"github.com/Wmuga/go-patterns/models"
	"github.com/Wmuga/go-patterns/models/item_repo/postgres"
	"github.com/Wmuga/go-patterns/models/item_repo/sqlite"
)

// mustGetRepo - factory method
func mustGetRepo(storage string) models.ItemRepo {
	switch storage {
	case "postgres", "pg":
		return postgres.MustNew("postgresql://postgres:postgres@127.0.0.1/postgres?ssl_mode=disable", true)
	case "sqlite", "sqlite3":
		return sqlite.MustNew("sqlite.db", true)
	default:
		panic("unknown storage " + storage)
	}
}

func main() {
	putGetItem(mustGetRepo("postgres"))
	putGetItem(mustGetRepo("sqlite"))
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
