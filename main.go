package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var build = "0" // build number set at compile-time

func main() {
	app := cli.NewApp()
	app.Name = "npm plugin"
	app.Usage = "npm plugin"
	app.Action = run
	app.Version = fmt.Sprintf("1.0.0+%s", build)
	app.Flags = []cli.Flag{
		// NPM options
		cli.StringFlag{
			Name:   "username",
			Usage:  "NPM username",
			EnvVar: "PLUGIN_USERNAME,NPM_USERNAME",
		},
		cli.StringFlag{
			Name:   "password",
			Usage:  "NPM password",
			EnvVar: "PLUGIN_PASSWORD,NPM_PASSWORD",
		},
		cli.StringFlag{
			Name:   "email",
			Usage:  "NPM email",
			EnvVar: "PLUGIN_EMAIL,NPM_EMAIL",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "NPM deploy token",
			EnvVar: "PLUGIN_TOKEN,NPM_TOKEN",
		},
		cli.StringFlag{
			Name:   "registry",
			Usage:  "NPM registry",
			Value:  GlobalRegistry,
			EnvVar: "PLUGIN_REGISTRY,NPM_REGISTRY",
		},
		cli.StringFlag{
			Name:   "folder",
			Usage:  "folder containing package.json",
			EnvVar: "PLUGIN_FOLDER",
		},
		cli.BoolFlag{
			Name:   "skip_verify",
			Usage:  "skip SSL verification",
			EnvVar: "PLUGIN_SKIP_VERIFY",
		},

		// Greenkeeper options
		cli.BoolFlag{
			Name:   "update",
			Usage:  "update lockfile",
			EnvVar: "PLUGIN_UPDATE",
		},
		cli.BoolFlag{
			Name:   "upload",
			Usage:  "upload lockfile",
			EnvVar: "PLUGIN_UPLOAD",
		},
		cli.BoolFlag{
			Name:   "github_token",
			Usage:  "Github token",
			EnvVar: "PLUGIN_GITHUB_TOKEN,GREENKEEPER_GITHUB_TOKEN",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			Username:    c.String("username"),
			Password:    c.String("password"),
			Token:       c.String("token"),
			Email:       c.String("email"),
			Registry:    c.String("registry"),
			SkipVerify:  c.Bool("skip_verify"),
			Folder:      c.String("folder"),
			Update:      c.Bool("update"),
			Upload:      c.Bool("upload"),
			GithubToken: c.String("github_token"),
		},
	}

	return plugin.Exec()
}
