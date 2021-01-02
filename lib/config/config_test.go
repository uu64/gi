package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRemoteConfig(t *testing.T) {
	t.Run("can get default values", func(t *testing.T) {
		rc := GetRemoteConfig()
		assert.Equal(t, defaultOwner, rc.Owner)
		assert.Equal(t, defaultRepository, rc.Repository)
		assert.Equal(t, defaultRef, rc.Ref)
	})
	// TODO: add test for values in config file
}

func TestGetTuiConfig(t *testing.T) {
	t.Run("can get default values", func(t *testing.T) {
		tc := GetTuiConfig()
		assert.Equal(t, defaultPageSize, tc.PageSize)
	})
	// TODO: add test for values in config file
}
