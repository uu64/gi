package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

// Config contains all configurations used by gi.
type Config struct {
	Auth authConfig `mapstructure:"auth"`
	Repo repoConfig `mapstructure:"repo"`
	Tui  tuiConfig  `mapstructure:"tui"`
}

type authConfig struct {
	Token string `mapstructure:"token"`
}

type repoConfig struct {
	Owner  string `mapstructure:"owner"`
	Name   string `mapstructure:"name"`
	Branch string `mapstructure:"branch"`
}

type tuiConfig struct {
	PageSize int `mapstructure:"pagesize"`
}

const (
	appName         = "gi"
	defaultOwner    = "github"
	defaultName     = "gitignore"
	defaultBranch   = "master"
	defaultPageSize = 20
	defaultToken    = ""
)

var c Config

func init() {
	viper.AddConfigPath(filepath.Join(getConfigDir(), appName))
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	// Set default values
	viper.SetDefault("auth.token", defaultToken)
	viper.SetDefault("repo.owner", defaultOwner)
	viper.SetDefault("repo.name", defaultName)
	viper.SetDefault("repo.branch", defaultBranch)
	viper.SetDefault("tui.pagesize", defaultPageSize)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// ignore
		} else {
			fmt.Printf("%+v", err)
			os.Exit(1)
		}
	}

	if err := viper.Unmarshal(&c); err != nil {
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

// This method reloads config from buffer for testing.
func reload(buf *bytes.Buffer) error {
	if err := viper.ReadConfig(buf); err != nil {
		return err
	}

	if err := viper.Unmarshal(&c); err != nil {
		return err
	}

	return nil
}
