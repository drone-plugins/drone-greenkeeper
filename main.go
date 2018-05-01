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
		cli.BoolFlag{
			Name:   "upload",
			Usage:  "upload lockfile",
			EnvVar: "PLUGIN_UPLOAD",
		},
		cli.BoolFlag{
			Name:   "update",
			Usage:  "update lockfile",
			EnvVar: "PLUGIN_UPDATE",
		},
		cli.StringFlag{
			Name:   "folder",
			Usage:  "folder containing package.json",
			EnvVar: "PLUGIN_FOLDER",
		},

		// Build options
		cli.StringFlag{
			Name:   "repo",
			Usage:  "Repo slug",
			EnvVar: "DRONE_REPO",
		},
		cli.StringFlag{
			Name:   "remote",
			Usage:  "Remote url",
			EnvVar: "DRONE_REMOTE_URL",
		},
		cli.StringFlag{
			Name:   "event",
			Usage:  "Build event",
			EnvVar: "DRONE_BUILD_EVENT",
		},
		cli.StringFlag{
			Name:   "branch",
			Usage:  "Build branch",
			EnvVar: "DRONE_COMMIT_BRANCH",
		},
		cli.StringFlag{
			Name:   "message",
			Usage:  "Commit message",
			EnvVar: "DRONE_COMMIT_MESSAGE",
		},
		cli.StringFlag{
			Name:   "job",
			Usage:  "Job number",
			EnvVar: "DRONE_JOB_NUMBER",
		},

		// NPM options
		cli.StringFlag{
			Name:   "npm_registry",
			Usage:  "NPM registry",
			Value:  GlobalRegistry,
			EnvVar: "PLUGIN_REGISTRY,NPM_REGISTRY",
		},
		cli.StringFlag{
			Name:   "npm_email",
			Usage:  "NPM email",
			EnvVar: "PLUGIN_EMAIL,NPM_EMAIL",
		},
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
			Name:   "npm_token",
			Usage:  "NPM deploy token",
			EnvVar: "PLUGIN_TOKEN,NPM_TOKEN",
		},
		cli.BoolFlag{
			Name:   "npm_skip_verify",
			Usage:  "skip SSL verification",
			EnvVar: "PLUGIN_SKIP_VERIFY",
		},

		// Greenkeeper options
		cli.StringFlag{
			Name:   "gk_token",
			Usage:  "Greenkeeper token",
			EnvVar: "PLUGIN_GK_TOKEN,GK_TOKEN",
		},
		cli.StringFlag{
			Name:   "gk_name",
			Usage:  "Greenkeeper lock commit name",
			EnvVar: "PLUGIN_GK_NAME",
		},
		cli.StringFlag{
			Name:   "gk_email",
			Usage:  "Greenkeeper lock commit email",
			EnvVar: "PLUGIN_GK_EMAIL",
		},
		cli.BoolFlag{
			Name:   "gk_ammend",
			Usage:  "Greenkeeper lockfile commit should ammend",
			EnvVar: "PLUGIN_GK_AMMEND",
		},
		cli.StringFlag{
			Name:   "gk_yarn_opts",
			Usage:  "Greenkeeper yarn lock options",
			EnvVar: "PLUGIN_GK_YARN_OPTS",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := Plugin{
		Config: Config{
			Update: c.Bool("update"),
			Upload: c.Bool("upload"),
			Folder: c.String("folder"),
		},
		Build: Build{
			Repo:    c.String("repo"),
			Remote:  c.String("remote"),
			Event:   c.String("event"),
			Branch:  c.String("branch"),
			Message: c.String("message"),
			Job:     c.String("job"),
		},
		Npm: Npm{
			Registry:   c.String("npm_registry"),
			Username:   c.String("npm_username"),
			Email:      c.String("npm_email"),
			Password:   c.String("npm_password"),
			Token:      c.String("npm_token"),
			SkipVerify: c.Bool("npm_skip_verify"),
		},
		Greenkeeper: Greenkeeper{
			Token:    c.String("gk_token"),
			Name:     c.String("gk_name"),
			Email:    c.String("gk_email"),
			Ammend:   c.Bool("gk_ammend"),
			YarnOpts: c.String("gk_yarn_opts"),
		},
	}

	return plugin.Exec()
}
