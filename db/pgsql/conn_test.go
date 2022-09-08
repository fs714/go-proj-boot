package pgsql

import (
	"testing"

	"github.com/fs714/go-proj-boot/pkg/config"
)

func TestInitSqlx(t *testing.T) {
	host := "192.168.75.230"
	port := "5432"
	user := "mikasa"
	pass := "mikasa"
	name := "titan"
	err := InitSqlx(host, port, user, pass, name)
	if err != nil {
		t.FailNow()
		t.Logf("failed to init DB by sqlx, %s:%s@%s:%s/%s\n", user, pass, host, port, name)
	}
}

func TestInitMssqlx(t *testing.T) {
	dbConfig := config.Database{
		User: "mikasa",
		Pass: "mikasa",
		Name: "titan",
		Master: config.DBNodeGroup{
			MaxOpenConnection: 2,
			MaxIdleConnection: 1,
			MaxLifeTime:       21600,
			Nodes: []config.DBNode{
				{
					Host: "192.168.75.230",
					Port: "5432",
				},
			},
		},
		Slave: config.DBNodeGroup{
			MaxOpenConnection: 2,
			MaxIdleConnection: 1,
			MaxLifeTime:       21600,
			Nodes: []config.DBNode{
				{
					Host: "192.168.75.231",
					Port: "5432",
				},
				{
					Host: "192.168.75.232",
					Port: "5432",
				},
				{
					Host: "192.168.75.233",
					Port: "5432",
				},
			},
		},
	}

	err := InitMssqlx(dbConfig)
	if err != nil {
		t.FailNow()
		t.Logf("failed to init DB by mssqlx with err: %+v\n", err)
	}
}
