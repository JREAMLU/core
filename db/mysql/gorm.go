package mysql

import (
	"errors"

	"github.com/JREAMLU/core/com"
	"github.com/JREAMLU/core/io"
	// mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	yaml "gopkg.in/yaml.v2"
)

// GormConf gorm config
type GormConf struct {
	Name          string `yaml:"name"`
	Driver        string `yaml:"driver"`
	Setting       string `yaml:"setting"`
	SingularTable bool   `yaml:"singularTable"`
	LogMode       bool   `yaml:"logMode"`
}

// GormConfs gorm configs
type GormConfs struct {
	GormConf []GormConf `yaml:"mysqlConfig"`
}

// X gorm.DB
var X *gorm.DB

// XS gorm.DB map
var XS map[string]*gorm.DB

// InitGorm "root:root@tcp(localhost:3306)/plucron?charset=utf8&parseTime=True&loc=Local"
func (gconf *GormConf) InitGorm() error {
	var err error
	X, err = gorm.Open(gconf.Driver, gconf.Setting)
	if err != nil {
		return err
	}
	X.SingularTable(gconf.SingularTable)
	X.LogMode(gconf.LogMode)
	return nil
}

// InitGorms init gorms
func (gconfs *GormConfs) InitGorms(filePath string) error {
	XS = make(map[string]*gorm.DB)
	result, err := io.ReadAllBytes(filePath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(result, gconfs); err != nil {
		return err
	}
	for _, gconf := range gconfs.GormConf {
		XS[gconf.Name], err = gorm.Open(gconf.Driver, gconf.Setting)
		if err != nil {
			return errors.New(com.StringJoin("DB Alias \"", gconf.Name, "\": ", err.Error()))
		}
		XS[gconf.Name].SingularTable(gconf.SingularTable)
		XS[gconf.Name].LogMode(gconf.LogMode)
	}

	return nil
}

// GetXS get XS
func GetXS(serverName string) (*gorm.DB, error) {
	if _, ok := XS[serverName]; ok {
		return XS[serverName], nil
	}
	return nil, errors.New("map key <" + serverName + "> is nil")
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
