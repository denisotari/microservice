package config

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"go.uber.org/multierr"
)

const (
	envPrefix = "bootstrap" // TODO: change me
)

type Application struct {
	Env       string
	Addr      string
	Port      string
	Secret    string
	LogLevel  string
	LogFormat string
}

func (a *Application) IsProduction() bool {
	return a.Env == "production"
}

func (a *Application) validate() error {
	//if a.Addr == "" {
	//	return errors.New("empty address provided for an http server to start on")
	//}
	//if a.Secret == "" {
	//	return errors.New("empty secret provided")
	//}
	return nil
}

type Database struct {
	Host     string
	User     string
	Password string
	Port     int
	Db       string
}

func (d *Database) validate() error {
	if d.Host == "" {
		return errors.New("empty db host provided")
	}
	if d.Port == 0 {
		return errors.New("empty db port provided")
	}
	if d.User == "" {
		return errors.New("empty db user provided")
	}
	if d.Password == "" {
		return errors.New("empty db password provided")
	}
	if d.Db == "" {
		return errors.New("empty db name provided")
	}
	return nil
}

type Broker struct {
	UserURL         string
	UserCredits     string
	ExchangePrefix  string
	ExchangePostfix string
}

func (b *Broker) validate() error {
	if b.UserURL == "" {
		return errors.New("empty broker url provided")
	}
	if b.UserCredits == "" {
		return errors.New("empty broker credentials provided")
	}
	return nil
}

type Config struct {
	Application Application
	Database    Database
	Broker      Broker
}

func (c *Config) validate() error {
	return multierr.Combine(
		c.Application.validate(),
		c.Database.validate(),
		c.Broker.validate(),
	)
}

// Parse will parse the configuration from the environment variables and a file with the specified path.
// Environment variables have more priority than ones specified in the file.
func Parse(filepath string) (*Config, error) {
	setDefaults()

	// Parse the file
	viper.SetConfigFile(filepath)
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "failed to read the config file")
	}

	bindEnvVars() // remember to parse the environment variables

	// Unmarshal the config
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal the configuration")
	}

	// Validate the provided configuration
	if err := cfg.validate(); err != nil {
		return nil, errors.Wrap(err, "failed to validate the config")
	}
	return &cfg, nil
}

func (c *Config) Print() {
	if c.Application.IsProduction() {
		return
	}
	inspected := *c // get a copy of an actual object
	// Hide sensitive data
	inspected.Application.Secret = ""
	inspected.Database.User = ""
	inspected.Database.Password = ""
	inspected.Broker.UserCredits = ""
	fmt.Printf("%+v\n", inspected)
}

// TODO: set the default values here
func setDefaults() {
	viper.SetDefault("application.env", "production")
	viper.SetDefault("application.loglevel", "debug")
	viper.SetDefault("application.logformat", "text")
	viper.SetDefault("application.port", "8080")
}

func bindEnvVars() {
	viper.SetEnvPrefix(envPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()
}
