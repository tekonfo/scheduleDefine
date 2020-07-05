package model

type Location struct {
	id   int
	name string
	// minute
	playTimes      []int
	isContinuePlay bool
	changeTime     int
}

func initialize() []Location {
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

	return []Location{cafe, street}
}

// TODO: カフェ→ストの移動時間は5分以上、スト→カフェの移動時間は10分以上になるようにする必要がある
// これどうやって実装しようか？
