package model

import "reflect"

// ImpossibleBandOrder is 実現不可能なバンドの順番を記録するためのtype
// ロールバックされる際に追加される
// この配列は1日目~最終日まで継続して全ての場所のバンド情報を保持する
type ImpossibleBandOrder []int

// IsExistBandOrder は現在のBandOrderが不可能順に含まれているのかをチェックする関数
func (currentBandOrder ImpossibleBandOrder) IsExistBandOrder(band Band, impossibleBandOrder []ImpossibleBandOrder) bool {
	return true
}

// DeleteBandOrder は現在のBandOrderを一つ削除する
// もしスライスがゼロならerrorを返す
// ASK: これは[]intなので、bandOrderを返さなくてもいいのか？
func (currentBandOrder ImpossibleBandOrder) DeleteBandOrder() (ImpossibleBandOrder, error) {
	return currentBandOrder, nil
}

// AddImpossibleBandOrders は現在の順番を追加したImpossiblebandOrdersを返す関数
func (currentBandOrder ImpossibleBandOrder) AddImpossibleBandOrders(impossibleBandOrders []ImpossibleBandOrder) ([]ImpossibleBandOrder, error) {

	return impossibleBandOrders, nil
}

// Deepequal はbandOrderが引数のintsと一致しているのかどうかを判定する関数
func (currentBandOrder ImpossibleBandOrder) Deepequal(ints []int) bool {
	orderInt := []int{}
	for _, order := range currentBandOrder {
		orderInt = append(orderInt, order)
	}
	return reflect.DeepEqual(orderInt, ints)
}
