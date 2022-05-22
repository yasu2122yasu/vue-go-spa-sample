package accessor

import (
	"fmt"
	"log"
	"time"

	"github.com/go-gorp/gorp"
)

type IDBAccessor interface {
	initDb() *gorp.DbMap
	connectDb() *gorp.DbMap
	atTransactional(*gorp.DbMap) (*gorp.Transaction, error)
}

func AccessDB(acsr IDBAccessor) *gorp.DbMap {
	return acsr.initDb()
}
func ConnectDb(acsr IDBAccessor) *gorp.DbMap {
	return acsr.connectDb()
}

func StartTransaction(acsr IDBAccessor, dbmap *gorp.DbMap) (*gorp.Transaction, error) {
	return scsr.atTransactional(dbmap)
}

type DBBase struct {
	dbType   string
	user     string
	pass     string
	protocol string
	dbName   string
}

//共通エラーハンドリング
func checkError(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
		fmt.Println(err)
	}
}

type TableBase struct {
	Created time.Time `json:"created" db:"created,notnull"`
	Updated time.Time `json:"updated" db:"updated,notnull"`
}

func (tb *TableBase) PreInsert(s gorp.SqlExecutor) error {
	tb.Created = time.Now()
	tb.Updated = tb.Created
	return nil
}

func (tb *TableBase) PreUpdate(s gorp.SqlExecutor) error {
	tb.Updated = time.Now()
	return nil
}

type Hoge struct {
	ID   int64  `json:"id" db:"id,primaryKey,"`
	Name string `json:"name" db:"name,notnull,size:200"`
	TableBase
}

type Message struct {
	ID   int64  `json:"id" db:"id,primarykey,autoincrement`
	Text string `json:"text" db:"text,notnull,size:200"`
	TableBase
}

//アスタリスクで型名を作るとポインタ型となる。
//アスタリスクは、ポインタ型の宣言をするときに利用する。
