package repository

import (
	"fmt"
	"github.com/dgraph-io/ristretto"
	"time"
)

type InMemoryRepository struct{
	cache        *ristretto.Cache
}

// NewInMemoryRepository constructs and returns InMemoryRepository.
func NewInMemoryRepository(ristrettoCache *ristretto.Cache) *InMemoryRepository {
	return &InMemoryRepository{
		cache: ristrettoCache,
	}
}

// GetUsers returns 50000 dummy users from the in-memory repository.
func (r *InMemoryRepository) GetUsers() map[int]string {
	key := "users"
	value, found := r.cache.Get(key)
	if !found {
		users := make(map[int]string)
		for i:=1; i <= 50000; i++ {
			users[i] = fmt.Sprintf("User %d", i)
		}

		r.cache.SetWithTTL(key, users, 1, 1*time.Hour)
		time.Sleep(10 * time.Millisecond)
		return users
	}
	return value.(map[int]string)
}