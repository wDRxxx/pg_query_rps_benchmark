package main

import (
	"context"
	"log"
	"sync"
)

func main() {
	ctx := context.Background()
	var wg sync.WaitGroup
	config := NewConfig()

	db, err := NewPostgreSQL(ctx, config.DSN())
	defer db.Close()

	if err != nil {
		log.Fatal("Error creating new database: ", err)
	}

	benchmarkCtx, cancel := context.WithTimeout(ctx, config.Duration())
	defer cancel()

	bench := NewBenchmark(db, config.Query(), &wg)
	result := bench.RunBenchmark(benchmarkCtx)

	log.Println("Benchmark finished with result:", result)
}
