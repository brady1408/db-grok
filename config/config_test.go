package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	tt := []struct {
		name           string
		configFile     string
		makeAssertions func(*testing.T, string)
	}{
		{
			name:       "Good Path",
			configFile: "../config.json",
			makeAssertions: func(t *testing.T, c string) {
				cfg, err := LoadConfig(c)
				if err != nil {
					t.Errorf("Error pulling config: %s", err)
					return
				}
				assert.Equal(t, "postgres://postgres:mysecretpassword@tcp(192.168.23.29:3306)/atlascloud?sslmode=disable", cfg.ConnectionString)
				assert.Equal(t, "models", cfg.BaseImportPath)
				assert.Equal(t, 2, len(cfg.Models))
				assert.Equal(t, "TableA", cfg.Models[0].FromTableName)
				assert.Equal(t, "TableB", cfg.Models[1].FromTableName)
			},
		},
		{
			name:       "Bad Path",
			configFile: "badConfig.json",
			makeAssertions: func(t *testing.T, c string) {
				cfg, err := LoadConfig(c)
				assert.Nil(t, cfg)
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeAssertions(t, tc.configFile)
		})
	}
}
