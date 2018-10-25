package slacker

import (
	"github.com/urfave/cli"
)

var (
	DBHostFlags   = cli.StringFlag{Name: "host", Value: "127.0.0.1:3306"}
	DBUserFlags   = cli.StringFlag{Name: "user", Value: "root"}
	DBPasswdFlags = cli.StringFlag{Name: "passwd", Value: "123456"}
	DBNameFlags   = cli.StringFlag{Name: "db", Value: "test"}
)
