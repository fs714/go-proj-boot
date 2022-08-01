package pgsql

import (
	"embed"
	"time"

	"github.com/fs714/go-proj-boot/pkg/utils/config"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

var DB *sqlx.DB

var (
	//go:embed migrations/*.sql
	MigratesFs embed.FS
)

func PostgreDbInitFromConfig() (err error) {
	err = PostgreDbInit(
		config.Config.Database.Host,
		config.Config.Database.Port,
		config.Config.Database.User,
		config.Config.Database.Pass,
		config.Config.Database.Name,
	)

	if err != nil {
		errors.WithMessage(err, "failed to init db from config")
		return
	}

	return
}

func PostgreDbInit(host string, port string, user string, password string, dbName string) (err error) {
	count := 0
	for {
		if count >= 30 {
			break
		}
		count++

		err = doPostgreDbInit(host, port, user, password, dbName)
		if err != nil {
			time.Sleep(10 * time.Second)
			continue
		} else {
			break
		}
	}

	return
}

func doPostgreDbInit(host string, port string, user string, password string, dbName string) (err error) {
	dbUrl := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName + "?sslmode=disable"
	DB, err = sqlx.Connect("pgx", dbUrl)
	if err != nil {
		return errors.Wrap(err, "failed to connection to PostgreSql DB")
	}

	DB.DB.SetMaxOpenConns(config.Config.Database.MaxOpenConnection)
	DB.DB.SetMaxIdleConns(config.Config.Database.MaxIdleConnection)
	DB.DB.SetConnMaxLifetime(time.Duration(config.Config.Database.MaxLifeTime) * time.Second)

	err = DB.Ping()
	if err != nil {
		return errors.Wrap(err, "failed to ping PostgreSql DB")
	}

	return
}
