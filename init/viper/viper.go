package viper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

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
	files, err := getFilesInDir(path)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filename := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))
		if filename != "config" {
			viper.SetConfigName(filename)
			viper.AddConfigPath(path)
			err := viper.MergeInConfig()
			if err != nil {
				panic(err)
			}
		}
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	if err := viper.MergeInConfig(); err != nil {
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

func getFilesInDir(dir string) ([]os.DirEntry, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	return files, nil
}
