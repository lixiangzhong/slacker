package slacker

import (
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/cosmtrek/air/runner"
	"github.com/urfave/cli/v2"
)

func Run() *cli.Command {
	return &cli.Command{
		Name:  "run",
		Usage: "live reload",
		Action: func(c *cli.Context) error {
			cfgfile := "slacker.run.toml"
			_, err := os.Lstat(cfgfile)
			if err != nil {
				if os.IsNotExist(err) {
					err = CreateDefaultRunConfig(cfgfile)
					if err != nil {
						return err
					}
				}
			}
			e, err := runner.NewEngine(cfgfile, true)
			if err != nil {
				return err
			}
			go e.Run()
			waitSignal()
			e.Stop()
			return nil
		},
	}
}

func waitSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	<-c
	signal.Reset()
}

func CreateDefaultRunConfig(filename string) error {
	s := `
# Config file for [Air](https://github.com/cosmtrek/air) in TOML format

# Working directory
# . or absolute path, please note that the directories following must be under root.
root = "." 
tmp_dir = "tmp"

[build]
# Just plain old shell command. You could use make as well.
cmd = "go build -o ./tmp/main ."
# Binary file yields from cmd.
bin = "tmp/main"
# Watch these filename extensions.
include_ext = ["go", "tpl", "tmpl", "html"]
# Ignore these filename extensions or directories.
exclude_dir = ["dist", "tmp", "vendor", "src", "node_modules", "markdowndoc"]
# Watch these directories if you specified.
include_dir = []
# Exclude files.
exclude_file = []
# It's not necessary to trigger build each time file changes if it's too frequent.
delay = 1000 # ms
# Stop to run old binary when build errors occur.
stop_on_error = true
# This log file places in your tmp_dir.
log = "air_errors.log"

[log]
# Show log time
time = false

[color]
# Customize each part's color. If no color found, use the raw app log.
main = "magenta"
watcher = "cyan"
build = "yellow"
runner = "green"

[misc]
# Delete tmp directory on exit
clean_on_exit = true
`
	return ioutil.WriteFile(filename, []byte(s), 0755)
}
