package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("can get default values", func(t *testing.T) {
		cfg := Get()
		assert.Equal(t, defaultOwner, cfg.Remote.Owner)
		assert.Equal(t, defaultRepository, cfg.Remote.Repository)
		assert.Equal(t, defaultRef, cfg.Remote.Ref)
		assert.Equal(t, defaultPageSize, cfg.Tui.PageSize)
	})
	// TODO: add test for values in config file
}
