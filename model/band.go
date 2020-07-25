package model

import (
	"errors"
	"time"
)

const MAINBAND = 1
const TMPBAND = 2
const OBBAND = 3

type Band struct {
	ID              int
	Name            string
	Email           string
	Members         []Member
	ImpossibleTimes []ImpossibleTime
	// 希望出演場所
	DesireLocationID int
	BandType         BandType
	IsMultiPlay      bool
	// TODO: これらの情報CSVに追加しないといけない
	WantPrayTime                     map[int]int
	IsHavingMemberAttendingMainStage bool
	IsWantPracticeWithMachine        bool
	IsMapped                         bool
}

type ImpossibleTime struct {
	Date  time.Time
	Start time.Time
	End   time.Time
}

// バンドの種類は、本バンド、企画バンド、OBバンドの三種類が存在する
type BandType struct {
	Id   int
	Name string
}

// IsPlay は対象のバンドが既に歌っているのかどうかをチェックする関数
func (band Band) IsPlay() bool {
	return band.IsMapped
}

// IsImpossibleTime は対象の時間がバンドの不可能時間に一致していないのかチェックする関数
func (band Band) IsImpossibleTime() bool {
	return true
}

func (band Band) IsMatchLocation(locationID int) bool {
	return band.DesireLocationID == locationID
}

func SetBandType(num int) (BandType, error) {
	var bandTp BandType

	switch num {
	case MAINBAND:
		bandTp = BandType{1, "本バンド"}
	case TMPBAND:
		bandTp = BandType{2, "企画バンド"}
	case OBBAND:
		bandTp = BandType{3, "OBバンド"}
	default:
		return bandTp, errors.New("this num is not registered")
	}

	return bandTp, nil
}

// AddImpossibleTimeFromID is function to add impossible time to band which has the args's ID
// TODO:　値をコピーしまくって最適でない。ポインタを使う
func AddImpossibleTimeFromID(ID int, bands []Band, impossibleTime ImpossibleTime) ([]Band, error) {

	for i, band := range bands {
		if band.ID == ID {
			band.ImpossibleTimes = append(band.ImpossibleTimes, impossibleTime)
			bands[i] = band
			return bands, nil
		}
	}

	return bands, errors.New("no band which has such ID")
}

// AddBandIsMapped は名前が一致しているバンドに対してのマップをtrueにする
func (band Band) AddBandIsMapped() (Band, error) {
	band.IsMapped = true
	return band, nil
}
