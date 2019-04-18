package store

import (
	"database/sql"
)

// GetAllPks Return map with a list of all primary keys in the database or an error
func GetAllPks(db *sql.DB) (map[string]string, error) {
	rows, err := db.Query(
		`select kcu.table_name, kcu.column_name
		from information_schema.key_column_usage kcu
		join information_schema.table_constraints tco
			 on tco.constraint_name = kcu.constraint_name
			 and tco.constraint_schema = kcu.constraint_schema
			 and tco.constraint_name = kcu.constraint_name
		where tco.constraint_type = 'PRIMARY KEY'
		order by kcu.table_name,
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
