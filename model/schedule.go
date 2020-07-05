package model

import "time"

// Schedule は各開催日の出演バンドが記載された時系列のスケジュール
type Schedule struct {
	day    time.Time
	start  time.Time
	end    time.Time
	evnets []event
}

type event struct {
	start         time.Time
	end           time.Time
	band          Band
	isCodeSetting bool
}
