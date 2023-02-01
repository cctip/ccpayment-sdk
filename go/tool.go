package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"sync"
	"time"
)

func BillId() string {
	return fmt.Sprintf("%s%d",
		Format(time.Now(), CompactLayout),
		GlobalSnowflakeNode().Generate().Int64())
}

func InitGlobalSnowflakeNode(nodeId int64) (err error) {
	globalSnowflakeNodeSync.Do(func() {
		globalSnowflakeNode, err = snowflake.NewNode(nodeId)
	})

	return err
}

func GlobalSnowflakeNode() *snowflake.Node {
	if globalSnowflakeNode == nil {
		if err := InitGlobalSnowflakeNode(1); err != nil {
			panic(err)
		}
	}

	return globalSnowflakeNode
}

func Format(t time.Time, layout ...string) string {
	if len(layout) != 0 {
		return t.Format(layout[0])
	}
	return t.Format(DefaultLayout)
}

var (
	globalSnowflakeNode     *snowflake.Node
	globalSnowflakeNodeSync sync.Once
)

const CompactLayout = "20060102150405"
const DefaultLayout = "2006-01-02 15:04:05"
