package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Wmuga/go-patterns/models"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	db *pgxpool.Pool
}

var _ models.ItemRepo = (*Postgres)(nil)

// Get implements models.ItemRepo.
func (s *Postgres) Get(ctx context.Context, id int) (*models.Item, error) {
	row := s.db.QueryRow(ctx, `SELECT id, name FROM items WHERE id = $1;`, id)
	item := new(models.Item)

	var err error
	if err = row.Scan(&item.ID, &item.Name); err == nil {
		return item, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: %d", models.ErrorNotFound, id)
	}

	return nil, err
}

// Put implements models.ItemRepo.
func (s *Postgres) Put(ctx context.Context, item *models.Item) error {
	var err error
	if item.ID == 0 {
		item.ID, err = s.getNextId(ctx)
		if err != nil {
			return err
		}
	}

	_, err = s.db.Exec(ctx, `INSERT INTO items(id,name) VALUES($1, $2) ON CONFLICT(id) DO UPDATE name = excluded.name;`, item.ID, item.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Postgres) getNextId(ctx context.Context) (int, error) {
	var (
		id  int
		err error
	)

	row := s.db.QueryRow(ctx, `SELECT id FROM items ORDER BY id DESC LIMIT 1;`)
	if err = row.Scan(&id); err == nil {
		return id + 1, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return 1, nil
	}

	return 0, err
}

func MustNew(constr string, createTables bool) *Postgres {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	db, err := pgxpool.New(ctx, constr)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(ctx); err != nil {
		panic(err)
	}

	if !createTables {
		return &Postgres{db: db}
	}

	_, err = db.Exec(ctx, `CREATE TABLE IF NOT EXITST items(
		ID INTEGER PRIMARY KEY,
		name TEXT
	);`)

	if err != nil {
		panic(err)
	}

	return &Postgres{db: db}
}
