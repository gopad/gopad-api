package upload

import (
	"net/http"

	"github.com/pkg/errors"
)

var (
	// ErrUnknownDriver defines a named error for unknown upload drivers.
	ErrUnknownDriver = errors.New("unknown upload driver")
)

// Upload provides the interface for the upload implementations.
type Upload interface {
	Close() error
	Handler() http.Handler
}
