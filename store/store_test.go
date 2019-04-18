package store

import (
	"testing"

	"bitbucket.org/atlascloudapp/db-grok/db"
	"bitbucket.org/atlascloudapp/db-grok/test"
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
				}
				pks, err := GetAllPks(sdb)
				assert.Greater(t, len(pks), 0)
				assert.NoError(t, err)
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tc.makeAssertions(t)
		})
	}
}
