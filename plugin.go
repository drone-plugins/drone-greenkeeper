package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"os/exec"
	"os/user"
	"path"
	"strings"

	log "github.com/sirupsen/logrus"
)

type (
	// Config for the plugin.
	Config struct {
		Username    string
		Password    string
		Token       string
		Email       string
		Registry    string
		Folder      string
		SkipVerify  bool
		Update      bool
		Upload      bool
		GithubToken string
	}

	// Plugin values
	Plugin struct {
		Config Config
	}
)

// GlobalRegistry defines the default NPM registry.
const GlobalRegistry = "https://registry.npmjs.org"

// Exec executes the plugin.
func (p Plugin) Exec() error {
	if p.Config.Upload && p.Config.Update {
		return errors.New("Both update and upload are specified")
	} else if !(p.Config.Upload || p.Config.Update) {
		return errors.New("Neither update nor upload is specified")
	}

	// Print versions of commands
	err := showVersions(p.Config)

	if err != nil {
		return err
	}

	// Setup the NPM registry
	err = setupRegistry(p.Config)

	if err != nil {
		return err
	}

	// See if authentication is required
	if p.Config.Username == "" || p.Config.Token == "" {
		log.Info("NPM credentials are being used")

		// write npmrc for authentication
		err := writeNpmrc(p.Config)

		if err != nil {
			return err
		}

		// attempt to authenticate
		err = authenticate(p.Config)

		if err != nil {
			return err
		}
	} else {
		log.Info("Anonymous NPM credentials are being used")
	}

	var cmds []*exec.Cmd

	if p.Config.Update {
		cmds = append(cmds, updateCommand())
	} else {
		cmds = append(cmds, uploadCommand())
	}

	return runCommands(cmds, p.Config.Folder)
}

func showVersions(config Config) error {
	var cmds []*exec.Cmd

	// write the version command
	cmds = append(cmds, versionCommand())
	cmds = append(cmds, yarnVersionCommand())

	return runCommands(cmds, config.Folder)
}

func setupRegistry(config Config) error {
	var cmds []*exec.Cmd

	// write registry command
	if config.Registry != GlobalRegistry {
		cmds = append(cmds, registryCommand(config.Registry))
	}

	// write skip verify command
	if config.SkipVerify {
		cmds = append(cmds, skipVerifyCommand())
	}

	// run commands
	return runCommands(cmds, config.Folder)
}

/// writeNpmrc creates a .npmrc in the folder for authentication
func writeNpmrc(config Config) error {
	var npmrcContents string

	// check for an auth token
	if config.Token == "" {
		// check for a username
		if config.Username == "" {
			return errors.New("No username provided")
		}

		// check for an email
		if config.Email == "" {
			return errors.New("No email address provided")
		}

		// check for a password
		if config.Password == "" {
			log.Warning("No password provided")
		}

		log.WithFields(log.Fields{
			"username": config.Username,
			"email":    config.Email,
		}).Info("Specified credentials")

		npmrcContents = npmrcContentsUsernamePassword(config)
	} else {
		log.Info("Token credentials being used")

		npmrcContents = npmrcContentsToken(config)
	}

	// write npmrc file
	home := "/root"
	user, err := user.Current()
	if err == nil {
		home = user.HomeDir
	}
	npmrcPath := path.Join(home, ".npmrc")

	log.WithFields(log.Fields{
		"path": npmrcPath,
	}).Info("Writing npmrc")

	return ioutil.WriteFile(npmrcPath, []byte(npmrcContents), 0644)
}

/// authenticate atempts to authenticate with the NPM registry.
func authenticate(config Config) error {
	var cmds []*exec.Cmd

	cmds = append(cmds, alwaysAuthCommand())
	cmds = append(cmds, whoamiCommand())

	return runCommands(cmds, config.Folder)
}

// npmrcContentsUsernamePassword creates the contents from a username and
// password
func npmrcContentsUsernamePassword(config Config) string {
	// get the base64 encoded string
	authString := fmt.Sprintf("%s:%s", config.Username, config.Password)
	encoded := base64.StdEncoding.EncodeToString([]byte(authString))

	// create the file contents
	return fmt.Sprintf("_auth = %s\nemail = %s", encoded, config.Email)
}

/// Writes npmrc contents when using a token
func npmrcContentsToken(config Config) string {
	registry, _ := url.Parse(config.Registry)
	return fmt.Sprintf("//%s/:_authToken=%s", registry.Host, config.Token)
}

// versionCommand gets the npm version
func versionCommand() *exec.Cmd {
	return exec.Command("npm", "--version")
}

// yarnVersionCommand gets the yarn version
func yarnVersionCommand() *exec.Cmd {
	return exec.Command("yarn", "--version")
}

// registryCommand sets the NPM registry.
func registryCommand(registry string) *exec.Cmd {
	return exec.Command("npm", "config", "set", "registry", registry)
}

// alwaysAuthCommand forces authentication.
func alwaysAuthCommand() *exec.Cmd {
	return exec.Command("npm", "config", "set", "always-auth", "true")
}

// skipVerifyCommand disables ssl verification.
func skipVerifyCommand() *exec.Cmd {
	return exec.Command("npm", "config", "set", "strict-ssl", "false")
}

// whoamiCommand creates a command that gets the currently logged in user.
func whoamiCommand() *exec.Cmd {
	return exec.Command("npm", "whoami")
}

func updateCommand() *exec.Cmd {
	return exec.Command("greenkeeper-lockfile-update")
}

func uploadCommand() *exec.Cmd {
	return exec.Command("greenkeeper-lockfile-upload")
}

// trace writes each command to standard error (preceded by a ‘$ ’) before it
// is executed. Used for debugging your build.
func trace(cmd *exec.Cmd) {
	fmt.Fprintf(os.Stdout, "+ %s\n", strings.Join(cmd.Args, " "))
}

// runCommands executes the list of cmds in the given directory.
func runCommands(cmds []*exec.Cmd, dir string) error {
	for _, cmd := range cmds {
		err := runCommand(cmd, dir)

		if err != nil {
			return err
		}
	}

	return nil
}

func runCommand(cmd *exec.Cmd, dir string) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = dir
	trace(cmd)

	return cmd.Run()
}
