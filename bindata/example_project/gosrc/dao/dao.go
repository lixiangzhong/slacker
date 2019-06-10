package dao

import (
	"github.com/jmoiron/sqlx"
)

type Dao struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Dao {
	return &Dao{
		db: db,
	}
}
