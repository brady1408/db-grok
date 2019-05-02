package models

type EmptyObject struct {
}

// BuildConfig is a container for the config.json sent to the application as a parameter of the Build command.
type BuildConfig struct {
	ConnectionString string
	BaseImportPath   string
	Models           []ModelConfig
}

type ModelConfig struct {
	FromTableName     string
	OverrideModelName string
	Properties        ModelProps
	Validations       Validations
}

type ModelProps struct {
	JSONOverrideColumns map[string]string
	Validations         map[string]string
	ModelGenComment     string
	OverridePkName      string
	ShouldGenModelOnly  bool
	IgnoreForGeneration bool //you don't want anything generated
	DoNotGenerateModel  bool //You are using a custom model
}

type OutputModel struct {
	ModelName       string
	ModelGenComment string
	ModelVariables  []ModelVariable
	HasDate         bool
}

type FKInfo struct {
	Name string
}

type ScanVariableAssigment struct {
	Text     string
	Name     string
	TypeName string
}

type ValidationBase struct {
}

type Validations struct {
	MinLength MinLength
}

type MinLength struct {
	MinLength int
}

type SelectMethod struct {
	Parameters               []Parameter
	Conditions               []Condition
	SingleRow                bool
	MethodName               string
	ShouldIncludeInApi       bool
	ShouldImplementPaging    bool
	ShouldAllowFilterDeleted bool
	PagingField              string
}

type DeleteMethod struct {
	Parameters []Parameter
	MethodName string
}

type Condition struct {
	Name              string
	Comparison        string
	CompareTo         string
	CompareToConstant string
}

type Parameter struct {
	Name        string
	Type        string
	Comparison  string
	Alias       string
	ExistsTable string
}

type Routine struct {
	Name   string
	Params []RoutineParameter
}

type RoutineParameter struct {
	Name string
	Type string
}
