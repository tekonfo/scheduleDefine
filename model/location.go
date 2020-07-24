package model

import "errors"

type Location struct {
	ID   int
	name string
	// minute
	PlayTimes      []int
	isContinuePlay bool
	changeTime     int
}

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

func CheckValidLocationID(num int, locations map[int]Location) error {
	if _, ok := locations[num]; ok {
		return nil
	}
	return errors.New("this ID is not in locations")
}

// TODO: カフェ→ストの移動時間は5分以上、スト→カフェの移動時間は10分以上になるようにする必要がある
// これどうやって実装しようか？
