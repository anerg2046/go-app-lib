package database

import (
	elastic6 "github.com/olivere/elastic/v6"
	"github.com/olivere/elastic/v7"
)

func ElasticConn(address ...string) *elastic.Client {
	esClient, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(address...))
	if err != nil {
		panic(err)
	}
	return esClient
}
func Elastic6Conn(address ...string) *elastic6.Client {
	esClient, err := elastic6.NewClient(elastic6.SetSniff(false), elastic6.SetURL(address...))
	if err != nil {
		panic(err)
	}
	return esClient
}
