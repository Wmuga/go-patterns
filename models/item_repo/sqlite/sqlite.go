package sqlite

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Wmuga/go-patterns/models"

	_ "modernc.org/sqlite"
)

type Sqlite struct {
	db *sql.DB
}

var _ models.ItemRepo = (*Sqlite)(nil)

// Get implements models.ItemRepo.
func (s *Sqlite) Get(ctx context.Context, id int) (*models.Item, error) {
	row := s.db.QueryRowContext(ctx, `SELECT id, name FROM items WHERE id = $1;`, id)
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
func (s *Sqlite) Put(ctx context.Context, item *models.Item) error {
	var err error
	if item.ID == 0 {
		item.ID, err = s.getNextId(ctx)
		if err != nil {
			return err
		}
	}

	_, err = s.db.ExecContext(ctx, `INSERT OR REPLACE INTO items(id,name) VALUES($1, $2);`, item.ID, item.Name)
	if err != nil {
		return err
	}

	return nil
}

func (s *Sqlite) getNextId(ctx context.Context) (int, error) {
	var (
		id  int
		err error
	)

	row := s.db.QueryRowContext(ctx, `SELECT id FROM items ORDER BY id DESC LIMIT 1;`)
	if err = row.Scan(&id); err == nil {
		return id + 1, nil
	}

	if errors.Is(err, sql.ErrNoRows) {
		return 1, nil
	}

	return 0, err
}

func MustNew(constr string, createTables bool) *Sqlite {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	db, err := sql.Open("sqlite", constr)
	if err != nil {
		panic(err)
	}
	if err = db.PingContext(ctx); err != nil {
		panic(err)
	}

	if !createTables {
		return &Sqlite{db: db}
	}

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXITST items(
		ID INTEGER PRIMARY KEY,
		name TEXT
	);`)

	if err != nil {
		panic(err)
	}

	return &Sqlite{db: db}
}
