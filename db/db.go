package db

import (
	"context"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
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
func SetupConnection(dsn string) (*sqlx.DB, error) {
	// Case the dns schema to find the type of database
	log := logrus.New()
	log.SetOutput(os.Stdout)
	// log.SetLevel(logrus.DebugLevel)
	log.Info("Connecting to database server")
	sdb, err := sqlx.Open("postgres", dsn)
	log.Infof("opened connection with dsn: %s", dsn)

	// for loop to retry ping. This is configurable in the consts (DB_PING_ATTEMPTS and DB_PING_TIMEOUT_SECS)
	for i := 0; i < DB_PING_ATTEMPTS; i++ {
		log.Info("Pinging SQL postgres database")
		ctx, cancel := context.WithTimeout(context.Background(), DB_PING_TIMEOUT_SECS)
		defer cancel()
		err = sdb.PingContext(ctx)
		if err == nil {
			break
		} else {
			if i == DB_PING_ATTEMPTS-1 {
				err = errorx.Decorate(err, "Failed to ping DB, server will exit err:")
				return sdb, err
			}
			err = errorx.Decorate(err, "Failed to ping DB retrying in %d seconds err:", DB_PING_TIMEOUT_SECS)
			time.Sleep(DB_PING_TIMEOUT_SECS)
		}
	}
	// Logging successfully connected only after a successful ping to the database
	log.Info("Successfully Connected to database server")

	return sdb, err
}
