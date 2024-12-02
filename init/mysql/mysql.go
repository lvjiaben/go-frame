package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/lvjiaben/go-wheel/init/viper"
)

var Db *gorm.DB

func Load() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
		viper.Conf.Mysql.User,
		viper.Conf.Mysql.Pass,
		viper.Conf.Mysql.Host,
		viper.Conf.Mysql.Port,
		viper.Conf.Mysql.Dbname,
		viper.Conf.Mysql.Charset,
	)
	var err error
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	sqlDb, err := Db.DB()
	sqlDb.SetMaxOpenConns(viper.Conf.Mysql.MaxOpenConns)
	sqlDb.SetMaxIdleConns(viper.Conf.Mysql.MaxIdleConns)
}

func Close() {
	sqlDb, _ := Db.DB()
	_ = sqlDb.Close()
}
