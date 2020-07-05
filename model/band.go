package model

import "time"

type Band struct {
	members         []Member
	impossibleTimes []ImpossibleTime
	// 希望出演場所
	desireLocationID int
	bandType         int
	isMultiPlay      bool
}

type ImpossibleTime struct {
	day   time.Time
	start time.Time
	end   time.Time
}

// バンドの種類は、本バンド、企画バンド、OBバンドの三種類が存在する
type bandType struct {
	id   int
	name string
}
