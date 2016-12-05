package mysql

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*
CREATE TABLE `redirect` (
    `redirect_id` INT(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT "短网址唯一id,自增长",
    `long_url` VARCHAR(255) NOT NULL DEFAULT "" COMMENT "原始url|JREAMLU|2016-10-10",
    `short_url` CHAR(25) NOT NULL DEFAULT "" COMMENT "短url|JREAMLU|2016-10-10",
    `long_crc` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT "原始url crc|JREAMLU|2016-10-10",
    `short_crc` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT "短url crc|JREAMLU|2016-10-10",
    `status` TINYINT(4) UNSIGNED NOT NULL DEFAULT 0 COMMENT "状态 0:删除 1:正常|JREAMLU|2016-10-10",
    `created_by_ip` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT "创建者ip|JREAMLU|2016-10-10",
    `updated_by_ip` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT "更新者ip|JREAMLU|2016-10-10",
    `created_at` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT "创建时间timestamp|JREAMLU|2016-10-10",
    `updated_at` INT(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT "更新时间timestamp|JREAMLU|2016-10-10",
    PRIMARY KEY (`redirect_id`),
    KEY `long_crc` (`long_crc`),
    KEY `short_url` (`short_url`)
) ENGINE = INNODB DEFAULT CHARSET=utf8 COMMENT="基建|短网址表|JREAMLU|2016-10-10";
*/

var (
	driver        = "mysql"
	setting       = "root:123@tcp(127.0.0.1:3306)/jream?charset:utf8&parsetime:true&loc:local"
	singulartable = true
	logmode       = false
)

type Redirect struct {
	ID          uint64 `gorm:"primary_key;column:redirect_id"`
	LongURL     string `gorm:"column:long_url"`
	ShortURL    string `gorm:"column:short_url"`
	LongCrc     uint64 `gorm:"column:long_crc"`
	ShortCrc    uint64 `gorm:"column:short_crc"`
	Status      uint8  `gorm:"column:status"`
	CreatedByIP uint64 `gorm:"column:created_by_ip"`
	UpdateByIP  uint64 `gorm:"column:updated_by_ip"`
	CreateAT    uint64 `gorm:"column:created_at"`
	UpdateAT    uint64 `gorm:"column:updated_at"`
}

func TestConGorm(t *testing.T) {
	Convey("func ConGorm()", t, func() {
		var gc GormConf
		gc.Driver = driver
		gc.Setting = setting
		gc.SingularTable = singulartable
		gc.LogMode = logmode
		err := gc.InitGorm()
		So(err, ShouldBeNil)

		// X.Close()
	})
}

func TestConGorms(t *testing.T) {
	Convey("func ConGorms()", t, func() {
		var gcs GormConfs
		err := gcs.InitGorms("./mysql.yml")
		So(err, ShouldBeNil)
	})
}

func TestInsert(t *testing.T) {
	var r Redirect
	r.LongURL = "http://o9d.cn/"
	r.ShortURL = "http://goo.lu/XI13G"
	r.LongCrc = 12345
	r.ShortCrc = 54321
	r.Status = 1
	r.CreatedByIP = 4294967295
	r.UpdateByIP = 4294967295
	r.CreateAT = 1473057106
	r.UpdateAT = 1473057106

	Convey("func Insert()", t, func() {
		id, err := insert(r)
		So(err, ShouldBeNil)
		So(id, ShouldBeGreaterThan, 0)
	})
}

func TestUpdate(t *testing.T) {
	var r Redirect
	r.LongURL = "http://www.o9d.cn/"
	r.ShortURL = "http://www.goo.lu/XI13G"
	r.LongCrc = 123456
	r.ShortCrc = 654321
	r.Status = 1
	r.CreatedByIP = 4294967295
	r.UpdateByIP = 4294967295
	r.CreateAT = 1473057106
	r.UpdateAT = 1473057106

	ids := []uint64{}
	var i uint64
	for i = 1; i <= 20000; i++ {
		ids = append(ids, i)
	}

	Convey("func Update()", t, func() {
		err := update(r, ids)
		So(err, ShouldBeNil)
	})
}

func BenchmarkInsert(b *testing.B) {
	var r Redirect
	r.LongURL = "http://o9d.cn/"
	r.ShortURL = "http://goo.lu/XI13G"
	r.LongCrc = 12345
	r.ShortCrc = 54321
	r.Status = 1
	r.CreatedByIP = 4294967295
	r.UpdateByIP = 4294967295
	r.CreateAT = 1473057106
	r.UpdateAT = 1473057106

	Convey("bench Insert()", b, func() {
		for i := 0; i < b.N; i++ {
			insert(r)
		}
	})
}

func BenchmarkUpdate(b *testing.B) {
	var r Redirect
	r.LongURL = "http://www.o9d.cn/"
	r.ShortURL = "http://www.goo.lu/XI13G"
	r.LongCrc = 123456
	r.ShortCrc = 654321
	r.Status = 1
	r.CreatedByIP = 4294967295
	r.UpdateByIP = 4294967295
	r.CreateAT = 1473057106
	r.UpdateAT = 1473057106

	ids := []uint64{}
	var i uint64
	for i = 1; i <= 200000; i++ {
		ids = append(ids, i)
	}

	Convey("bench Update()", b, func() {
		for i := 0; i < b.N; i++ {
			update(r, ids)
		}
	})
}

func insert(r Redirect) (uint64, error) {
	res := X.Create(&r)
	if res.Error != nil {
		return 0, res.Error
	}
	return r.ID, nil
}

func update(r Redirect, id []uint64) error {
	res := X.Table("redirect").Where("redirect_id IN (?)", id).Updates(r)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func delete(id []uint64) error {
	res := X.Delete(Redirect{}, "id IN (?)", id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func Select(id []uint64) (r Redirect, err error) {
	sql := `
SELECT  *
FROM    redirect
WHERE 	redirect_id IN (?)
`
	res := X.Raw(sql, id).Scan(&r)
	if res.Error != nil {
		return r, res.Error
	}
	return r, nil
}
