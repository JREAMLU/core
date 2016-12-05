package mongo

import (
	"errors"
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

// MgoClient mongo client
type MgoClient struct {
	session *mgo.Session
	DBName  string
	URL     string
}

// NewMongoClient new mongo client
func NewMongoClient(url, dbName string) *MgoClient {
	return &MgoClient{
		DBName: dbName,
		URL:    url,
	}
}

// Session session
func (client *MgoClient) Session() (*mgo.Session, error) {
	var err error
	if client.session == nil {
		client.session, err = mgo.Dial(client.URL)
		if err != nil {
			return nil, err
		}
		if client.session == nil {
			return nil, errors.New("session is nil")
		}
	}
	return client.session.Clone(), nil
}

// CollectExists Returns true if the collection exists.
func (client *MgoClient) CollectExists(db *mgo.Database, collectName string) bool {
	c := db.C(`system.namespaces`)
	query := c.Find(map[string]string{`name`: fmt.Sprintf(`%s.%s`, client.DBName, collectName)})
	count, _ := query.Count()
	if count > 0 {
		return true
	}
	return false
}

//e.g
/*
main.go
//init mongodb
model.InitMongo(yaml.Yconf.Platform.Mongodb, yaml.Yconf.Platform.Mgo_db_name)
*/
/* =========================================*/
/*
model
package model

import (
	"fmt"
	"log"

	"git.corp.plu.cn/phpgo/core/mongo"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Jlog struct {
	Id        bson.ObjectId `bson:"_id"`
	Taskid    uint64        `bson:"taskid"`
	Pid       uint64        `bson:"pid"`
	Log       string        `bson:"log"`
	Shell     string        `bson:"shell"`
	ErrorLog  string        `bson:"errorlog"`
	CreatedAt int64         `bson:"create_at"`
}

var mongoClient *mongo.MongoClient

func InitMongo(url, dbName string) {
	mongoClient = mongo.NewMongoClient(url, dbName)
}

func GetLogs() []Jlog {
	session, err := mongoClient.Session()
	if err != nil {
		log.Println("mongodb err: ", err)
	}
	defer session.Close()
	c := session.DB(mongoClient.DBName).C("cron")

	var jlog []Jlog
	err = c.Find(bson.M{"log": "ppppppp"}).All(&jlog)
	if err != nil {
		log.Fatal(err)
	}

	return jlog
}

func GetLogsByObjectId() *Jlog {
	session, err := mongoClient.Session()
	if err != nil {
		log.Println("mongodb err: ", err)
	}
	defer session.Close()
	c := session.DB(mongoClient.DBName).C("cron")

	if err != nil {
		log.Println("mongodb err: ", err)
	}

	id := "574c1a1020733e03a077c771"
	objectId := bson.ObjectIdHex(id)
	jlog := new(Jlog)
	c.Find(bson.M{"_id": objectId}).One(&jlog)
	fmt.Println(jlog)
	return jlog
}

func AddLogs(j *Jlog) (id interface{}, err error) {
	session, err := mongoClient.Session()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(mongoClient.DBName).C("cron")

	if !mongoClient.CollectExists(session.DB(mongoClient.DBName), "cron") {
		cIndex := mgo.Index{Key: []string{"pid"}, Unique: false, Background: false}
		c.EnsureIndex(cIndex)
	}

	j.Id = bson.NewObjectId()
	err = c.Insert(j)
	if err != nil {
		return nil, err
	}
	return j.Id, nil
}

func InsertLogs() error {
	session, err := mongoClient.Session()
	if err != nil {
		log.Println("mongodb err: ", err)
	}
	defer session.Close()
	c := session.DB(mongoClient.DBName).C("cron")

	err = c.Insert(
		&Jlog{Id: bson.NewObjectId(), Pid: 17, Log: "ppppppp", Action: "do17"},
		&Jlog{Id: bson.NewObjectId(), Pid: 18, Log: "zzzzzzzz", Action: "do18"})

	if err != nil {
		log.Fatal(err)
		return err
	}

	log.Printf("Insert Success %v", bson.NewObjectId())
	return nil
}

func UpdateLogs() error {
	session, err := mongoClient.Session()
	if err != nil {
		log.Println("mongodb err: ", err)
	}
	defer session.Close()
	c := session.DB(mongoClient.DBName).C("cron")
	return c.Update(bson.M{"pid": 17}, bson.M{"$set": bson.M{"log": "LBJ", "action": "do17.1"}})
}

func DeteleLogs(id interface{}) error {
	session, err := mongoClient.Session()
	if err != nil {
		log.Println("mongodb err: ", err)
	}
	defer session.Close()
	c := session.DB(mongoClient.DBName).C("cron")
	return c.Remove(bson.M{"_id": id}) //bson.ObjectIdHex("574c1a1020733e03a077c772")
}
func AddLogsBatch(docs []interface{}) (id interface{}, err error) {
	session, err := mongoClient.Session()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(mongoClient.DBName).C("cron")

	if !mongoClient.CollectExists(session.DB(mongoClient.DBName), "cron") {
		cIndex := mgo.Index{Key: []string{"pid"}, Unique: false, Background: false}
		c.EnsureIndex(cIndex)
	}

	b := c.Bulk()
	b.Insert(docs...)
	// b.Insert(docs...)
	// b.Insert(docs...)
	res, err := b.Run()

	fmt.Println("res", res)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func AddLogsBat(docs []interface{}) (id interface{}, err error) {
	session, err := mongoClient.Session()
	if err != nil {
		return nil, err
	}
	defer session.Close()

	c := session.DB(mongoClient.DBName).C("cron")

	if !mongoClient.CollectExists(session.DB(mongoClient.DBName), "cron") {
		cIndex := mgo.Index{Key: []string{"pid"}, Unique: false, Background: false}
		c.EnsureIndex(cIndex)
	}

	err = c.Insert(docs...)

	if err != nil {
		return nil, err
	}
	return nil, nil
}
*/
