package models

// TableColumn information about the sql table columns
type TableColumn struct {
	ColumnName    string
	DataType      string
	FromTableName string
	IsNullable    bool
}

// ForeignKeyInfo information about the parent table foregin keys
type ForeignKeyInfo struct {
	tableName            string
	columnName           string
	referencedColumnName string
}
