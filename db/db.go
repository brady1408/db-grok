package db

import (
	"context"
	"os"
	"time"

	"database/sql"

	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"

	// Importing for database connection.
	_ "github.com/lib/pq"
)

const (
	// DB_PING_ATTEMPTS is the number of times a ping will be attempted after connecting with the database before application quits
	DB_PING_ATTEMPTS int = 3
	// DB_PING_TIMEOUT_SECS is the time between ping attempts
	DB_PING_TIMEOUT_SECS time.Duration = 3 * time.Second
)

// SetupConnection will connect to a database and test the connection with a ping.
func SetupConnection(dsn string) (*sql.DB, error) {
	// Case the dns schema to find the type of database
	log := logrus.New()
	log.SetOutput(os.Stdout)
	// log.SetLevel(logrus.DebugLevel)
	log.Info("Connecting to database server")
	// Skipping error checking on Open because Open will not always make a connection to the databae.
	// This reflects negatively on our tests. The connection test will be found below in the Ping.
	// Ping will open a connection if one does not yet exists, which is rare that it exists if nothing
	// has tried to use it yet.
	db, _ := sql.Open("postgres", dsn)
	log.Infof("opened connection with dsn: %s", dsn)

	// for loop to retry ping. This is configurable in the consts (DB_PING_ATTEMPTS and DB_PING_TIMEOUT_SECS)
	for i := 0; i < DB_PING_ATTEMPTS; i++ {
		log.Info("Pinging SQL postgres database")
		ctx, cancel := context.WithTimeout(context.Background(), DB_PING_TIMEOUT_SECS)
		defer cancel()
		err := db.PingContext(ctx)
		if err == nil {
			break
		} else {
			if i == DB_PING_ATTEMPTS-1 {
				err = errorx.Decorate(err, "Failed to ping DB, serer will exit err:")
				return nil, err
			}
			err = errorx.Decorate(err, "Failed to ping DB retrying in %d seconds err:", DB_PING_TIMEOUT_SECS)
			time.Sleep(DB_PING_TIMEOUT_SECS)
		}
	}
	// Logging successfully connected only after a successful ping to the database for reasons described above.
	// See comments above sql.Open()
	log.Info("Successfully Connected to database server")

	return db, nil
}
