package model

// ImpossibleBandOrder is 実現不可能なバンドの順番を記録するためのtype
// ロールバックされる際に追加される
// この配列は1日目~最終日まで継続して全ての場所のバンド情報を保持する
type ImpossibleBandOrder []int

// IsExistBandOrder は現在のBandOrderが不可能順に含まれているのかをチェックする関数
func (currentBandOrder ImpossibleBandOrder) IsExistBandOrder(band Band, impossibleBandOrder []ImpossibleBandOrder) bool {
	return true
}
