package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type GormConf struct {
	Driver        string
	Setting       string
	SingularTable bool
	LogMode       bool
}

var X *gorm.DB

func (gc *GormConf) InitGorm() error {
	var err error
	//"root:root@tcp(localhost:3306)/plucron?charset=utf8&parseTime=True&loc=Local"
	X, err = gorm.Open(gc.Driver, gc.Setting)
	if err != nil {
		return err
	}
	X.SingularTable(gc.SingularTable)
	X.LogMode(gc.LogMode)
	return nil
}
