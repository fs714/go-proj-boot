package pgtable

import (
	"reflect"
	"strings"
	"time"

	"github.com/jmoiron/sqlx/reflectx"
)

type TableBaseColumn struct {
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

func GetStructFields(i interface{}, skips []string) []string {
	fields := make([]string, 0)

	skipMap := make(map[string]struct{}, len(skips))
	for _, f := range skips {
		skipMap[f] = struct{}{}
	}

	m := reflectx.NewMapperFunc("db", strings.ToLower)
	fm := m.FieldMap(reflect.ValueOf(i))
	for k := range fm {
		if _, ok := skipMap[k]; ok {
			continue
		}

		fields = append(fields, k)
	}

	return fields
}

func GetSelectColumnsSql(cols []string) string {
	return strings.Join(cols, ", ")
}

func GetInsertColumnsSql(cols []string) string {
	return strings.Join(cols, ", ")
}

func GetInsertNamedValuesSql(cols []string) string {
	s := make([]string, 0, len(cols))
	for _, c := range cols {
		s = append(s, ":"+c)
	}

	return strings.Join(s, ", ")
}

func GetUpdateSetSql(cols []string) string {
	s := make([]string, 0, len(cols))
	for _, c := range cols {
		s = append(s, c+"=:"+c)
	}

	return strings.Join(s, ", ")
}
