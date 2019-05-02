package models

// WriteToModel contains information to write out model template
type WriteToModel struct {
	ModelName        string
	VariableMappings []VariableMapping
	ModelVariables   []ModelVariable
	HasDate          bool
}

type VariableMapping struct {
	Source      string
	Dest        string
	DestLowered string
	Order       int
}

type ModelVariable struct {
	Name             string
	LowerName        string
	TypeName         string
	JSONName         string
	FieldValidations string
	HasValidations   bool
	ForeignKey       string
	HasForeignKey    bool
	IsPrimaryKey     bool
}

type ModelRepresentations struct {
	// BaseImportPath contains the base path of the project. example "github.com/brady1408/db-grok"
	// any generated import will be branched off of this base path. example "import (github.com/brady1408/db-grok/models)"
	BaseImportPath string
	Models         []ModelRepresentation
}

type ModelRepresentation struct {
	BaseImportPath     string
	HasPrimaryKey      bool
	HasSQLDateTime     bool
	ModelGenComment    string
	ModelName          string
	OverridePkName     string
	ShouldGenModelOnly bool
	SQLTypeAbbrev      string
	SQLTypeName        string
	TableName          string
	WriteToModel       map[string]WriteToModel
}
