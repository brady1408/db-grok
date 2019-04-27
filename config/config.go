package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/joomcode/errorx"

	"github.com/brady1408/db-grok/models"
)

// LoadConfig returns a models.config
// pass LoadConfig the config from the command parameters,
// this should be the path to a config.json
func LoadConfig(c string) (*models.BuildConfig, error) {
	Config := &models.BuildConfig{}

	b, err := ioutil.ReadFile(c)
	if err != nil {
		err = errorx.Decorate(err, "Error reading the config json: ")
		return nil, err
	}

	err = json.Unmarshal(b, Config)
	if err != nil {
		err = errorx.Decorate(err, "Error unmarshaling config: ")
		return nil, err
	}

	if len(Config.ConnectionString) == 0 {
		err = errorx.Decorate(err, "Invalid connection string: ")
		return nil, err
	}

	return Config, nil
}
