package command

import (
	"bitbucket.org/atlascloudapp/db-grok/config"
	"bitbucket.org/atlascloudapp/db-grok/db"
	"bitbucket.org/atlascloudapp/db-grok/models"
	"bitbucket.org/atlascloudapp/db-grok/store"
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
	BuildCmd.Flags().StringP("config", "c", "builer_config.json", "Path to the builder_conig.json file")
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

	db, err := db.SetupConnection(cfg.ConnectionString)
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

	allPrimaryKeys, err := store.GetAllPks(db)
	if err != nil {
		log.Errorf("Error getting primary keys: %v", err)
		return err
	}
	//TEMP
	log.Infof("%v", allPrimaryKeys)

	return nil //TODO FIX NIL

}
