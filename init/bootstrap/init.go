package bootstrap

import (
	"fmt"

	"github.com/lvjiaben/go-wheel/init/viper"
)

func init() {
	viper.Load()
	fmt.Println(2)
}
