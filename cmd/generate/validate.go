package main

import (
	"flag"

	"github.com/lvjiaben/go-wheel/tools/validate"
)

// go run cmd/generate/validate.go -package_name=model -table_name=admin -path=/usr/local/var/golang/wheel/ -cover=true
func main() {
	var (
		TableName   string
		PackageName string
		Path        string
		Cover       bool
	)
	flag.StringVar(&TableName, "table_name", "", "数据表名")
	flag.StringVar(&PackageName, "package_name", "", "Package名称")
	flag.StringVar(&Path, "path", "", "路径地址")
	flag.BoolVar(&Cover, "cover", false, "是否覆盖")
	flag.Parse()
	validate.Genertate(TableName, PackageName, Path, Cover)
}
