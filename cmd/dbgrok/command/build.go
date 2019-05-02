package command

import (
	"bytes"
	"fmt"

	"bitbucket.org/atlascloudapp/modelbuilder/dataaccess"
	"bitbucket.org/atlascloudapp/modelbuilder/utils"
	"github.com/brady1408/db-grok/config"
	"github.com/brady1408/db-grok/db"
	"github.com/brady1408/db-grok/models"
	"github.com/brady1408/db-grok/store/postgres"
	"github.com/iancoleman/strcase"
	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var BuildCmd = &cobra.Command{
	Use:     "build",
	Short:   "Build go files",
	Long:    "Connect to database and create go files from schema",
	Example: `  build --out store`,
	RunE:    buildCmdF,
}

func init() {
	BuildCmd.Flags().StringP("out", "o", ".", "project root directory to build into")
	BuildCmd.Flags().StringP("config", "c", "builder_config.json", "Path to the builder_config.json file")
	BuildCmd.Flags().StringSliceP("tables", "t", []string{}, "build for a single table")
}

func buildCmdF(command *cobra.Command, args []string) error {
	log := logrus.New()
	log.Info("Building...")
	cfg, err := config.LoadConfig(viper.GetString("config"))
	if err != nil {
		log.Errorf("Error Loading Config: %v", err)
		return err
	}
	log.Info("Config loaded successfully")

	sdb, err := db.SetupConnection(cfg.ConnectionString)
	if err != nil {
		log.Errorf("Error setting up the database connection. Err: %v", err)
		return err
	}

	modelReps := &models.ModelRepresentations{
		BaseImportPath: cfg.BaseImportPath,
	}

	// TEMP
	log.Infof("%v", modelReps)

	modelList := map[string]string{}
	// Check to see if table name was included in args.
	// if it was add it to the modelList.
	// this allows us to build for only a specific list of tables
	// and not the whole BuildConfig
	for _, table := range cfg.Models {
		if !inArgs(viper.GetStringSlice("tables"), table.FromTableName) {
			continue
		}
		modelList[table.FromTableName] = table.FromTableName
	}

	allPrimaryKeys, err := postgres.GetAllPks(sdb)
	if err != nil {
		log.Errorf("Error getting primary keys: %v", err)
		return err
	}

	for _, table := range cfg.Models {
		var (
			comments string
			markdown bytes.Buffer
		)
		// modVars := []models.ModelVariable{}
		modelBuild := make(map[string]models.WriteToModel)

		fmt.Println("processing " + table.FromTableName)
		markdown.WriteString(fmt.Sprintf("%s\r\n\r\n", table.FromTableName))
		markdown.WriteString("|Property|Type|Comments|\r\n")
		markdown.WriteString("|--------|----|--------|\r\n")

		//Check if tables were sent in args and include only tables in args
		if !inArgs(viper.GetStringSlice("tables"), table.FromTableName) {
			continue
		}

		//SKIP all generation if this property is set to true
		if table.Properties.IgnoreForGeneration {
			continue
		}

		prettyTableName := strcase.ToLowerCamel(table.FromTableName)
		tableType, err := postgres.GetTableType(sdb, table.FromTableName)
		if err != nil {
			return errorx.Decorate(err, "Error getting table type: ")
		}
		isView := tableType == "VIEW"
		pkName, err := postgres.GetPkName(sdb, table.FromTableName)
		if err != nil {
			return errorx.Decorate(err, "Error getting primary key name, is there a pk? if not you should set the property noPK on this table in the config")
		}
		if table.Properties.OverridePkName != "" {
			pkName = table.Properties.OverridePkName
		}
		tableColumns, err := postgres.GetColumns(sdb, table.FromTableName)
		if err != nil {
			return errorx.Decorate(err, "Error getting table columns: ")
		}
		for k, v := range tableColumns {
			var (
				hasDate     bool
				jsonName    string
				validations string
			)
			if isView && v.DataType == "datetime" {
				tableColumns[k].IsNullable = true
			}
			isPrimaryKey := pkName == v.ColumnName
			prettyColumnName := strcase.ToLowerCamel(v.ColumnName)
			columnDataType := postgres.ConvertDataType(v)

			if jsonOverride, ok := table.Properties.JsonOverrideColumns[v.ColumnName]; !ok {
				jsonName = prettyColumnName
				markdown.WriteString(fmt.Sprintf("|%s|%s|%s|\r\n", prettyColumnName, columnDataType, comments))
			} else {
				jsonName = jsonOverride
			}
			if val, ok := table.Properties.Validations[v.ColumnName]; ok {
				validations = val
			}
			if columnDataType == "timestamp with time zone" {
				hasDate = true
			}

			modVar := models.ModelVariable{
				TypeName:         columnDataType,
				Name:             v.ColumnName,
				LowerName:        prettyColumnName,
				JSONName:         jsonName,
				FieldValidations: validations,
				HasValidations:   validations != "",
				IsPrimaryKey:     isPrimaryKey,
			}
			// modVars = append(modVars, modVar) // WHY?
			if _, ok := modelBuild[table.FromTableName]; !ok {
				if table.OverrideModelName != "" {
					modelBuild[table.FromTableName] = models.WriteToModel{
						ModelName:        table.OverrideModelName,
						VariableMappings: []models.VariableMapping{},
						HasDate:          hasDate,
					}
				} else {
					modelBuild[table.FromTableName] = models.WriteToModel{
						ModelName:        table.FromTableName,
						VariableMappings: []models.VariableMapping{},
						HasDate:          hasDate,
					}
				}
			}
			obj, _ := modelBuild[table.FromTableName]
			// obj.VariableMappings = append(obj.VariableMappings, varmap)
			obj.ModelVariables = append(obj.ModelVariables, modVar)
			obj.HasDate = hasDate
			modelBuild[table.FromTableName] = obj
		}
		if table.Properties.ParentFKType == "" {
			table.Properties.ParentFKType = "int"
		}
		if table.OverrideModelName != "" {
			modelName = table.OverrideModelName
		} else {
			modelName := tableName
		}

		m := &models.ModelRepresentation{
			BaseImportPath:     utils.Config.BaseImportPath,
			HasPrimaryKey:      pkName != "",
			HasSqlDateTime:     hasDate,
			ModelGenComment:    table.Properties.ModelGenComment,
			ModelName:          modelName,
			OverridePkName:     table.Properties.OverridePkName,
			ShouldGenModelOnly: table.Properties.ShouldGenModelOnly,
			SqlTypeAbbrev:      "pd",
			SqlTypeName:        fmt.Sprintf("Sql%vData", modelName),
			TableName:          utils.PrimaryTable(tableName),
			WriteToModel:       modelBuild,
		}

		safeFileName := prettyTableName

		for _, v := range modelBuild {
			//don't generate model if specified
			if table.Properties.DoNotGenerateModel {
				continue
			}

			om := &models.OutputModel{
				ModelName:       v.ModelName,
				ModelVariables:  v.ModelVariables,
				HasDate:         v.HasDate,
				ModelGenComment: table.Properties.ModelGenComment,
			}

			fks := dataaccess.GetFks(v.ModelName)
		}

	}
	//TEMP
	log.Infof("%v", allPrimaryKeys)

	return nil //TODO FIX NIL

}

func buildModels() error {

}
