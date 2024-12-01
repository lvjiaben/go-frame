package bootstrap

import (
	"fmt"

	"github.com/lvjiaben/go-wheel/init/viper"
)

func init() {
	viper.Config
	fmt.Println(2)
}
