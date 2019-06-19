package main

import (
	"os"

	"github.com/lixiangzhong/slacker"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "一键生成后台管理项目"
	app.Version = "2.0.0 dao版本"
	app.Commands = []cli.Command{
		slacker.New(),
		slacker.Add(),
	}
	app.Run(os.Args)
}
