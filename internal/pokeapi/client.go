package pokeapi

import (
	"net/http"
	"time"

	"github.com/The-Legolas/build_pokedex/internal/pokecache"
)

// Client -
type Client struct {
	Pokedex    map[string]Pokemon
	cache      *pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
