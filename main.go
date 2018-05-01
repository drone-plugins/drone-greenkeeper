package main

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	version = "0.0.0"
	build   = "0"
)

func main() {
	app := cli.NewApp()
	app.Name = "npm plugin"
	app.Usage = "npm plugin"
	app.Action = run
	app.Version = fmt.Sprintf("%s+%s", version, build)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "folder",
			Usage:  "folder containing package.json",
			EnvVar: "PLUGIN_FOLDER",
		},

		// NPM options
		cli.StringFlag{
			Name:   "npm_username",
			Usage:  "NPM username",
			EnvVar: "PLUGIN_USERNAME,NPM_USERNAME",
		},
		cli.StringFlag{
			Name:   "npm_password",
			Usage:  "NPM password",
			EnvVar: "PLUGIN_PASSWORD,NPM_PASSWORD",
		},
		cli.StringFlag{
			Name:   "npm_email",
			Usage:  "NPM email",
			EnvVar: "PLUGIN_EMAIL,NPM_EMAIL",
		},
		cli.StringFlag{
			Name:   "npm_token",
			Usage:  "NPM deploy token",
			EnvVar: "PLUGIN_TOKEN,NPM_TOKEN",
		},
		cli.StringFlag{
			Name:   "npm_registry",
			Usage:  "NPM registry",
			Value:  GlobalRegistry,
			EnvVar: "PLUGIN_REGISTRY,NPM_REGISTRY",
		},
		cli.BoolFlag{
			Name:   "npm_skip_verify",
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
			Name:   "gk_token",
			Usage:  "Greenkeeper token",
			EnvVar: "PLUGIN_GK_TOKEN,GK_TOKEN",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			Folder: c.String("folder"),
			Update: c.Bool("update"),
			Upload: c.Bool("upload"),
		},
		Npm: Npm{
			Username:   c.String("username"),
			Password:   c.String("password"),
			Token:      c.String("token"),
			Email:      c.String("email"),
			Registry:   c.String("registry"),
			SkipVerify: c.Bool("skip_verify"),
		},
		Greenkeeper: Greenkeeper{
			Token: c.String("gk_token"),
		},
	}

	return plugin.Exec()
}
