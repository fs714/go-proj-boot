package pgsql

import (
	"fmt"
	"time"

	"github.com/fs714/go-proj-boot/pkg/config"
	"github.com/fs714/go-proj-boot/pkg/utils/log"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/linxGnu/mssqlx"
	"github.com/pkg/errors"
)

var DBMssqlx *mssqlx.DBs

func InitMssqlxFromConfig() (err error) {
	err = InitMssqlx(config.Config.Database)
	if err != nil {
		errors.WithMessage(err, "failed to init db from config")
		return
	}

	return
}

func InitMssqlx(dbConfig config.Database) (err error) {
	count := 0
	for {
		if count >= 30 {
			break
		}
		count++

		err = doInitMssqlx(dbConfig)
		if err != nil {
			log.Errorf("failed to init db by mssqlx %d times with %+v", count, err)
			DBMssqlx.Destroy()
			time.Sleep(10 * time.Second)
			continue
		} else {
			break
		}
	}

	return
}

func doInitMssqlx(dbConfig config.Database) (err error) {
	var masterDSNs, slaveDSNs []string

	for _, n := range dbConfig.Master.Nodes {
		log.Infow("master db", "url", getDataSourceName(n.Host, n.Port, dbConfig.User, dbConfig.Pass, dbConfig.Name))
		masterDSNs = append(masterDSNs, getDataSourceName(n.Host, n.Port, dbConfig.User, dbConfig.Pass, dbConfig.Name))
	}

	for _, n := range dbConfig.Slave.Nodes {
		log.Infow("slave db", "url", getDataSourceName(n.Host, n.Port, dbConfig.User, dbConfig.Pass, dbConfig.Name))
		slaveDSNs = append(slaveDSNs, getDataSourceName(n.Host, n.Port, dbConfig.User, dbConfig.Pass, dbConfig.Name))
	}

	var errs []error
	DBMssqlx, errs = mssqlx.ConnectMasterSlaves("pgx", masterDSNs, slaveDSNs, mssqlx.WithReadQuerySource(mssqlx.ReadQuerySourceSlaves))
	for _, err = range errs {
		if err != nil {
			log.Errorf("failed to connect master and slave by mssqlx with err: %s", err.Error())
			return errors.Wrap(err, "failed to connect master and slave by mssqlx")
		}
	}

	DBMssqlx.SetHealthCheckPeriod(500)

	DBMssqlx.SetMasterMaxOpenConns(dbConfig.Master.MaxOpenConnection)
	DBMssqlx.SetMasterMaxIdleConns(dbConfig.Master.MaxIdleConnection)
	DBMssqlx.SetMasterConnMaxLifetime(time.Duration(dbConfig.Master.MaxLifeTime) * time.Second)

	DBMssqlx.SetSlaveMaxOpenConns(dbConfig.Master.MaxOpenConnection)
	DBMssqlx.SetSlaveMaxIdleConns(dbConfig.Master.MaxIdleConnection)
	DBMssqlx.SetSlaveConnMaxLifetime(time.Duration(dbConfig.Master.MaxLifeTime) * time.Second)

	errs = DBMssqlx.Ping()
	for _, err = range errs {
		if err != nil {
			return errors.Wrap(err, "failed to connect master and slave by mssqlx")
		}
	}

	return
}

func getDataSourceName(host string, port string, user string, password string, dbName string) (dsn string) {
	dsnTemplate := "postgres://%s:%s@%s:%s/%s?sslmode=disable"
	dsn = fmt.Sprintf(dsnTemplate, user, password, host, port, dbName)
	return
}
