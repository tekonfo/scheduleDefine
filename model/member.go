package model

import "time"

// Member は名前と所属しているバンドを持っている構造体
type Member struct {
	ID     int
	Name   string
	Bands  []Band
	Events []Event
}

// CheckDoubleCheck はあるメンバーが同時に出演していないかをチェックする関数
func (member Member) CheckDoubleCheck() bool {

	return true
}

// IsViolateTimeRule はメンバーが引数の時間に他の場所の関係による制約に引っかかっていないのかをチェックする関数
func (member Member) IsViolateTimeRule(targetTime time.Time) bool {
	// 同時刻にバンドメンバーが別の場所で歌っていない
	if false {
		return false
	}

	// 移動時間の制約に引っかからない
	if false {
		return false
	}

	return true
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
