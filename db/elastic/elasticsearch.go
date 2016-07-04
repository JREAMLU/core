package elastic

import (
	"strconv"

	"github.com/astaxie/beego"

	"gopkg.in/olivere/elastic.v3"
)

var ESClient *elastic.Client

func InitElastic(url string) error {
	var err error
	ESClient, err = elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return err
	}

	info, code, err := ESClient.Ping(url).Do()
	if err != nil {
		return err
	}
	beego.Trace("Elasticsearch returned with code " + strconv.Itoa(code) + "and version " + info.Version.Number)

	esversion, err := ESClient.ElasticsearchVersion(url)
	if err != nil {
		return err
	}
	beego.Trace("Elasticsearch version " + esversion)

	return nil
}

func CreateIndexElastic(index string) error {
	exists, err := ESClient.IndexExists(index).Do()
	if err != nil {
		return err
	}
	if !exists {
		createIndex, err := ESClient.CreateIndex(index).Do()
		if err != nil {
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			beego.Trace("Not acknowledged: " + strconv.FormatBool(createIndex.Acknowledged))
		}
	}
	return nil
}
