package mysql

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

/*
CREATE TABLE `redirect` (
  `redirect_id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '短网址唯一id,自增长',
  `long_url` varchar(255) NOT NULL DEFAULT '' COMMENT '原始url',
  `short_url` char(25) NOT NULL DEFAULT '' COMMENT '短url',
  `long_crc` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '原始url crc',
  `short_crc` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '短url crc',
  `status` tinyint(4) unsigned NOT NULL DEFAULT '0' COMMENT '状态 0:删除 1:正常',
  `created_by_ip` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ip',
  `updated_by_ip` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新者ip',
  `created_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间timestamp',
  `updated_at` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间timestamp',
  PRIMARY KEY (`redirect_id`),
  KEY `long_crc` (`long_crc`),
  KEY `short_url` (`short_url`)
) ENGINE=InnoDB AUTO_INCREMENT=4208 DEFAULT CHARSET=utf8 COMMENT='短网址表'
*/

var (
	driver        = "mysql"
	setting       = "root:123@tcp(127.0.0.1:3306)/jream?charset:utf8&parsetime:true&loc:local"
	singulartable = true
	logmode       = false
)

type Redirect struct {
	Id          uint64 `gorm:"primary_key;column:redirect_id"`
	LongUrl     string `gorm:"column:long_url"`
	ShortUrl    string `gorm:"column:short_url"`
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

func TestInsert(t *testing.T) {
	var r Redirect
	r.LongUrl = "http://o9d.cn/"
	r.ShortUrl = "http://goo.lu/XI13G"
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
	r.LongUrl = "http://www.o9d.cn/"
	r.ShortUrl = "http://www.goo.lu/XI13G"
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
	r.LongUrl = "http://o9d.cn/"
	r.ShortUrl = "http://goo.lu/XI13G"
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
	r.LongUrl = "http://www.o9d.cn/"
	r.ShortUrl = "http://www.goo.lu/XI13G"
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
	return r.Id, nil
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
