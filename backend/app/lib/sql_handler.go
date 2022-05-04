//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ sql_handler.go [Ver.2022_05_04] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_04
//┃
//┃──┨ SQL管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_04

package lib

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//===================================================|0
//                    構造体定義
//==========================================|2022_05_04

type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

//===================================================|0
//                    SQLHandler新規作成
//==========================================|2022_05_04

func NewSQLHandler() *SQLHandler {
	//------------------------------
	// DB接続定義
	dsn := GenerateDsn()
	var db *SQLHandler
	db.DB, db.Err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if db.Err != nil {
		log.Fatalln(db.Err)
	}

	//------------------------------
	// DB定義
	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatalln(err)
	}

	//------------------------------
	// DB各種設定
	//コネクションプールの最大接続数を設定。
	sqlDB.SetMaxIdleConns(100)
	//接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
	sqlDB.SetMaxOpenConns(100)
	//接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
	sqlDB.SetConnMaxLifetime(100 * time.Second)

	//------------------------------
	// SQLHandler新規作成
	sqlHandler := new(SQLHandler)
	db.DB.Logger.LogMode(4)
	sqlHandler.DB = db.DB

	return sqlHandler
}

//===================================================|0
//                    DSN生成
//==========================================|2022_05_04

func GenerateDsn() (dsn string) {
	//------------------------------
	// DB接続用変数定義
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbName := os.Getenv("MYSQL_DATABASE")
	tz := os.Getenv("TZ")

	//------------------------------
	// DB接続定義
	dsn = fmt.Sprintf(
		"%[1]s:%[2]s@tcp(%[3]s:%[4]s)/%[5]s?parseTime=true&loc=%[6]s",
		user, password, host, port, dbName, tz,
	)

	return dsn
}

//===================================================|0
//                    DB接続
//==========================================|2022_05_04

func DBOpen() {
	dbConn = NewSQLHandler()
}

//===================================================|0
//                    DB切断
//==========================================|2022_05_04

func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

//===================================================|0
//                    DB接続取得
//==========================================|2022_05_04

func GetDBConn() *SQLHandler {
	return dbConn
}

//===================================================|0
//                    トランザクション開始
//==========================================|2022_05_04

func BeginTransaction() *gorm.DB {
	dbConn.DB = dbConn.DB.Begin()
	return dbConn.DB
}

//===================================================|0
//                    ロールバック
//==========================================|2022_05_04

func RollBack() {
	dbConn.DB.Rollback()
}
