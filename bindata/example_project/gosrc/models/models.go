package models

import "sync"

type Table interface {
	TableName() string
}

var (
	Tables = new(sync.Map)
)

func AutoMigrate(v Table) {
	Tables.Store(v.TableName(), v)
}
