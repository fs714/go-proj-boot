package pgsql

import (
	"embed"
	"time"

	"github.com/fs714/go-proj-boot/pkg/config"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
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
		config.Config.Database.Master.Nodes[0].Host,
		config.Config.Database.Master.Nodes[0].Port,
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
			log.Infow("failed to connect to db", "host", host, "port", port, "count", count)
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

	DB.DB.SetMaxOpenConns(config.Config.Database.Master.MaxOpenConnection)
	DB.DB.SetMaxIdleConns(config.Config.Database.Master.MaxIdleConnection)
	DB.DB.SetConnMaxLifetime(time.Duration(config.Config.Database.Master.MaxLifeTime) * time.Second)

	err = DB.Ping()
	if err != nil {
		return errors.Wrap(err, "failed to ping PostgreSql DB")
	}

	return
}
