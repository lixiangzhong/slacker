package slacker

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/lixiangzhong/slacker/bindata"
	"github.com/urfave/cli"
	"os"
	"time"
)

func New() cli.Command {
	return cli.Command{
		Name:  "new",
		Usage: "create new project",
		Flags: []cli.Flag{
			DBHostFlags,
			DBUserFlags,
			DBPasswdFlags,
			DBNameFlags,
		},
		Action: func(c *cli.Context) error {
			projectname := c.Args().First()
			if projectname == "" {
				return cli.NewExitError("must project name", 0)
			}
			if InGOPATH() == false {
				return cli.NewExitError("Not in $GOPATH", 1)
			}
			if err := Mkdir(projectname); err != nil {
				return cli.NewExitError(err, 1)
			}
			if err := os.Chdir(projectname); err != nil {
				return cli.NewExitError(err, 1)
			}
			var mysqlconfig = MysqlConfig(c)
			var tpldata TemplateData
			tpldata.ProjectName = projectname
			tpldata.ProjectPath = ProjectPath()
			tpldata.MysqlConfig = mysqlconfig
			if err := bindata.Restore(tpldata); err != nil {
				return cli.NewExitError(err, 1)
			}
			fmt.Printf("\nCreate Project %v Success !!!\n\n", projectname)
			fmt.Println("Install npm dependency packages:")
			fmt.Println("\tcd", projectname)
			fmt.Println("\tnpm install")
			fmt.Println("How to run:")
			fmt.Println("\tnpm run dev")
			fmt.Printf("\tgo run %v.go\n", projectname)
			fmt.Println("How to build")
			fmt.Println("\tnpm run build")
			fmt.Println("\tmake build")

			return nil
		},
	}

}

func MysqlConfig(c *cli.Context) *mysql.Config {
	var cfg = mysql.NewConfig()
	cfg.Addr = c.String("host")
	cfg.User = c.String("user")
	cfg.Passwd = c.String("passwd")
	cfg.DBName = c.String("db")
	cfg.Net = "tcp"
	cfg.Loc = time.Local
	// cfg.ParseTime = true
	return cfg
}
