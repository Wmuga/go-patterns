package models

import "context"

type Item struct {
	ID   int
	Name string
}

type ItemRepo interface {
	Get(ctx context.Context, id int) (*Item, error)
	Put(ctx context.Context, item *Item) error
}
