package command

import (
	"github.com/brady1408/db-grok/config"
	"github.com/brady1408/db-grok/db"
	"github.com/brady1408/db-grok/models"
	"github.com/brady1408/db-grok/store"
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

	allPrimaryKeys, err := store.GetAllPks(sdb)
	if err != nil {
		log.Errorf("Error getting primary keys: %v", err)
		return err
	}

	for _, table := range cfg.Models {
		//Check if tables were sent in args and include only tables in args
		if !inArgs(viper.GetStringSlice("tables"), table.FromTableName) {
			continue
		}

		//SKIP all generation if this property is set to true
		if table.Properties.IgnoreForGeneration {
			continue
		}

		tableName := table.FromTableName
		tableType, err := store.GetTableType(sdb, tableName)
		if err != nil {
			return errorx.Decorate(err, "Error getting table type: ")
		}
	}
	//TEMP
	log.Infof("%v", allPrimaryKeys)

	return nil //TODO FIX NIL

}
