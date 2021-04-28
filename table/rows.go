package table

import (
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
	"github.com/oleiade/reflections"
)

func getRowChange(item interface{}, cond tablestore.RowExistenceExpectation) (*tablestore.PutRowChange, error) {
	fields, err := reflections.Fields(item)
	if err != nil {
		return nil, err
	}

	fieldToJSONMap, _, err := getFieldNameMap(item)
	if err != nil {
		return nil, err
	}

	putRowChange := new(tablestore.PutRowChange)
	primaryKeyColumn, err := primaryKey(item)
	if err != nil {
		return nil, err
	}
	putRowChange.PrimaryKey = primaryKeyColumn

	for _, field := range fields {
		value, _ := reflections.GetField(item, field)
		column := fieldToJSONMap[field]
		putRowChange.AddColumn(column, value)
	}

	putRowChange.TableName = getTableName(item)
	putRowChange.SetCondition(cond)

	return putRowChange, nil
}
