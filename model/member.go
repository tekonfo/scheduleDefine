package model

import (
	"time"
)

// MOVETIME は考慮される移動時間
const MOVETIME = 10

// Member は名前と所属しているバンドを持っている構造体
type Member struct {
	ID      int
	Name    string
	BandIDs []int
	Events  []Event
}

// IsViolateTimeRule はメンバーが引数の時間に他の場所の関係による制約に引っかかっていないのかをチェックする関数
func (member Member) IsViolateTimeRule(startTime time.Time, endTime time.Time) bool {
	for _, event := range member.Events {
		// 同時刻にバンドメンバーが別の場所で歌っている
		if event.Start.Unix() <= startTime.Unix() && startTime.Unix() <= event.End.Unix() {
			return true
		}

		if event.Start.Unix() <= endTime.Unix() && endTime.Unix() <= event.End.Unix() {
			return true
		}

		// 移動時間の制約に引っかからない
		// MOVETIMEを加算したEventの終了時刻
		nextPossibleTime := event.End.Add(time.Minute * time.Duration(MOVETIME))
		if startTime.Unix() <= nextPossibleTime.Unix() && nextPossibleTime.Unix() <= endTime.Unix() {
			return true
		}
	}

	return false
}

// CheckValidMembers は 登録しているMemberの配列の中にidが存在しているのかをチェックする関数
func CheckValidMembers(ids []int, members map[int]Member) (bool, error) {
	for _, id := range ids {
		if id == 0 {
			continue
		}
		if _, ok := members[id]; !ok {
			return false, nil
		}
	}

	return true, nil
}
