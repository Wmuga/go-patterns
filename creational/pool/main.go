package main

import (
	"sync"

	"github.com/Wmuga/go-patterns/models"
)

type Pool struct {
	pool sync.Pool
}

func (p *Pool) Get() *models.Item {
	item := p.pool.Get()
	return item.(*models.Item)
}

func (p *Pool) Return(i *models.Item) {
	p.pool.Put(i)
}

func NewPool() *Pool {
	return &Pool{
		pool: sync.Pool{
			New: func() any {
				return &models.Item{}
			},
		},
	}
}

func main() {
	p := NewPool()
	item := p.Get()
	// do smth
	p.Return(item)
}
