package store

import (
	"github.com/gopad/gopad-api/pkg/service/teams"
	"github.com/gopad/gopad-api/pkg/service/users"
	"github.com/pkg/errors"
)

var (
	// ErrUnknownDriver defines a named error for unknown store drivers.
	ErrUnknownDriver = errors.New("unknown database driver")
)

// Store provides the interface for the store implementations.
type Store interface {
	Info() map[string]interface{}
	Prepare() error
	Open() error
	Close() error
	Ping() error
	Migrate() error
	Admin(string, string, string) error
	Teams() teams.Store
	Users() users.Store
}
