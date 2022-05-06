//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ /migration/main.go [Ver.2022_05_04] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_04
//┃
//┃──┨ マイグレーション管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_04

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"fullcalendar/lib"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/source/file"
	// "github.com/joho/godotenv"
	// "github.com/pkg/errors"
)

//===================================================|0
//                    グローバル変数
//==========================================|2022_05_04

var migrationFilePath = "file://migrations/"

//===================================================|0
//                    メイン処理
//==========================================|2022_05_04

func main() {
	fmt.Println("Start Migration")

	//------------------------------
	// コマンドライン引数の取得
	flag.Parse()
	command := flag.Arg(0)
	migrationFileName := flag.Arg(1)
	if command == "" {
		showUsage()
		os.Exit(1)
	}

	//------------------------------
	// マイグレーション情報取得
	m := newMigrate()
	version, dirty, _ := m.Version()
	force := flag.Bool("f", false, "force execute fixed sql")
	if dirty && *force {
		fmt.Println("force=true: force execute current version sql")
		m.Force(int(version))
	}

	//------------------------------
	// 処理の分岐
	switch command {
	case "new":
		newMigration(migrationFileName)
	case "up":
		up(m)
	case "down":
		down(m)
	case "drop":
		drop(m)
	case "version":
		showVersionInfo(m.Version())
	default:
		fmt.Println("\nerror: invalid command '", command, "'")
		showUsage()
		os.Exit(0)
	}
}

//===================================================|0
//                    マイグレーション新規実行
//==========================================|2022_05_04

func newMigrate() *migrate.Migrate {
	//------------------------------
	// DB接続
	dsn := lib.GenerateDsn()
	driverName := os.Getenv("DB_DRIVER")
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	//------------------------------
	// Driver取得
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	fileDriver, err := (&file.File{}).Open(migrationFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	//------------------------------
	// DB作成
	m, err := migrate.NewWithInstance(
		"file",
		fileDriver,
		driverName,
		driver,
	)
	if err != nil {
		log.Fatalln(err)
	}

	return m
}

//===================================================|0
//                    マイグレーション新規作成
//==========================================|2022_05_04

func newMigration(name string) {
	if name == "" {
		fmt.Println("\nerror: migration file name must be supplied as an argument")
		os.Exit(1)
	}

	//------------------------------
	// マイグレーションファイル生成
	base := fmt.Sprintf(
		"./migration/migrations/%[1]s_%[2]s",
		time.Now().Format("20060102030405"), name,
	)
	ext := ".sql"
	lib.CreateFile(base + ".up" + ext)
	lib.CreateFile(base + ".down" + ext)
}

//===================================================|0
//                    UP
//==========================================|2022_05_04

func up(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Up()
	if err != nil {
		if err.Error() != "no change" {
			log.Fatalln(err)
		}
		fmt.Println("\nno change")
	} else {
		fmt.Println("\nUpdated:")
		showVersionInfo(m.Version())
	}
}

//===================================================|0
//                    down
//==========================================|2022_05_04

func down(m *migrate.Migrate) {
	fmt.Println("Before:")
	showVersionInfo(m.Version())
	err := m.Steps(-1)
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("\nUpdated:")
		showVersionInfo(m.Version())
	}
}

//===================================================|0
//                    drop
//==========================================|2022_05_04

func drop(m *migrate.Migrate) {
	err := m.Drop()
	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("Dropped all migrations")
		return
	}
}

//===================================================|0
//                    利用方法表示
//==========================================|2022_05_04

func showUsage() {
	fmt.Println(`
		-------------------------------------
		Usage:
			go run migration/main.go <command>
		Commands:
			new FILENAME  Create new up & down migration files
			up        Apply up migrations
			down      Apply down migrations
			drop      Drop everything
			version   Check current migrate version
		-------------------------------------
	`)
}

//===================================================|0
//                    バージョン表示
//==========================================|2022_05_04

func showVersionInfo(version uint, dirty bool, err error) {
	fmt.Println("-------------------")
	fmt.Println("version : ", version)
	fmt.Println("dirty   : ", dirty)
	fmt.Println("error   : ", err)
	fmt.Println("-------------------")
}
