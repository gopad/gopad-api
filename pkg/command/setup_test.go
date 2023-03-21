package command

import (
	"testing"

	"github.com/gopad/gopad-api/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestSetupLogger(t *testing.T) {
	assert.NoError(t, setupLogger(config.Load()))
}
