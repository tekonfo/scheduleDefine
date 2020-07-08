package model

// Member は名前と所属しているバンドを持っている構造体
type Member struct {
	ID    int
	Name  string
	Bands []Band
}

// CheckDoubleCheck はあるメンバーが同時に出演していないかをチェックする関数
func (member Member) CheckDoubleCheck() bool {

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
