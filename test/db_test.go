package test

import (
	"os"
	"testing"

	"github.com/brady1408/db-grok/db"
	"github.com/stretchr/testify/assert"
)

func TestGetTestDBInfo(t *testing.T) {
	// Create struct to hold previous env variables in case you have custom variables set.
	prevEnv := GetTestDBInfo()
	// Defer will handle pushing original env variables back so you can continue as normal.
	defer func() {
		var err error
		err = os.Setenv("DBGROK_DB_HOST", prevEnv.Host)
		err = os.Setenv("DBGROK_DB_NAME", prevEnv.Name)
		err = os.Setenv("DBGROK_DB_USER", prevEnv.User)
		err = os.Setenv("DBGROK_DB_PASS", prevEnv.Pass)
		err = os.Setenv("DBGROK_DB_PORT", prevEnv.Port)
		if err != nil {
			t.Error("Error setting previous env variables")
		}
	}()

	tt := []struct {
		name           string
		makeAssertions func(*testing.T)
	}{
		{
			name: "Get Custom",
			makeAssertions: func(t *testing.T) {
				// setting custom env variables in case the users already has variables set
				var err error
				err = os.Setenv("DBGROK_DB_HOST", "testHost")
				err = os.Setenv("DBGROK_DB_NAME", "testName")
				err = os.Setenv("DBGROK_DB_USER", "testUser")
				err = os.Setenv("DBGROK_DB_PASS", "testPass")
				err = os.Setenv("DBGROK_DB_PORT", "1234")
				if err != nil {
					t.Error("Error setting custom env variables")
				}
				currEnv := GetTestDBInfo()
				assert.Equal(t, "testHost", currEnv.Host)
				assert.Equal(t, "testName", currEnv.Name)
				assert.Equal(t, "testUser", currEnv.User)
				assert.Equal(t, "testPass", currEnv.Pass)
				assert.Equal(t, "1234", currEnv.Port)
			},
		},
		{
			name: "Get Default",
			makeAssertions: func(t *testing.T) {
				// clearing env variables in case the users already has variables set
				var err error
				err = os.Setenv("DBGROK_DB_HOST", "")
				err = os.Setenv("DBGROK_DB_NAME", "")
				err = os.Setenv("DBGROK_DB_USER", "")
				err = os.Setenv("DBGROK_DB_PASS", "")
				err = os.Setenv("DBGROK_DB_PORT", "")
				if err != nil {
					t.Error("Error setting default env variables")
				}
				currEnv := GetTestDBInfo()
				assert.Equal(t, "localhost", currEnv.Host)
				assert.Equal(t, "dbgrok", currEnv.Name)
				assert.Equal(t, "postgres_user", currEnv.User)
				assert.Equal(t, "postgres_pass", currEnv.Pass)
				assert.Equal(t, "5432", currEnv.Port)
			},
		},
	}

	// Run actual tests
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeAssertions(t)
		})
	}
}

func TestSetupConnection(t *testing.T) {
	tt := []struct {
		name           string
		makeAssertions func(*testing.T)
	}{
		{
			name: "Test Setup Connection",
			makeAssertions: func(t *testing.T) {
				sdb, err := db.SetupConnection(CreateDSN(GetTestDBInfo()))
				assert.NotNil(t, sdb)
				assert.NoError(t, err)
			},
		},
		{
			name: "Test Bad Setup Connection",
			makeAssertions: func(t *testing.T) {
				dsn := CreateDSN(TestDBInfo{Host: "", Name: "", User: "", Pass: "", Port: ""})
				assert.Equal(t, "postgres://:@:/?sslmode=disable", dsn)
				_, err := db.SetupConnection(dsn)
				assert.Error(t, err)
			},
		},
		{
			name: "Test DB Select",
			makeAssertions: func(t *testing.T) {
				var one int
				sdb, err := db.SetupConnection(CreateDSN(GetTestDBInfo()))
				assert.NotNil(t, sdb)
				assert.Nil(t, err)
				if sdb == nil {
					t.Error("*sql.DB is null")
					return
				}
				err = sdb.QueryRow("Select 1;").Scan(&one)
				assert.Nil(t, err)
				assert.Equal(t, 1, one)
			},
		},
	}

	// Run actual tests
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeAssertions(t)
		})
	}
}
