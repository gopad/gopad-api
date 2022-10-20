package command

import (
	"os"

	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/version"
	"github.com/urfave/cli/v2"
)

const (
	defaultMetricsAddr   = "0.0.0.0:8000"
	defaultServerAddress = "0.0.0.0:8080"
)

// Run parses the command line arguments and executes the program.
func Run() error {
	cfg := config.Load()

	app := &cli.App{
		Name:     "gopad-api",
		Version:  version.String,
		Usage:    "Etherpad for markdown with go",
		Authors:  authorList(cfg),
		Flags:    globalFlags(cfg),
		Before:   globalBefore(cfg),
		Commands: globalCommands(cfg),
	}

	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "Show the help",
	}

	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "Print the version",
	}

	return app.Run(os.Args)
}

func authorList(_ *config.Config) []*cli.Author {
	return []*cli.Author{
		{
			Name:  "Thomas Boerger",
			Email: "thomas@webhippie.de",
		},
	}
}

func globalFlags(cfg *config.Config) []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    "config-file",
			Value:   "",
			Usage:   "Path to config file",
			EnvVars: []string{"GOPAD_API_CONFIG_FILE"},
		},
		&cli.StringFlag{
			Name:        "log-level",
			Value:       "info",
			Usage:       "Set logging level",
			EnvVars:     []string{"GOPAD_API_LOG_LEVEL"},
			Destination: &cfg.Logs.Level,
		},
		&cli.BoolFlag{
			Name:        "log-pretty",
			Value:       true,
			Usage:       "Enable pretty logging",
			EnvVars:     []string{"GOPAD_API_LOG_PRETTY"},
			Destination: &cfg.Logs.Pretty,
		},
		&cli.BoolFlag{
			Name:        "log-color",
			Value:       true,
			Usage:       "Enable colored logging",
			EnvVars:     []string{"GOPAD_API_LOG_COLOR"},
			Destination: &cfg.Logs.Color,
		},
	}
}

func globalBefore(cfg *config.Config) cli.BeforeFunc {
	return func(c *cli.Context) error {
		setupLogger(cfg)
		return setupConfig(cfg)
	}
}

func globalCommands(cfg *config.Config) []*cli.Command {
	return []*cli.Command{
		Gen(cfg),
		Server(cfg),
		Health(cfg),
	}
}
