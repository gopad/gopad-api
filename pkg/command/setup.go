package command

import (
	"io"
	"net/url"
	"os"
	"strings"

	"github.com/gopad/gopad-api/pkg/config"
	"github.com/gopad/gopad-api/pkg/store"
	"github.com/gopad/gopad-api/pkg/store/boltdb"
	"github.com/gopad/gopad-api/pkg/store/gormdb"
	"github.com/gopad/gopad-api/pkg/upload"
	"github.com/gopad/gopad-api/pkg/upload/file"
	"github.com/gopad/gopad-api/pkg/upload/s3"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"github.com/uber/jaeger-client-go"
	tracecfg "github.com/uber/jaeger-client-go/config"
)

func setupConfig(cfg *config.Config) error {
	if cfg.File != "" {
		viper.SetConfigFile(cfg.File)
	} else {
		viper.SetConfigName("api")

		viper.AddConfigPath("/etc/gopad")
		viper.AddConfigPath("$HOME/.gopad")
		viper.AddConfigPath("./config")
	}

	if err := viper.ReadInConfig(); err != nil {
		switch err.(type) {
		case viper.ConfigFileNotFoundError:
			log.Info().
				Msg("continue without config")
		case viper.UnsupportedConfigError:
			log.Fatal().
				Err(err).
				Msg("unsupported config type")

			return err
		default:
			log.Fatal().
				Err(err).
				Msg("failed to read config")

			return err
		}
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatal().
			Err(err).
			Msg("failed to parse config")

		return err
	}

	return nil
}

func setupLogger(cfg *config.Config) {
	switch strings.ToLower(cfg.Logs.Level) {
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "trace":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	if cfg.Logs.Pretty {
		log.Logger = log.Output(
			zerolog.ConsoleWriter{
				Out:     os.Stderr,
				NoColor: !cfg.Logs.Color,
			},
		)
	}
}

func setupTracing(cfg *config.Config) (io.Closer, error) {
	switch {
	case cfg.Tracing.Enabled:
		closer, err := tracecfg.Configuration{
			Sampler: &tracecfg.SamplerConfig{
				Type:  jaeger.SamplerTypeConst,
				Param: 1,
			},
			Reporter: &tracecfg.ReporterConfig{
				LocalAgentHostPort: cfg.Tracing.Endpoint,
			},
		}.InitGlobalTracer("gopad-api")

		if err != nil {
			return nil, err
		}

		log.Info().
			Str("addr", cfg.Tracing.Endpoint).
			Msg("application tracing is enabled")

		return closer, nil
	default:
		log.Info().
			Msg("application tracing is disabled")

		return nil, nil
	}
}

func setupUploads(cfg *config.Config) (upload.Upload, error) {
	parsed, err := url.Parse(cfg.Upload.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	switch parsed.Scheme {
	case "file":
		return file.New(cfg.Upload)
	case "s3":
		return s3.New(cfg.Upload)
	case "minio":
		return s3.New(cfg.Upload)
	}

	return nil, upload.ErrUnknownDriver
}

func setupStorage(cfg *config.Config) (store.Store, error) {
	parsed, err := url.Parse(cfg.Database.DSN)

	if err != nil {
		return nil, errors.Wrap(err, "failed to parse dsn")
	}

	switch parsed.Scheme {
	case "boltdb":
		return boltdb.New(cfg.Database)
	case "mysql", "mariadb":
		return gormdb.New(cfg.Database)
	case "postgres", "postgresql":
		return gormdb.New(cfg.Database)
	}

	return nil, store.ErrUnknownDriver
}
