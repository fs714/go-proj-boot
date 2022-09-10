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

var DBSqlx *sqlx.DB

var (
	//go:embed migrations/*.sql
	MigratesFs embed.FS
)

func InitSqlxFromConfig() (err error) {
	err = InitSqlx(
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

func InitSqlx(host string, port string, user string, password string, dbName string) (err error) {
	count := 0
	for {
		if count >= 30 {
			break
		}
		count++

		err = doInitSqlx(host, port, user, password, dbName)
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

func doInitSqlx(host string, port string, user string, password string, dbName string) (err error) {
	dbUrl := "postgres://" + user + ":" + password + "@" + host + ":" + port + "/" + dbName + "?sslmode=disable"
	log.Infow("database", "url", dbUrl)
	DBSqlx, err = sqlx.Connect("pgx", dbUrl)
	if err != nil {
		return errors.Wrap(err, "failed to connection to PostgreSql DB")
	}

	DBSqlx.DB.SetMaxOpenConns(config.Config.Database.Master.MaxOpenConnection)
	DBSqlx.DB.SetMaxIdleConns(config.Config.Database.Master.MaxIdleConnection)
	DBSqlx.DB.SetConnMaxLifetime(time.Duration(config.Config.Database.Master.MaxLifeTime) * time.Second)

	err = DBSqlx.Ping()
	if err != nil {
		return errors.Wrap(err, "failed to ping PostgreSql DB")
	}

	return
}
