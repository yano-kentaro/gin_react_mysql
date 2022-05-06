//┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
//┃
//┃──┨ events.go [Ver.2022_05_04] ┃
//┃
//┠──┨ Copyright(C) https://github.com/yano-kentaro
//┠──┨ https://www.kengineer.dev
//┠──┨ 開発開始日：2022_05_04
//┃
//┃──┨ Events管理 ┃
//┃
//┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

//===================================================|0
//                    依存関係
//==========================================|2022_05_04

package events

import (
	"fullcalendar/lib"
	"log"
)

//===================================================|0
//                    構造体定義
//==========================================|2022_05_04

type Event struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Start string `json:"start"`
	End   string `json:"end"`
}

type Events struct {
	Items []Event
}

//===================================================|0
//                    インスタンス定義
//==========================================|2022_05_04

func New() *Events {
	return &Events{}
}

//===================================================|0
//                    指定期間の予定取得
//==========================================|2022_05_04

func (r *Events) Get(start string, end string) (events []Event) {
	db := lib.GetDBConn().DB
	err := db.Where(
		"start >= ? AND end <= ?", start, end,
	).Find(&events).Error
	if err != nil {
		log.Fatalln(err)
	}
	return events
}

//===================================================|0
//                    新規作成
//==========================================|2022_05_04

func (r *Events) Create(c Event) {
	r.Items = append(r.Items, c)
	db := lib.GetDBConn().DB
	err := db.Create(&c).Error
	if err != nil {
		log.Fatalln(err)
	}
}

//===================================================|0
//                    更新
//==========================================|2022_05_05

func (r *Events) Update(id int, u Event) {
	db := lib.GetDBConn().DB
	err := db.Model(&u).Where("id = ?", id).Updates(&u).Error
	if err != nil {
		log.Fatalln(err)
	}
}

//===================================================|0
//                    削除
//==========================================|2022_05_04

func (e Events) Delete(id int) {
	db := lib.GetDBConn().DB
	err := db.Delete(&Event{}, id).Error
	if err != nil {
		log.Fatalln(err)
	}
}
