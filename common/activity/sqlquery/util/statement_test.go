package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMySql(t *testing.T) {
	dbType, err := GetDbHelper("mysql")

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ? and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParseOracle(t *testing.T) {
	dbType, err := GetDbHelper("oracle")

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = :foo and s = :bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParsePostgres(t *testing.T) {
	dbType, err := GetDbHelper("postgres")

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ? and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParseSqlLite(t *testing.T) {
	dbType, err := GetDbHelper("sqlite")

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ?", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = ? s = ? and q = \"tar\"", s.PreparedStatementSQL())
}

func TestParseSqlServer(t *testing.T) {
	dbType, err := GetDbHelper("sqlserver")

	sql := "select * from table where t = :foo and s = :bar"
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where t = @foo and s = @bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar"
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = @foo and s = @bar", s.PreparedStatementSQL())

	sql = "select * from table where a = \"ignore :blah\" and t = :foo and s = :bar and q = \"tar\""
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	assert.Equal(t, sql, s.String())
	assert.Equal(t, "select * from table where a = \"ignore :blah\" and t = @foo and s = @bar and q = \"tar\"", s.PreparedStatementSQL())
}

func TestFlattenSql(t *testing.T) {

	sql := "select * from table where t = :foo and s = :bar and r = :other"
	params := map[string]interface{}{"foo": true, "bar": 2, "other": "test"}
	dbType, err := GetDbHelper("mysql")
	s, err := NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	flattenedSQL := s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = true and s = 2 and r = 'test'", flattenedSQL)

	dbType, err = GetDbHelper("oracle")
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = 1 and s = 2 and r = 'test'", flattenedSQL)

	dbType, err = GetDbHelper("postgres")
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = TRUE and s = 2 and r = 'test'", flattenedSQL)

	dbType, err = GetDbHelper("sqlite")
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = 1 and s = 2 and r = 'test'", flattenedSQL)

	dbType, err = GetDbHelper("sqlserver")
	s, err = NewSQLStatement(dbType, sql)
	assert.Nil(t, err)
	flattenedSQL = s.ToStatementSQL(params)
	assert.Equal(t, "select * from table where t = TRUE and s = 2 and r = 'test'", flattenedSQL)
}
