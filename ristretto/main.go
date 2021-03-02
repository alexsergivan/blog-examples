package main

import (
	"github.com/alexsergivan/blog-examples/ristretto/repository"
	"github.com/dgraph-io/ristretto"
)

func main() {
	ristrettoCache, _ := ristretto.NewCache(&ristretto.Config{
		NumCounters: 1e7,     // Num keys to track frequency of (10M).
		MaxCost:     1 << 30, // Maximum cost of cache (1GB).
		BufferItems: 64,      // Number of keys per Get buffer.
	})

	for i:=0; i<100; i ++ {
		UsersGetter(repository.NewInMemoryRepository(ristrettoCache))
	}
}

func UsersGetter(repository repository.Repository) {
	repository.GetUsers()
}