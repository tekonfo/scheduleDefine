package model

type Member struct {
	name  string
	bands []Band
}

// CheckDoubleCheck はあるメンバーが同時に出演していないかをチェックする関数
func (member Member) CheckDoubleCheck() bool {

	return true
}
