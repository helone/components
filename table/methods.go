package table

import (
	"fmt"
	"github.com/aliyun/aliyun-tablestore-go-sdk/tablestore"
)

func (c *Client) Insert(item interface{}) error {
	return c.insert(item, tablestore.RowExistenceExpectation_EXPECT_NOT_EXIST)
}

func (c *Client) Update(item interface{}) error {
	return c.insert(item, tablestore.RowExistenceExpectation_EXPECT_EXIST)
}

func (c *Client) InsertOrUpdate(item interface{}) error {
	return c.insert(item, tablestore.RowExistenceExpectation_IGNORE)
}

func (c *Client) insert(item interface{}, cond tablestore.RowExistenceExpectation) error {
	rowReq := new(tablestore.PutRowRequest)
	change, err := getRowChange(item, cond)
	rowReq.PutRowChange = change
	_, err = c.table.PutRow(rowReq)

	if err != nil {
		return fmt.Errorf("insert error:%w", err)
	}
	return nil
}
