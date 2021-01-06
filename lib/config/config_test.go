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
repo:
  owner: uu64
  name: gi
  branch: feature/test
`)
var incorrectConfig = []byte(`incorrect format`)

func TestGet(t *testing.T) {
	t.Run("can get default values", func(t *testing.T) {
		err := reload(bytes.NewBuffer(emptyConfig))
		if assert.NoError(t, err) {
			cfg := Get()
			assert.Equal(t, defaultOwner, cfg.Repo.Owner)
			assert.Equal(t, defaultName, cfg.Repo.Name)
			assert.Equal(t, defaultBranch, cfg.Repo.Branch)
			assert.Equal(t, defaultPageSize, cfg.Tui.PageSize)
			assert.Equal(t, defaultToken, cfg.Auth.Token)
		}
	})

	t.Run("can load value in config file", func(t *testing.T) {
		err := reload(bytes.NewBuffer(correctConfig))
		if assert.NoError(t, err) {
			cfg := Get()
			assert.Equal(t, "uu64", cfg.Repo.Owner)
			assert.Equal(t, "gi", cfg.Repo.Name)
			assert.Equal(t, "feature/test", cfg.Repo.Branch)
			assert.Equal(t, 30, cfg.Tui.PageSize)
			assert.Equal(t, "0123456789abc", cfg.Auth.Token)
		}
	})

	t.Run("cannot load incorrect yaml format", func(t *testing.T) {
		err := reload(bytes.NewBuffer(incorrectConfig))
		assert.Error(t, err)
	})
}
