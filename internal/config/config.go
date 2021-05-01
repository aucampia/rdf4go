package config

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

const (
	// EnvPrefix is the environment prefix for the config.
	EnvPrefix   string = "gosvc"
	defaultHost string = "127.0.0.1"
	defaultPort int    = 22505
)

var (
	// EnvPrefixU is the upper case environment prefix for the config.
	EnvPrefixU string = strings.ToUpper(EnvPrefix)
)

// Config contains configuration for this service
type Config struct {
	// The host to listen on, e.g. "1.2.3.4".
	Host string
	// The port to listen on, e.g. 1234.
	Port int
}

// Address returns the Address of the config as "{{ .Host }}:{{ .Port }".
func (c *Config) Address() string {
	return fmt.Sprintf("%s:%v", c.Host, c.Port)
}

func (c *Config) setDefaults() {
	c.Host = defaultHost
	c.Port = defaultPort
}

// Load loads values into the config struct.
func (c *Config) Load(vi *viper.Viper) error {
	if vi == nil {
		vi = viper.New()
	}
	var err error
	defaultConfig := NewConfig()
	// This is to deal with some quirks of viper
	// see https://github.com/spf13/viper/issues/188#issuecomment-413368673
	defaultConfigYaml, err := yaml.Marshal(defaultConfig)
	if err != nil {
		return fmt.Errorf("failed to marshall dto: %v", err)
	}
	defaultmBytes := bytes.NewReader(defaultConfigYaml)
	vi.SetConfigType("yaml")
	if err := vi.MergeConfig(defaultmBytes); err != nil {
		return fmt.Errorf("failed to merge config: %v", err)
	}
	vi.AutomaticEnv()
	vi.SetEnvPrefix(EnvPrefix)

	err = vi.Unmarshal(c)
	if err != nil {
		return fmt.Errorf("failed to unmarshall dto: %v", err)
	}
	return nil
}

// NewConfig returns a config struct with default values.
func NewConfig() *Config {
	config := Config{}
	config.setDefaults()
	return &config
}

// LoadConfig loads the configuration and returns a config struct.
func LoadConfig(vi *viper.Viper, config *Config) (*Config, error) {
	if config == nil {
		config = NewConfig()
	}
	err := config.Load(nil)
	if err != nil {
		return nil, err
	}
	return config, nil
}
