package main

import (
	"FestivalSchedule/handler"
	"FestivalSchedule/model"
	"fmt"
	"log"
)

const (
	memberCSV = "csv/member.csv"
	bandCSV   = "csv/band.csv"
)

func main() {
	members, err := handler.ImportMember(memberCSV)
	if err != nil {
		log.Fatal(err)
	}

	locations := model.InitializeLocation()

	bands, err := handler.ImportBand(bandCSV, members, locations)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(bands)
}
