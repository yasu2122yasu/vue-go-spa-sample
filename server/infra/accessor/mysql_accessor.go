package accessor

import (
	"database/sql"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlAccessor struct {
	DBBase
}

func (a MysqlAccessor) atTransactional(dbmap *gorp.DbMap) (*gorp.Transaction, error) {
	return dbmap.Begin()
}

func (a MysqlAccessor) intiDb() *gorp.DbMap {
	a.dbType = "mysql"
	a.user = "root"
	a.pass = "root"
	a.protocol = "tcp(db:3306)"
	a.dbName = "vgs"

	CONNECT := a.user + ":" + a.pass + "@" + a.pass + "@" + a.protocol + "/"
	CONNECTDB := CONNECT + a.dbName + "?parseTime=true"

	db, err := sql.Open("mysql", CONNECT)
	checkError(err, "sql.Open failed")

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	_, err = dbmap.Exec("CREATE DATABASE IF NOT EXISTS vgs DEFAULT CHARACTER SET utf8;")
	checkError(err, "create db failed")

	db, err = sql.Open("mysql", CONNECTDB)
	checkError(err, "sql.Open failed")

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

}

//多重代入
//引数の中はSQL文を記述している
