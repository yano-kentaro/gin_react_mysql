//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ events.go [Ver.2022_05_04] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_04
//┃
//┃──┨ Eventsテーブル関数 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_04

package handler

import (
	"fullcalendar/models/events"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//===================================================|0
//                    構造体定義
//==========================================|2022_05_04

type RequestPostEvent struct {
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

//===================================================|0
//                    GET
//==========================================|2022_05_04

func GetEvents(m *events.Events) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// クエリパラメータ取得
		start := ctx.Query("start")
		end := ctx.Query("end")

		// レスポンス作成
		response := m.Get(start, end)

		// JSON化
		ctx.JSON(http.StatusOK, response)
	}
}

//===================================================|0
//                    POST
//==========================================|2022_05_04

func CreateEvent(m *events.Events) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// リクエストBody取得
		requestBody := RequestPostEvent{}
		ctx.ShouldBindJSON(&requestBody)

		// 予定登録
		item := events.Event{
			Title: requestBody.Title,
			Start: requestBody.Start,
			End:   requestBody.End,
		}
		m.Create(item)

		ctx.Status(http.StatusNoContent)
	}
}

//===================================================|0
//                    PUT
//==========================================|2022_05_04

func UpdateEvent(m *events.Events) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 対象ID取得
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalln(err)
		}
		// リクエストBody取得
		requestBody := RequestPostEvent{}
		ctx.ShouldBindJSON(&requestBody)

		// 予定更新
		item := events.Event{
			Title: requestBody.Title,
			Start: requestBody.Start,
			End:   requestBody.End,
		}
		m.Update(id, item)

		ctx.Status(http.StatusNoContent)
	}
}

//===================================================|0
//                    Delete
//==========================================|2022_05_04

func DeleteEvent(m *events.Events) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 対象ID取得
		idStr := ctx.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Fatalln(err)
		}

		// 予定削除
		m.Delete(id)

		ctx.Status(http.StatusNoContent)
	}
}
