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

			var table Table
			var ok bool
			tablename := c.String("table")
			tables := Tables(mysqlconfig.DBName)
			for _, t := range tables {
				if t.Name == tablename {
					table = t
					ok = true
					break
				}
			}
			if !ok {
				return cli.NewExitError("table does not exist", 1)
			}

			var mvc = c.Args().First()
			var dir = c.String("dir")
			switch mvc {
			case "m", "model", "models":
				table.ExecTemplate("m", dir)
			case "v", "vue", "view":
				table.ExecTemplate("v", dir)
			case "js":
				table.ExecTemplate("js", dir)
			case "c", "controller", "controllers":
				table.ExecTemplate("c", dir)
			case "s", "sql":
				table.ExecTemplate("sql", dir)
				addInitCreateTableFile(table.Name)
				ReBindata(dir, "sql")
			case "a", "all":
				table.ExecTemplate("m", MVCDefaultDir)
				table.ExecTemplate("v", MVCDefaultDir)
				table.ExecTemplate("c", MVCDefaultDir)
				table.ExecTemplate("js", MVCDefaultDir)
				{
					table.ExecTemplate("sql", MVCDefaultDir)
					addInitCreateTableFile(table.Name)
					ReBindata(dir, "sql")
				}
			case "":
				return cli.NewExitError("missing action !!!\nexample:\n\tslacker add <action>", 0)
			default:
				return cli.NewExitError("action: "+mvc+" unknown", 0)
			}

			return nil
		},
	}
}
