package main

import (
	"os"

	version "github.com/drone-stack/drone-release-version"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type formatter struct{}

func (*formatter) Format(entry *logrus.Entry) ([]byte, error) {
	return []byte(entry.Message), nil
}

func init() {
	// logrus.SetFormatter(&logrus.TextFormatter{
	// 	DisableTimestamp: true,
	// 	DisableColors:    true,
	// })
	logrus.SetFormatter(new(formatter))
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	// Load env-file if it exists first
	if env := os.Getenv("PLUGIN_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	app := cli.NewApp()
	app.Name = "release version plugin"
	app.Usage = "release version plugin"
	app.Action = run
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "name",
			Usage:  "plugin name",
			EnvVar: "PLUGIN_NAME",
		},
		cli.StringFlag{
			Name:   "release",
			Usage:  "plugin release",
			EnvVar: "PLUGIN_RELEASE",
		},
		cli.StringFlag{
			Name:   "url",
			Usage:  "plugin url",
			EnvVar: "PLUGIN_URL",
		},
		cli.StringFlag{
			Name:   "token",
			Usage:  "plugin token",
			EnvVar: "PLUGIN_TOKEN",
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}

func run(c *cli.Context) error {
	plugin := version.Plugin{
		URL:     c.String("url"),
		Token:   c.String("token"),
		Name:    c.String("name"),
		Release: c.String("release"),
		Type:    c.String("type"),
	}

	if err := plugin.Exec(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
	return nil
}
