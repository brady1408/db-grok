package store

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// GetAllPks Return map with a list of all primary keys in the database or an error
func GetAllPks(db *sqlx.DB) (map[string]string, error) {
	rows, err := db.Query(
		`SELECT kcu.table_name, kcu.column_name
			FROM information_schema.key_column_usage kcu
			JOIN information_schema.table_constraints tco
				ON tco.constraint_name = kcu.constraint_name
				AND tco.constraint_schema = kcu.constraint_schema
				AND tco.constraint_name = kcu.constraint_name
			WHERE tco.constraint_type = 'PRIMARY KEY'
			ORDER BY kcu.table_name,
					 kcu.column_name;`)
	if err != nil && err != sql.ErrNoRows {
		panic("could not get primary keys") // NO PANICS!!!
	}
	primaryKeys := map[string]string{}
	for rows.Next() {
		var (
			tableName, columnName string
		)

		err := rows.Scan(&tableName, &columnName)
		if err != nil {
			return nil, err
		}
		primaryKeys[tableName] = columnName
	}
	return primaryKeys, nil
}

// GetTableType returns the type of table of the table passed in tableName
func GetTableType(db *sqlx.DB, tableName string) (string, error) {
	var tableType string
	err := db.Get(
		&tableType,
		db.Rebind(
			`SELECT table_type 
				FROM information_schema.tables 
				WHERE table_name = ? 
				AND table_schema = ?`,
		),
		tableName,
		"public",
	)
	if err != nil {
		return "", err
	}

	return tableType, nil
}
