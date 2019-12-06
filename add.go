package slacker

import (
	"github.com/urfave/cli"
)

func Add() cli.Command {
	return cli.Command{
		Name:  "add",
		Usage: "add table",
		Flags: []cli.Flag{
			DBAddrFlag,
			DBUserFlag,
			DBPasswdFlag,
			DBNameFlag,
			DBTableFlag,
			DirFlag,
		},
		Action: func(c *cli.Context) error {
			mysqlconfig := MysqlConfig(c)
			if err := ConnectDB(mysqlconfig); err != nil {
				return cli.NewExitError(err, 1)
			}

			tablename := c.String("table")
			var tables = make([]Table, 0)

			if tablename == "" {
				tables = Tables(GoModName(), mysqlconfig.DBName)
			} else {
				for _, t := range Tables(GoModName(), mysqlconfig.DBName) {
					if t.Name == tablename {
						tables = append(tables, t)
						break
					}
				}
			}

			if len(tables) == 0 {
				return cli.NewExitError("table does not exist", 1)
			}

			var mvc = c.Args().First()
			var dir = c.String("dir")

			for _, table := range tables {
				switch mvc {
				case "m", "model", "models":
					table.ExecTemplate("m", dir)
				case "v", "vue", "view":
					table.ExecTemplate("v", dir)
				case "comp", "component":
					table.ExecTemplate("component", dir)
				case "js":
					table.ExecTemplate("js", dir)
				case "c", "controller", "controllers":
					table.ExecTemplate("c", dir)
				case "dao", "d":
					table.ExecTemplate("dao", dir)
				case "service", "s":
					table.ExecTemplate("service", dir)
				case "markdown", "md":
					table.ExecTemplate("md", dir)
				case "go":
					table.ExecTemplate("m", dir)
					table.ExecTemplate("dao", dir)
					table.ExecTemplate("service", dir)
					table.ExecTemplate("c", dir)
				case "a", "all":
					table.ExecTemplate("m", dir)
					table.ExecTemplate("v", dir)
					table.ExecTemplate("component", dir)
					table.ExecTemplate("c", dir)
					table.ExecTemplate("js", dir)
					table.ExecTemplate("dao", dir)
					table.ExecTemplate("md", dir)
					table.ExecTemplate("service", dir)
				case "":
					return cli.NewExitError("missing action !!!\nexample:\n\tslacker add <action>", 0)
				default:
					return cli.NewExitError("action: "+mvc+" unknown", 0)
				}
			}
			return nil
		},
	}
}
