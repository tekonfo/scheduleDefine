package main

import (
	"FestivalSchedule/handler"
	"FestivalSchedule/model"
	"FestivalSchedule/services"
	"fmt"
	"log"
)

const (
	memberCSV         = "../../csv/member.csv"
	bandCSV           = "../../csv/band.csv"
	scheduleCSV       = "../../csv/schedule.csv"
	impossibleTimeCSV = "../../csv/impossible.csv"
)

func tmp(s *string) error {
	*s = *s + " tmp"
	return nil
}

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

	schedules, _ := handler.ImportSchedule(scheduleCSV)

	bands, err = handler.ImportImpossibleTime(impossibleTimeCSV, bands, schedules)
	if err != nil {
		log.Fatal(err)
	}

	schedules, err = services.DefineSchedules(schedules, bands, members, locations)
	fmt.Println(schedules)

	hoge := []string{"aaa", "bbb", "ccc"}

	for i := 0; i < len(hoge); i++ {
		tmp(&hoge[i])
		fmt.Println(hoge[i])
	}

}
