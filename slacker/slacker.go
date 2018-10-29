package main

import (
	"github.com/lixiangzhong/slacker"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Usage = "一键生成后台管理项目"
	app.Version = "0.0.2"
	app.Commands = []cli.Command{
		slacker.New(),
		slacker.Add(),
	}
	app.Run(os.Args)
}
