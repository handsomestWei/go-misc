package nosql

import (
	"context"
	"github.com/olivere/elastic"
)

type esClient struct {
	es *elastic.Client
}

// es客户端连接
func NewElasticClient(conn string) *esClient {
	elasticClient, err := elastic.NewClient(elastic.SetURL(conn))
	if err != nil {
		panic(err)
		return nil
	}
	return &esClient{
		es: elasticClient,
	}
}

func (c *esClient) Gets(index, typ, id string) *elastic.GetResult {
	// 通过id查找
	get1, err := c.es.Get().Index(index).Type(typ).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
		return nil
	}
	return get1
}
