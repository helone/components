package table

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/oleiade/reflections"
	"os"
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

	pk := new(tablestore.PrimaryKey)
	for field, tag := range tags {
		//name, index := strings.Split(tag, ",")

		fmt.Println(field)
		fmt.Println(tag)
	}
	os.Exit(1)
	pk.AddPrimaryKeyColumn("uuid", "123")
	return pk, nil
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
