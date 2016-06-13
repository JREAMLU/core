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

/*
func Insert(cron Cronlist) (uint64, error) {
	res := core.X.Create(&cron)
	if res.Error != nil {
		return 0, res.Error
	}
	return cron.ID, nil
}

func Update(cron Cronlist, id []uint64) error {
	res := core.X.Table("cronlist").Where("id IN (?)", id).Updates(cron)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func Delete(id []uint64) error {
	res := core.X.Delete(Cronlist{}, "id IN (?)", id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func Select(id []uint64) (cronlist Cronlist, err error) {
	sql := `
SELECT	 id, name, type
FROM  	  cronlist
WHERE 	id IN (?)
`
	res := core.X.Raw(sql, id).Scan(&cronlist)
	if res.Error != nil {
		return cronlist, res.Error
	}
	return cronlist, nil
}

//transaction
func Transact() error {
	tx := core.X.Begin()
	cronlist := Cronlist{
		Name: "Iversion",
		Type: 2,
	}
	res := core.X.Create(&cronlist)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}

	res = core.X.Delete(Cronlist{}, "id IN (?)", []uint64{361})
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	tx.Commit()
	return nil
}
*/
