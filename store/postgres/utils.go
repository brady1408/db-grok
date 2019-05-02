package postgres

import "github.com/brady1408/db-grok/models"

// ConvertDataType returns golang types for pq columns
func ConvertDataType(tc models.TableColumn) string {
	if tc.DataType == "" {
		return ""
	}

	if tc.DataType == "int" && !tc.IsNullable {
		return "int"
	} else if tc.DataType == "int" && tc.IsNullable {
		return "*int"
	} else if tc.DataType == "bigint" && !tc.IsNullable {
		return "int64"
	} else if tc.DataType == "boolean" {
		return "bool"
	} else if (tc.DataType == "decimal" || tc.DataType == "numeric") && !tc.IsNullable {
		return "float32"
	} else if (tc.DataType == "decimal" || tc.DataType == "numeric") && tc.IsNullable {
		return "*float32"
	} else if tc.DataType == "bigint" && tc.IsNullable {
		return "sql.NullInt64"
	} else if tc.DataType == "varchar" && tc.IsNullable {
		return "sql.NullString"
	} else if (tc.DataType == "text" || tc.DataType == "mediumtext" || tc.DataType == "longtest") && tc.IsNullable {
		return "sql.NullString"
	} else if (tc.DataType == "json" || tc.DataType == "jsonb") && tc.IsNullable {
		return "sql.NullString"
	} else if tc.DataType == "varchar" && tc.IsNullable {
		return "*string"
	} else if tc.DataType == "varchar" && !tc.IsNullable {
		return "string"
	} else if tc.DataType == "timestamp with time zone" && !tc.IsNullable {
		return "time.Time"
	} else if tc.DataType == "timestamp with time zone" && tc.IsNullable {
		return "*time.Time"
	} else if tc.DataType == "blob" || tc.DataType == "mediumblob" {
		return "[]byte"
	} else if tc.DataType == "bytea" {
		return "[]byte"
	} else if tc.DataType == "bit" {
		return "[]byte"
	}
	return "string"
}
