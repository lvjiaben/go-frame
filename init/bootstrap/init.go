package bootstrap

import (
	"fmt"

	"github.com/lvjiaben/go-wheel/init/viper"
)

func init() {
	conf := viper.Load()
	fmt.Println(2)
	fmt.Println(conf)
}
