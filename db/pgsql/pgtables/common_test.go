package pgtable

import (
	"strings"
	"testing"
)

type TestUser struct {
	Uuid   string
	Name   string `db:"name"`
	Email  string `db:"email"`
	Age    int    `db:"age"`
	Role   string `db:"role"`
	Points int    `db:"points"`
	TableBaseColumn
}

func TestGetStructFields(t *testing.T) {
	fields := GetStructFields(TestUser{}, []string{})
	t.Log(strings.Join(fields, ", "))
}

func TestGetSelectColumnsSql(t *testing.T) {
	fields := GetStructFields(TestUser{}, []string{})
	sql := GetSelectColumnsSql(fields)
	t.Log(sql)
}

func TestGetInsertColumnsSql(t *testing.T) {
	fields := GetStructFields(TestUser{}, []string{})
	sql := GetInsertColumnsSql(fields)
	t.Log(sql)
}

func TestGetInsertNamedValuesSql(t *testing.T) {
	fields := GetStructFields(TestUser{}, []string{})
	sql := GetInsertNamedValuesSql(fields)
	t.Log(sql)
}

func TestGetUpdateSetSql(t *testing.T) {
	fields := GetStructFields(TestUser{}, []string{})
	sql := GetUpdateSetSql(fields)
	t.Log(sql)
}
