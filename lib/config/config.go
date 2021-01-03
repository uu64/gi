package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// Config contains all configurations used by gi.
type Config struct {
	Remote remoteConfig `mapstructure:"remote"`
	Tui    tuiConfig    `mapstructure:"tui"`
}

type remoteConfig struct {
	Owner      string `mapstructure:"owner"`
	Repository string `mapstructure:"repository"`
	Ref        string `mapstructure:"ref"`
}

type tuiConfig struct {
	PageSize int `mapstructure:"pagesize"`
}

const (
	appName           = "gi"
	defaultOwner      = "github"
	defaultRepository = "gitignore"
	defaultRef        = "master"
	defaultPageSize   = 20
)

var c Config

func init() {
	viper.AddConfigPath(filepath.Join(getConfigDir(), appName))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Set default values
	viper.SetDefault("remote.owner", defaultOwner)
	viper.SetDefault("remote.repository", defaultRepository)
	viper.SetDefault("remote.ref", defaultRef)
	viper.SetDefault("tui.pagesize", defaultPageSize)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// ignore
		} else {
			fmt.Printf("%+v", err)
			os.Exit(1)
		}
	}

	err := viper.Unmarshal(&c)
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
}

func getConfigDir() string {
	var configDir string

	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		configDir = os.Getenv("APPDATA")
	} else {
		configDir = filepath.Join(home, ".config")
	}

	return configDir
}

// Get returns a Config object
func Get() *Config {
	return &c
}
