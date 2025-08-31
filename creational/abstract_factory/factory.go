package main

import (
	"github.com/Wmuga/go-patterns/models"
	"github.com/Wmuga/go-patterns/models/item_repo/postgres"
	"github.com/Wmuga/go-patterns/models/item_repo/sqlite"
)

type PostgresConfig struct {
	ConnectionString string
	CreateTable      bool
}

type SqliteConfig struct {
	ConnectionString string
	CreateTable      bool
}

type StorageConfig struct {
	Postgres *PostgresConfig
	Sqlite   *SqliteConfig
}

type StorageFactory struct {
	config *StorageConfig
}

func (s *StorageFactory) MustGetRepo(storage string) models.ItemRepo {
	switch storage {
	case "postgres", "pg":
		return postgres.MustNew("postgresql://postgres:postgres@127.0.0.1/postgres?ssl_mode=disable", true)
	case "sqlite", "sqlite3":
		return sqlite.MustNew("sqlite.db", true)
	default:
		panic("unknown storage " + storage)
	}
}

func FactoryFromConfig(cfg *StorageConfig) *StorageFactory {
	return &StorageFactory{config: cfg}
}
