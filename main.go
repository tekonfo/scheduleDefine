package main

import (
	"FestivalSchedule/handler"
	"FestivalSchedule/model"
	"fmt"
)

func main() {
	members := map[int]model.Member{
		1: {ID: 1, Name: "taisho"},
		2: {ID: 2, Name: "haruki"},
		3: {ID: 3, Name: "katsuya"},
		4: {ID: 4, Name: "hinako"},
		5: {ID: 5, Name: "rino"},
		6: {ID: 6, Name: "miran"},
	}
	locations := model.InitializeLocation()
	fileName := "csv/test/bandっっっd.csv"
	bands, err := handler.ImportBand(fileName, members, locations)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(bands)
}
