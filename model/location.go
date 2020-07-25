package model

import "errors"

// Location は歌う場所の情報を定義している構造体
type Location struct {
	ID   int
	name string
	// minute
	PlayTimes      []int
	isContinuePlay bool
}

// InitializeLocation は初期のLocation情報を返す関数
func InitializeLocation() map[int]Location {
	cafe := Location{
		ID:             1,
		name:           "street",
		PlayTimes:      []int{5, 10},
		isContinuePlay: false,
	}

	street := Location{
		ID:             2,
		name:           "cafe",
		PlayTimes:      []int{10, 15},
		isContinuePlay: true,
	}

	return map[int]Location{
		1: cafe,
		2: street,
	}
}

// CheckValidLocationID は数字がlocationIDとして定義されているのかを判定する関数
func CheckValidLocationID(num int, locations map[int]Location) error {
	if _, ok := locations[num]; ok {
		return nil
	}
	return errors.New("this ID is not in locations")
}
