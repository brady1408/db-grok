package test

import (
	"fmt"
	"os"
)

// TestDBInfo hold database connection information for the test database.
type TestDBInfo struct {
	Host string
	Name string
	User string
	Pass string
	Port string
}

type testArgs struct {
	nilConnection bool
	badConnection bool
}

// getTestDBInfo returns a object with all the database connections information in it.
func GetTestDBInfo() TestDBInfo {
	testDBInfo := TestDBInfo{}
	testDBInfo.Host = os.Getenv("DBGROK_DB_HOST")
	if testDBInfo.Host == "" {
		testDBInfo.Host = "localhost"
	}
	testDBInfo.Name = os.Getenv("DBGROK_DB_NAME")
	if testDBInfo.Name == "" {
		testDBInfo.Name = "dbgrok"
	}
	testDBInfo.User = os.Getenv("DBGROK_DB_USER")
	if testDBInfo.User == "" {
		testDBInfo.User = "postgres_user"
	}
	testDBInfo.Pass = os.Getenv("DBGROK_DB_PASS")
	if testDBInfo.Pass == "" {
		testDBInfo.Pass = "postgres_pass"
	}
	testDBInfo.Port = os.Getenv("DBGROK_DB_PORT")
	if testDBInfo.Port == "" {
		testDBInfo.Port = "5432"
	}
	return testDBInfo
}

func CreateDSN(t TestDBInfo) string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", t.User, t.Pass, t.Host, t.Port, t.Name)
}
