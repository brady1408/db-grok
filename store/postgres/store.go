package postgres

import (
	"database/sql"

	"github.com/brady1408/db-grok/models"
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
				AND table_schema = ?;`,
		),
		tableName,
		"public",
	)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return tableType, nil
}

// GetPkName returns the column name of the primary key column. Likely id.
func GetPkName(db *sqlx.DB, tableName string) (string, error) {
	var column string
	err := db.Get(
		&column,
		db.Rebind(
			`SELECT Column_Name 
				FROM information_schema.key_column_usage 
				WHERE table_name = ? 
				AND table_schema = ?;`,
		),
		tableName,
		"public",
	)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	return column, err
}

// GetColumns returns list of columns in a table
func GetColumns(db *sqlx.DB, tableName string) ([]models.TableColumn, error) {
	tableColumns := []models.TableColumn{}
	err := db.Select(
		&tableColumns,
		db.Rebind(
			`SELECT
				  column_name
				, case is_nullable when 'YES' then true else false end as is_nullable
				, data_type
				, table_name
			 FROM information_schema.columns
			 WHERE table_name = ?
			 AND table_schema = ?;`,
		),
		tableName,
		"public",
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return tableColumns, err
}

// GetFks returns a list of foreign keys related to the parent table
func GetFks(db *sqlx.DB, parentTable string) ([]models.ForeignKeyInfo, error) {
	foreginKeyInfo := []models.ForeignKeyInfo{}
	err := db.Select(
		&foreginKeyInfo,
		db.Rebind(
			`SELECT
				  referenced_table_name
				, column_name
				, referenced_column_name
			 FROM information_schema.key_column_usage
			 WHERE table_name = ?
			 AND constraint_name != 'primary'
			 AND table_schema = ?;`,
		),
		parentTable,
		"public",
	)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return foreginKeyInfo, err
}
