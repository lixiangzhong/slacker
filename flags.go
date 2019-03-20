package slacker

import (
	"github.com/urfave/cli"
)

const (
	MVCDefaultDir = "mvc_default_dir"
)

var (
	DBName string
)

var (
	DBAddrFlag   = cli.StringFlag{Name: "addr", Value: "127.0.0.1:3306"}
	DBUserFlag   = cli.StringFlag{Name: "user", Value: "root"}
	DBPasswdFlag = cli.StringFlag{Name: "passwd", Value: "123456"}
	DBNameFlag   = cli.StringFlag{Name: "db", Value: "test", Destination: &DBName}
	DBTableFlag  = cli.StringFlag{Name: "table", Value: ""}

	//

	DirFlag = cli.StringFlag{Name: "dir", Value: MVCDefaultDir}
)
