package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/spf13/viper"
)

type config struct {
	Remote RemoteConfig `mapstructure:"remote"`
	Tui    TuiConfig    `mapstructure:"tui"`
}

type RemoteConfig struct {
	Owner      string `mapstructure:"owner"`
	Repository string `mapstructure:"repository"`
	Ref        string `mapstructure:"ref"`
}

type TuiConfig struct {
	PageSize int `mapstructure:"pagesize"`
}

const (
	defaultOwner      = "github"
	defaultRepository = "gitignore"
	defaultRef        = "master"
	defaultPageSize   = 20
)

var c config

func init() {
	// TODO: app name should be managed
	viper.AddConfigPath(filepath.Join(getConfigDir(), "gi"))

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

func GetRemoteConfig() RemoteConfig {
	return c.Remote
}

func GetTuiConfig() TuiConfig {
	return c.Tui
}
