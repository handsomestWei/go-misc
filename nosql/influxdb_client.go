package nosql

import (
	"github.com/influxdata/influxdb/client/v2"
)

func NewInfluxdbClient(url, username, password string) (client.Client, error) {
	// 创建http客户端
	return client.NewHTTPClient(client.HTTPConfig{
		InsecureSkipVerify: true,
		Addr:               url,
		Username:           username,
		Password:           password,
	})
}
