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
	BandType         bandType
	IsMultiPlay      bool
	// TODO: これらの情報CSVに追加しないといけない
	wantCafePlayTime                 int
	wantStPlayTime                   int
	isHavingMemberAttendingMainStage bool
	isWantPracticeWithMachine        bool
	IsMapped                         bool
}

type ImpossibleTime struct {
	Date  time.Time
	Start time.Time
	End   time.Time
}

// バンドの種類は、本バンド、企画バンド、OBバンドの三種類が存在する
type bandType struct {
	id   int
	name string
}

// IsPlay は対象のバンドが既に歌っているのかどうかをチェックする関数
func (band Band) IsPlay() bool {
	return band.IsMapped
}

// IsImpossibleTime は対象の時間がバンドの不可能時間に一致していないのかチェックする関数
func (band Band) IsImpossibleTime() bool {
	return true
}

func SetBandType(num int) (bandType, error) {
	var bandTp bandType

	switch num {
	case MAINBAND:
		bandTp = bandType{id: 1, name: "本バンド"}
	case TMPBAND:
		bandTp = bandType{2, "企画バンド"}
	case OBBAND:
		bandTp = bandType{3, "OBバンド"}
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
