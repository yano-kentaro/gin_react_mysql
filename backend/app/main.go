//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ main.go [Ver.2022_05_03] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_03
//┃
//┃──┨ メイン処理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_03

package main

import (
	"fmt"
	"fullcalendar/handler"
	"fullcalendar/lib"
	"fullcalendar/models/events"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//===================================================|0
//                    メイン処理
//==========================================|2022_05_04

func main() {
	//------------------------------
	// DB接続＋終了時切断
	lib.DBOpen()
	defer lib.DBClose()

	//------------------------------
	// APIサーバー設定
	r := gin.Default() // r <- router
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://frontend:80",
		},
		AllowMethods: []string{
			"GET", "POST", "PUT", "DELETE", "OPTIONS",
		},
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           24 * time.Hour,
	}))

	//------------------------------
	// Events
	events := events.New()
	r.GET("/api/events", handler.GetEvents(events))
	r.POST("/api/event/create", handler.CreateEvent(events))
	r.PUT("/api/event/update/:id", handler.UpdateEvent(events))
	r.DELETE("/api/event/delete/:id", handler.DeleteEvent(events))

	//------------------------------
	// APIサーバー起動
	httpHost := os.Getenv("HTTP_HOST")
	httpPort := os.Getenv("HTTP_PORT")
	r.Run(fmt.Sprintf(
		"%[1]s:%[2]s", httpHost, httpPort,
	))
}
