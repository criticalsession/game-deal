package api

import (
	"sync"
	"time"

	"github.com/criticalsession/game-deal/internal/types/stores"
)

type storeCache struct {
	entries map[int]stores.Store
	mux     *sync.Mutex
	expires time.Time
}

func NewStoreCache() *storeCache {
	s := &storeCache{
		entries: make(map[int]stores.Store),
		mux:     &sync.Mutex{},
	}

	go s.reapLoop()

	return s
}

func (c *storeCache) Get(id int) (stores.Store, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	entry, ok := c.entries[id]
	if !ok {
		return stores.Store{}, false
	}

	return entry, true
}

func (c *storeCache) Set(s map[int]stores.Store) {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.entries = s
}

func (c *storeCache) reapLoop() {
	for {
		time.Sleep(30 * time.Minute)

		c.mux.Lock()
		if time.Now().After(c.expires) {
			c.entries = make(map[int]stores.Store)
		}
		c.mux.Unlock()
	}
}
