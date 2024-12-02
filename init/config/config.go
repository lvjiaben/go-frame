package config

import (
	"fmt"

	"github.com/lvjiaben/go-wheel/tools"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config map[string]interface{}

func (config *Config) Load() *Config {
	path, err := tools.GetRootDir()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("system")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件已修改")
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Println(err)
		}
	})
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	return config
}
