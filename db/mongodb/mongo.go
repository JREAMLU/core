package mongodb

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"
)

var mbDB *mgo.Database
var ErrCollectionDoesNotExist = errors.New(`Collection does not exist.`)
var dbName string

//getCollection
func GetCollection(collectName string) (*mgo.Collection, error) {
	// if err := initDB(); err != nil {
	// 	return nil, err
	// }
	return mbDB.C(collectName), nil
}

//initDB
func InitMongodb(dbName string, mongodb string) error {
	dbName = dbName
	if mbDB == nil {
		// mgo.SetDebug(true)
		session, err := mgo.Dial(mongodb)
		if err != nil {
			return err
		}
		//TODO 创建索引
		mbDB = session.DB(dbName)
	}
	return nil
}

// CollectExists Returns true if the collection exists.
func CollectExists(collectName string) bool {
	c, err := GetCollection(`system.namespaces`)
	if err != nil {
		return false
	}

	query := c.Find(map[string]string{`name`: fmt.Sprintf(`%s.%s`, dbName, collectName)})
	count, _ := query.Count()
	if count > 0 {
		return true
	}
	return false
}
