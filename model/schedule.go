package model

import "time"

// Schedule は各開催日の出演バンドが記載された時系列のスケジュール
type Schedule struct {
	LocationID               int
	Day                      time.Time
	Start                    time.Time
	End                      time.Time
	Events                   []Event
	TimeFromBeforeCodeRollUP int
}

// Event はそれぞれのスケジュールが登録されている
type Event struct {
	Start         time.Time
	End           time.Time
	LocationID    int
	BandID        int
	IsCodeSetting bool
}
