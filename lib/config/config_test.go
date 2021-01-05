package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	t.Run("can get default values", func(t *testing.T) {
		cfg := Get()
		assert.Equal(t, defaultOwner, cfg.Repo.Owner)
		assert.Equal(t, defaultName, cfg.Repo.Name)
		assert.Equal(t, defaultBranch, cfg.Repo.Branch)
		assert.Equal(t, defaultPageSize, cfg.Tui.PageSize)
		assert.Equal(t, defaultToken, cfg.Auth.Token)
	})
}
