package table

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"reflect"
)

type Config struct {
	EndPoint        string
	InstanceName    string
	AccessKeyId     string
	AccessKeySecret string
}

type Client struct {
	table  *tablestore.TableStoreClient
	config *Config
}

var (
	tagToIndexTypeMap = map[string]tablestore.FieldType{
		"bool":   tablestore.FieldType_BOOLEAN,
		"float":  tablestore.FieldType_DOUBLE,
		"string": tablestore.FieldType_KEYWORD,
		"int":    tablestore.FieldType_LONG,
		"nested": tablestore.FieldType_NESTED,
		"text":   tablestore.FieldType_TEXT,
		"geo":    tablestore.FieldType_GEO_POINT,
	}

	kindToIndexTypeMap = map[reflect.Kind]tablestore.FieldType{
		reflect.Bool:    tablestore.FieldType_BOOLEAN,
		reflect.Float64: tablestore.FieldType_DOUBLE,
		reflect.String:  tablestore.FieldType_KEYWORD,
		reflect.Int64:   tablestore.FieldType_LONG,
	}
)
