package store

import (
	"os"
	"testing"

	"github.com/brady1408/db-grok/db"
	"github.com/brady1408/db-grok/test"
	"github.com/stretchr/testify/assert"
)

func TestGetAllPks(t *testing.T) {
	tt := []struct {
		name           string
		makeAssertions func(*testing.T)
	}{
		{
			name: "Test When PKs Exist",
			makeAssertions: func(t *testing.T) {
				sdb, err := db.SetupConnection(test.CreateDSN(test.GetTestDBInfo()))
				if err != nil {
					t.Error("Error connecting to database")
					return
				}
				pks, err := GetAllPks(sdb)
				assert.Greater(t, len(pks), 0)
				assert.NoError(t, err)
			},
		},
		{
			name: "Test When PKs Don't Exist",
			makeAssertions: func(t *testing.T) {
				prevDBName := os.Getenv("DBGROK_DB_NAME")
				os.Setenv("DBGROK_DB_NAME", "empty")
				sdb, err := db.SetupConnection(test.CreateDSN(test.GetTestDBInfo()))
				if err != nil {
					t.Error("Error connecting to database")
					return
				}
				pks, err := GetAllPks(sdb)
				assert.Equal(t, len(pks), 0)
				assert.NoError(t, err)
				os.Setenv("DBGROK_DB_NAME", prevDBName)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeAssertions(t)
		})
	}
}

func TestGetTableType(t *testing.T) {
	tt := []struct {
		name           string
		tableName      string
		makeAssertions func(*testing.T, string)
	}{
		{
			name:      "Good Table",
			tableName: "orders", // Needs to be standardized for boiler plate
			makeAssertions: func(t *testing.T, tn string) {
				sdb, err := db.SetupConnection(test.CreateDSN(test.GetTestDBInfo()))
				if err != nil {
					t.Fatal("Error connecting to database")
				}
				tableType, err := GetTableType(sdb, tn)
				if err != nil {
					t.Fatalf("Could not query table type for table %s. Error: %v", tn, err)
				}
				assert.Equal(t, "BASE TABLE", tableType)
			},
		},
		{
			name:      "Bad Table",
			tableName: "doesnt_exist",
			makeAssertions: func(t *testing.T, tn string) {
				sdb, err := db.SetupConnection(test.CreateDSN(test.GetTestDBInfo()))
				if err != nil {
					t.Fatal("Error connecting to database")
				}
				tableType, err := GetTableType(sdb, tn)
				assert.Error(t, err)
				assert.Equal(t, "", tableType)
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeAssertions(t, tc.tableName)
		})
	}
}
