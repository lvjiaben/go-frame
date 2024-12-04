package gorm

import (
	"strings"

	"github.com/lvjiaben/go-wheel/init/viper"

	"github.com/lvjiaben/go-wheel/init/mysql"
	"github.com/lvjiaben/go-wheel/pkg/file"

	"github.com/lvjiaben/go-wheel/pkg/util"
)

// 初始化需要传入一个model
func Genertate(TableName string, PackageName string, Path string, Cover bool) {
	tableNames := strings.Split(TableName, ",")
	viper.Load()
	mysql.Load()
	defer mysql.Close()
	db := mysql.Db
	tables := GetTables(db, tableNames, viper.Conf.Mysql.Dbname)
	for _, table := range tables {
		fields := GetFields(db, table.Name)
		generateModel(PackageName, Path, Cover, table, fields)
	}
}

// 生成Model
func generateModel(PackageName string, Path string, Cover bool, table Table, fields []Field) {

	var builder strings.Builder
	builder.WriteString("package " + PackageName + "\n\n")

	// 表注释
	if len(table.Comment) > 0 {
		builder.WriteString("// " + table.Comment + "\n")
	}

	// 生成结构体
	builder.WriteString("type " + util.Marshal(table.Name) + " struct {\n")

	// 文件内容填充
	for _, field := range fields {
		fieldName := field.Field
		/**
		字段名 字段类型 `json:"字段名" gorm:"column:字段名"` //注释
		*/
		builder.WriteString("\t" + util.Marshal(fieldName) + "\t" + GetFiledType(field) + "\t" +
			"`" + GetFieldJson(field) + "`\t" + GetFieldComment(field) + "\n")
	}
	builder.WriteString("}\n")

	// 函数名称返回自身
	/**
	func (e *结构体名) TableName() string {
	    return 结构体名
	}
	*/
	builder.WriteString("func (e *" + util.Marshal(table.Name) +
		") TableName() string { \n    return \"" + table.Name + "\"\n}")

	// 文件生成
	fileName := Path + table.Name + ".go"
	file.MakeFile(Path, fileName, builder.String(), Cover)
}
