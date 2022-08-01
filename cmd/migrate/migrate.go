package migrate

import (
	"github.com/fs714/go-proj-boot/db/pgsql"
	"github.com/fs714/go-proj-boot/pkg/utils/config"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	dbHost string
	dbPort string
	dbUser string
	dbPass string
	dbName string
	number int
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
	InitStartShowCmd()
	InitStartUpCmd()
	InitStartDownCmd()
	InitStartGotoCmd()
	InitStartForceCmd()

	StartCmd.AddCommand(StartShowCmd)
	StartCmd.AddCommand(StartUpCmd)
	StartCmd.AddCommand(StartDownCmd)
	StartCmd.AddCommand(StartGotoCmd)
	StartCmd.AddCommand(StartForceCmd)
}

var StartShowCmd = &cobra.Command{
	Use:          "show",
	Short:        "Show the currently active migration version",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return ShowCmd()
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartShowCmd() {
	addDatabaseFlags(StartShowCmd)
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
	Use:          "goto [flags] V",
	Short:        "Migrate to version V",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			log.Infow("migrate-goto called", "V", args[0])
			return nil
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartGotoCmd() {
	addDatabaseFlags(StartGotoCmd)
}

var StartForceCmd = &cobra.Command{
	Use:          "force [flags] V",
	Short:        "Set version V but don't run migration (ignores dirty state)",
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			log.Infow("migrate-force called", "V", args[0])
			return nil
		} else {
			cmd.Help()
			return nil
		}
	},
}

func InitStartForceCmd() {
	addDatabaseFlags(StartForceCmd)
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

func ShowCmd() (err error) {
	err = pgsql.PostgreDbInitFromConfig()
	if err != nil {
		log.Errorf("show cmd failed during init db:%+v", err)
		return
	}

	driver, err := postgres.WithInstance(pgsql.DB.DB, &postgres.Config{})
	if err != nil {
		errors.Wrap(err, "show cmd failed during new migrate db driver")
		log.Errorf("%+v", err)
		return
	}

	d, err := iofs.New(pgsql.MigratesFs, "db/migrations")
	if err != nil {
		errors.Wrap(err, "show cmd failed during new iofs from embed")
		log.Errorf("%+v", err)
		return
	}

	m, err := migrate.NewWithInstance("iofs", d, "postgres", driver)
	if err != nil {
		errors.Wrap(err, "show cmd failed during new migrate instance")
		log.Errorf("%+v", err)
		return
	}

	version, dirty, err := m.Version()
	if err != nil {
		errors.Wrap(err, "show cmd failed during invoke Version")
		log.Errorf("%+v", err)
		return
	}

	log.Infow("Current active migration version", "version", version, "dirty", dirty)

	return
}
