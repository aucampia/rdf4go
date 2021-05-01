package config

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	envHost = fmt.Sprintf("%s_HOST", EnvPrefixU)
	envPort = fmt.Sprintf("%s_PORT", EnvPrefixU)
)

func TestDefaultConfig(t *testing.T) {
	config := NewConfig()
	assert.Equal(t, defaultHost, config.Host)
	assert.Equal(t, defaultPort, config.Port)
	assert.Equal(t, fmt.Sprintf("%s:%v", defaultHost, defaultPort), config.Address())
}

func TestLoadConfigFromEnv(t *testing.T) {
	useHost := "1.2.3.4"
	usePort := 1234

	restore := map[string]*string{}
	kvps := map[string]string{
		envHost: useHost,
		envPort: strconv.Itoa(usePort),
	}
	defer func() {
		for key, value := range restore {
			if value != nil {
				os.Setenv(key, *value)
			} else {
				os.Unsetenv(key)
			}
		}
	}()
	for key, value := range kvps {
		val, valSet := os.LookupEnv(key)
		if valSet {
			restore[key] = &val
		} else {
			restore[key] = nil
		}
		os.Setenv(key, value)
	}
	config, err := LoadConfig(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, useHost, config.Host)
	assert.Equal(t, usePort, config.Port)
}
