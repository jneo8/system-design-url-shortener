package api

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"go.uber.org/dig"
)

// StoreOpts ...
type StoreOpts struct {
	dig.In
	Size     int      `name:"session_size"`
	Network  string   `name:"session_network"`
	Address  string   `name:"session_address"`
	Password string   `name:"session_password"`
	KeyPairs []string `name:"session_key_pairs"`
}

// NewRedisSessionStore return redis session store.
func NewRedisSessionStore(opts StoreOpts) (sessions.Store, error) {
	keyPairs := [][]byte{}
	for _, k := range opts.KeyPairs {
		keyPairs = append(keyPairs, []byte(k))
	}
	return redis.NewStore(opts.Size, opts.Network, opts.Address, opts.Password, keyPairs...)
}
