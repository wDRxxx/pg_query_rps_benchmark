package main

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Benchmark struct {
	wg *sync.WaitGroup

	db    *pgxpool.Pool
	query string
}

func NewBenchmark(db *pgxpool.Pool, query string, wg *sync.WaitGroup) *Benchmark {
	return &Benchmark{
		db:    db,
		query: query,
		wg:    wg,
	}
}

func (b *Benchmark) RunBenchmark(ctx context.Context) uint64 {
	var rps atomic.Uint64

	for {
		select {
		case <-ctx.Done():
			b.wg.Wait()
			return rps.Load()
		default:
			b.wg.Add(1)
			go func() {
				defer b.wg.Done()

				if ctx.Err() != nil {
					return
				}

				_, err := b.db.Exec(ctx, b.query)
				if err != nil {
					return
				}

				rps.Add(1)
			}()
		}
	}
}
