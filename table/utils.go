package table

import (
	"errors"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/oleiade/reflections"
	"reflect"
	"strings"
)

func getTableName(table interface{}) string {
	v := reflect.Indirect(reflect.ValueOf(table))
	return strings.ToLower(v.Type().Name())
}

func primaryKey(table interface{}) (*tablestore.PrimaryKey, error) {
	tags, err := reflections.Tags(table, "json")

	if err != nil {
		return nil, err
	}
	tablePrimaryKey := new(tablestore.PrimaryKey)

	for field, tag := range tags {
		index, _ := reflections.GetFieldTag(table, field, "primaryKey")
		value, _ := reflections.GetField(table, field)
		if index != "" {
			tablePrimaryKey.AddPrimaryKeyColumn(tag, value)
		}
	}

	if len(tablePrimaryKey.PrimaryKeys) == 0 {
		return nil, errors.New("not found table index row")
	}

	return tablePrimaryKey, nil
}

func getFieldNameMap(item interface{}) (map[string]string, map[string]string, error) {
	tags, err := reflections.Tags(item, "json")
	if err != nil {
		return nil, nil, err
	}

	fieldToJSONMap := map[string]string{}
	jsonToFieldMap := map[string]string{}

	for field, tag := range tags {
		fieldToJSONMap[field] = tag
		jsonToFieldMap[tag] = field
	}

	return fieldToJSONMap, jsonToFieldMap, nil
}
