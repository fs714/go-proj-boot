package migrate

import (
	"github.com/fs714/go-proj-boot/pkg/utils/config"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	"github.com/spf13/cobra"
)

var (
	dbHost  string
	dbPort  string
	dbUser  string
	dbPass  string
	dbName  string
	number  int
	version string
)

var StartCmd = &cobra.Command{
	Use:          "migrate",
	Short:        "Migrate database command line",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			cmd.Help()
		}

		return nil
	},
}

func InitStartCmd() {
	InitStartCreateCmd()
	InitStartUpCmd()
	InitStartDownCmd()
	InitStartGotoCmd()
	InitStartForceCmd()

	StartCmd.AddCommand(StartCreateCmd)
	StartCmd.AddCommand(StartUpCmd)
	StartCmd.AddCommand(StartDownCmd)
	StartCmd.AddCommand(StartGotoCmd)
	StartCmd.AddCommand(StartForceCmd)
}

var StartCreateCmd = &cobra.Command{
	Use:          "create [flags] name",
	Short:        "Create Database migration up/down files",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			log.Infow("migrate-created called", "name", args[0])
			return nil
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartCreateCmd() {
	addDatabaseFlags(StartCreateCmd)
}

var StartUpCmd = &cobra.Command{
	Use:          "up",
	Short:        "Apply all or N up migrations",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Infow("migrate-up called", "N", number)
			return nil
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartUpCmd() {
	addDatabaseFlags(StartUpCmd)
	StartUpCmd.Flags().IntVarP(&number, "number", "N", 1,
		"Number of Migration steps, 0 means all")
}

var StartDownCmd = &cobra.Command{
	Use:          "down",
	Short:        "Apply all or N down migrations",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Infow("migrate-down called", "N", number)
			return nil
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartDownCmd() {
	addDatabaseFlags(StartDownCmd)
	StartDownCmd.Flags().IntVarP(&number, "number", "N", 1,
		"Number of Migration steps, 0 means all")
}

var StartGotoCmd = &cobra.Command{
	Use:          "goto",
	Short:        "Migrate to version V",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Infow("migrate-goto called", "V", version)
			return nil
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartGotoCmd() {
	addDatabaseFlags(StartGotoCmd)
	StartGotoCmd.Flags().StringVarP(&version, "version", "V", "",
		"Migrate Version")
}

var StartForceCmd = &cobra.Command{
	Use:          "force",
	Short:        "Set version V but don't run migration (ignores dirty state)",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			log.Infow("migrate-force called", "V", version)
			return nil
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartForceCmd() {
	addDatabaseFlags(StartForceCmd)
	StartForceCmd.Flags().StringVarP(&version, "version", "V", "",
		"Migrate Version")
}

func addDatabaseFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().SortFlags = false
	cmd.Flags().SortFlags = false

	cmd.Flags().StringVarP(&dbHost, "host", "l", config.DefaultConfig.Database.Host,
		"Database server listening address")
	config.Viper.BindPFlag("database.host", cmd.Flags().Lookup("host"))
	config.Viper.BindEnv("database.port", "DATABASE_HOST")

	cmd.Flags().StringVarP(&dbPort, "port", "p", config.DefaultConfig.Database.Port,
		"Database server listening port")
	config.Viper.BindPFlag("database.port", cmd.Flags().Lookup("port"))
	config.Viper.BindEnv("database.port", "DATABASE_PORT")

	cmd.Flags().StringVarP(&dbUser, "user", "", config.DefaultConfig.Database.User,
		"Username for Database connection")
	config.Viper.BindPFlag("database.user", cmd.Flags().Lookup("user"))
	config.Viper.BindEnv("database.user", "DATABASE_USER")

	cmd.Flags().StringVarP(&dbPass, "pass", "", config.DefaultConfig.Database.Pass,
		"Password for Database connection")
	config.Viper.BindPFlag("database.pass", cmd.Flags().Lookup("pass"))
	config.Viper.BindEnv("database.pass", "DATABASE_PASS")

	cmd.Flags().StringVarP(&dbName, "name", "", config.DefaultConfig.Database.Name,
		"Database name to connect")
	config.Viper.BindPFlag("database.name", cmd.Flags().Lookup("name"))
	config.Viper.BindEnv("database.name", "DATABASE_NAME")
}
