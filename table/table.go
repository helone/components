package table

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func New(config *Config) *Client {
	return &Client{
		table:  tablestore.NewClient(config.EndPoint, config.InstanceName, config.AccessKeyId, config.AccessKeySecret),
		config: config,
	}
}
