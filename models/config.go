package models

type EmptyObject struct {
}

// BuildConfig is a container for the config.json sent to the application as a parameter of the Build command.
type BuildConfig struct {
	ConnectionString string
	BaseImportPath   string
	Models           []ModelConfig
}

type ModelRepresentations struct {
	BaseImportPath string
	Models         []ModelRepresentation
}

type ModelVariable struct {
	Name             string
	LowerName        string
	TypeName         string
	JsonName         string
	FieldValidations string
	HasValidations   bool
	ForeignKey       string
	HasForeignKey    bool
	IsPrimaryKey     bool
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

type WriteToModel struct {
	ModelName        string
	VariableMappings []VariableMapping
	ModelVariables   []ModelVariable
	HasDate          bool
	ShouldUseGorm    bool
}

type ModelRepresentation struct {
	BaseImportPath            string
	SqlTypeName               string
	InterfaceName             string
	SqlTypeAbbrev             string
	SelectColumnList          string
	TableName                 string
	FromStatement             string
	ModelName                 string
	ModelGenComment           string
	ScanVariableAssigments    []ScanVariableAssigment
	ScanArguments             string
	VariableMappings          []VariableMapping
	PrimaryVariableMappings   []VariableMapping
	PrimaryKeyPropertyName    string
	PrimaryKeyColumnName      string
	ShouldGenerateIsValid     bool
	ShouldGeneratePreSave     bool
	ShouldGeneratePreSaveAPI  bool
	ShouldGeneratePostSaveAPI bool
	ShouldIgnoreUpdatedDate   bool
	ShouldGenAPI              bool
	InsertParams              string
	UpdateParams              string
	ExecParams                string
	InsertColumnList          string
	InsertVarMappings         []VariableMapping
	UpdateVarMappings         []VariableMapping
	UpdateColumnList          string
	ParentFK                  string
	ParentFKType              string
	ParentFKLower             string
	FKHeirarchy               []FKInfo
	ShouldIgnoreParentFK      bool
	IsNotParentFKFacility     bool
	HasUpdateDate             bool
	ShouldImplementPaging     bool
	PagingField               string
	ShouldIncludeUtilsForSql  bool
	ShouldBulkInsert          bool
	ShouldGenerateUpsert      bool
	ShouldFilterByHeirarchy   bool
	ShouldOrderEntities       bool
	OrderBy                   string
	ShouldImplementSorting    bool
	ShouldImplementFilter     bool
	FilterColumns             []string
	HasPrimaryKey             bool
	HasSqlDateTime            bool
	ShouldValidate            bool
	ShouldAllowFilterDeleted  bool
	ShouldIncludeUtils        bool
	HasCacheControl           bool
	CacheControl              string
	ExtraSelectMethods        []SelectMethod
	ExtraDeleteMethods        []DeleteMethod
	ShouldDeleteByFK          bool
	ShouldGetOnly             bool
	WriteToModel              map[string]WriteToModel
	IsComposite               bool
	ShouldHardDelete          bool
	Routines                  []Routine
	IgnoreApiUpdate           []string
	OverridePkName            string
	IsReadOnly                bool
	IncludeViewRelation       []string
	ShouldSkipRelationships   bool
	ShouldUseGorm             bool
	ShouldGenModelOnly        bool
	ShouldIncludeReferences   bool
}

type ScanVariableAssigment struct {
	Text     string
	Name     string
	TypeName string
}

type VariableMapping struct {
	Source      string
	Dest        string
	DestLowered string
	Order       int
}

type ModelConfig struct {
	FromTableName     string
	OverrideModelName string
	Properties        ModelProps
	Validations       Validations
}

type ValidationBase struct {
}

type Validations struct {
	MinLength MinLength
}

type MinLength struct {
	MinLength int
}

type ModelProps struct {
	ShouldPreSave              bool
	ShouldPreSaveApi           bool
	ShouldPostSaveApi          bool
	ShouldCheckIsValid         bool
	ShouldIgnoreUpdatedDate    bool
	ShouldIgnoreParentFK       bool
	ParentFK                   string
	ParentFKType               string
	ShouldGenAPI               bool
	FKHeirarchy                []FKInfo
	ShouldImplementPaging      bool
	PagingField                string
	ShouldGenerateUpsert       bool
	ShouldOrderEntities        bool
	OrderBy                    string
	ShouldImplementSorting     bool
	JsonOverrideColumns        map[string]string
	ApiPath                    string
	ApiMethods                 map[string]string
	ShouldFilterByHeirarchy    bool
	ShouldBulkInsert           bool
	ShouldValidate             bool
	Validations                map[string]string
	ModelGenComment            string
	ShouldAllowFilterDeleted   bool
	CacheControl               string
	ExtraSelectMethods         []SelectMethod
	ExtraDeleteMethods         []DeleteMethod
	ShouldDeleteByFK           bool
	ShouldGetOnly              bool
	ShouldHardDelete           bool
	Routines                   []string
	IgnoreApiUpdate            []string
	OverridePkName             string
	IsReadOnly                 bool
	IncludeViewRelation        []string
	ShouldSkipRelationships    bool
	ShouldUseGorm              bool
	ShouldGenModelOnly         bool
	ShouldIncludeReferences    bool
	IgnoreForGeneration        bool //you don't want anything generated
	DontGenerateModel          bool //You are using a custom model
	DontGenerateSqlDataContext bool //you are using custom sql_blah_data so don't generate base class (no inheritance/overriding in gorm <frown>
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
