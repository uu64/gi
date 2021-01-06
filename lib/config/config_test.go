package config

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var emptyConfig = []byte(``)
var correctConfig = []byte(`
tui:
  pagesize: 30
auth:
  token: 0123456789abc
repos:
  - owner: uu64
    name: gi
    branch: main
  - owner: uu64
    name: gi
    branch: feature/test
`)
var incorrectConfig = []byte(`incorrect format`)

func TestGet(t *testing.T) {
	t.Run("can get default values", func(t *testing.T) {
		err := reload(bytes.NewBuffer(emptyConfig))
		if assert.NoError(t, err) {
			cfg := Get()
			assert.Equal(t, defaultPageSize, cfg.Tui.PageSize)
			assert.Equal(t, defaultToken, cfg.Auth.Token)
			assert.Len(t, cfg.Repos, 1)
			assert.Equal(t, defaultOwner, cfg.Repos[0].Owner)
			assert.Equal(t, defaultName, cfg.Repos[0].Name)
			assert.Equal(t, defaultBranch, cfg.Repos[0].Branch)
		}
	})

	t.Run("can load value in config file", func(t *testing.T) {
		err := reload(bytes.NewBuffer(correctConfig))
		if assert.NoError(t, err) {
			cfg := Get()
			assert.Equal(t, 30, cfg.Tui.PageSize)
			assert.Equal(t, "0123456789abc", cfg.Auth.Token)
			assert.Len(t, cfg.Repos, 2)
			assert.Equal(t, "uu64", cfg.Repos[0].Owner)
			assert.Equal(t, "gi", cfg.Repos[0].Name)
			assert.Equal(t, "main", cfg.Repos[0].Branch)
			assert.Equal(t, "uu64", cfg.Repos[1].Owner)
			assert.Equal(t, "gi", cfg.Repos[1].Name)
			assert.Equal(t, "feature/test", cfg.Repos[1].Branch)
		}
	})

	t.Run("cannot load incorrect yaml format", func(t *testing.T) {
		err := reload(bytes.NewBuffer(incorrectConfig))
		assert.Error(t, err)
	})
}
