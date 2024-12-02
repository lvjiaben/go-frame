package viper

import (
	"fmt"
	"path/filepath"

	"github.com/lvjiaben/go-wheel/pkg/file"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	App   App   `mapstructure:"app"`
	Log   Log   `mapstructure:"log"`
	Redis Redis `mapstructure:"redis"`
	Mysql Mysql `mapstructure:"mysql"`
}

var Conf = new(Config)

func Load() {
	path, err := file.GetRootDir()
	if err != nil {
		panic(err)
	}
	path = filepath.Join(path, "configs")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
		if err := viper.Unmarshal(Conf); err != nil {
			fmt.Println(err)
		}
	})
	if err := viper.Unmarshal(Conf); err != nil {
		panic(err)
	}
}
