package main

import (
	"os"

	"github.com/lixiangzhong/slacker"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Usage = "一键生成后台管理项目"
	app.Version = "1.1.0"
	app.Commands = []cli.Command{
		slacker.New(),
		slacker.Add(),
	}
	app.Run(os.Args)
}
