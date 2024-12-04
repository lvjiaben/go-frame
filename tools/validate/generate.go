package validate

import (
	"strings"

	gorm2 "github.com/lvjiaben/go-wheel/tools/gorm"

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
	tables := gorm2.GetTables(db, tableNames, viper.Conf.Mysql.Dbname)
	for _, table := range tables {
		fields := gorm2.GetFields(db, table.Name)
		generateValidate(PackageName, Path, Cover, table, fields)
	}
}

func getRequire(null string) string {
	if null == "YES" {
		return "-"
	} else {
		return "required"
	}
}

// 生成Model
func generateValidate(PackageName string, Path string, Cover bool, table gorm2.Table, fields []gorm2.Field) {

	var builder strings.Builder
	builder.WriteString("package " + PackageName + "\n\n")
	list := []string{"Create", "Update", "Delete", "Sort"}
	for _, item := range list {
		builder.WriteString("type " + util.Marshal(table.Name) + item + " struct {\n")
		for _, field := range fields {
			fieldName := field.Field
			if (item == "Create" || item == "Update") && !util.IsInSlice(fieldName, []string{"created_at", "create_time", "updated_at", "update_time", "deleted_at", "delete_time"}) {
				if item == "Create" && field.Key == "PRI" {
					continue
				}
				builder.WriteString("\t" + util.Marshal(fieldName) + "\t" + gorm2.GetFiledType(field) + "\t" +
					"`" + "json:\"" + fieldName + "\" binding:\"" + getRequire(field.Null) + "\" msg:\"" + gorm2.GetFieldZh(field) + "有误\"" + "`\n")
			}
			if item == "Delete" && field.Key == "PRI" {
				builder.WriteString("\t" + util.Marshal(fieldName) + "\t" + gorm2.GetFiledType(field) + "\t" +
					"`" + "json:\"" + fieldName + "\" binding:\"" + getRequire(field.Null) + "\" msg:\"" + gorm2.GetFieldZh(field) + "有误\"" + "`\n")
			}
		}
		builder.WriteString("}\n\n")
	}
	// 文件生成
	fileName := Path + table.Name + ".go"
	file.MakeFile(Path, fileName, builder.String(), Cover)
}
