package model

import "errors"

type Location struct {
	id   int
	name string
	// minute
	playTimes      []int
	isContinuePlay bool
	changeTime     int
}

func InitializeLocation() map[int]Location {
	cafe := Location{
		id:             1,
		name:           "cafe",
		playTimes:      []int{5, 10},
		isContinuePlay: false,
	}

	street := Location{
		id:             2,
		name:           "street",
		playTimes:      []int{10, 15},
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
