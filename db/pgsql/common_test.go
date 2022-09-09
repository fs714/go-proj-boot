package pgsql

import (
	"strings"
	"testing"
	"time"
)

func TestGetStructFields(t *testing.T) {
	p := struct {
		Name string    `db:"name"`
		Age  int       `db:"age"`
		Time time.Time `db:"time"`
	}{}

	fields := GetStructFields(p)
	t.Log(strings.Join(fields, ", "))
}
