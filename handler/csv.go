package handler

import (
	"FestivalSchedule/model"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	dateFormat = "2006-01-02"
	timeFormat = "3:04"
)

type bandCSVFormat struct {
	id         int
	name       string
	memberNum  int
	member1ID  int
	member2ID  int
	member3ID  int
	member4ID  int
	member5ID  int
	member6ID  int
	member7ID  int
	locID      int
	bandTypeID int
}

type memberCSVFormat struct {
	ID   int
	name string
}

type scheduleCSVFormat struct {
	date  time.Time
	start time.Time
	end   time.Time
}

type impossibleCSVFormat struct {
	groupID int
	date    time.Time
	start   time.Time
	end     time.Time
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func applyMemberSliceToStruct(record []string) (memberCSVFormat, error) {
	var memberStruct memberCSVFormat
	memberStruct.name = record[0]
	if record[1] == "" {
		return memberStruct, errors.New("idが入力されていません")
	}

	id, err := strconv.Atoi(record[1])
	if err != nil {
		return memberStruct, err
	}
	memberStruct.ID = id

	return memberStruct, nil
}

func applyImpossibleTimeSliceToStruct(record []string) (impossibleCSVFormat, error) {
	var impossibleStruct impossibleCSVFormat

	groupID, err := strconv.Atoi(record[0])
	if err != nil {
		log.Print(err)
		return impossibleStruct, err
	}

	impossibleStruct.groupID = groupID

	date, err := time.Parse(dateFormat, record[1])
	if err != nil {
		log.Print(err)
		return impossibleStruct, err
	}

	impossibleStruct.date = date

	startDuration, err := time.ParseDuration(record[2])
	if err != nil {
		log.Print(err)
		return impossibleStruct, err
	}
	impossibleStruct.start = date.Add(startDuration)

	endDuration, err := time.ParseDuration(record[3])
	if err != nil {
		log.Print(err)
		return impossibleStruct, err
	}
	impossibleStruct.end = date.Add(endDuration)

	return impossibleStruct, nil
}

// 情報がないmemberXは0になる
func applyBandSliceToStruct(record []string) (bandCSVFormat, error) {
	var bandStruct bandCSVFormat
	var memberNum int

	id, err := strconv.Atoi(record[0])
	if err != nil {
		log.Fatal(err)
		return bandStruct, err
	}
	bandStruct.id = id

	bandStruct.name = record[1]

	if record[2] != "" {
		member1, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member1ID = member1
		memberNum++
	}

	if record[3] != "" {
		member2, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member2ID = member2
		memberNum++
	}

	if record[4] != "" {
		member3, err := strconv.Atoi(record[4])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member3ID = member3
		memberNum++
	}

	if record[5] != "" {
		member4, err := strconv.Atoi(record[5])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member4ID = member4
		memberNum++
	}

	if record[6] != "" {
		member5, err := strconv.Atoi(record[6])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member5ID = member5
		memberNum++
	}

	if record[7] != "" {
		member6, err := strconv.Atoi(record[7])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member6ID = member6
		memberNum++
	}

	if record[8] != "" {
		member7, err := strconv.Atoi(record[8])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member7ID = member7
		memberNum++
	}

	bandStruct.memberNum = memberNum

	locID, err := strconv.Atoi(record[9])
	if err != nil {
		log.Fatal(err)
		return bandStruct, err
	}
	bandStruct.locID = locID

	bandTypeID, err := strconv.Atoi(record[10])
	if err != nil {
		log.Fatal(err)
		return bandStruct, err
	}
	bandStruct.bandTypeID = bandTypeID

	return bandStruct, nil
}

func applyScheduleSliceToStruct(record []string) (scheduleCSVFormat, error) {
	scheduleStruct := scheduleCSVFormat{}

	date, err := time.Parse(dateFormat, record[0])
	if err != nil {
		log.Print(err)
		return scheduleStruct, err
	}

	scheduleStruct.date = date

	startDuration, err := time.ParseDuration(record[1])
	if err != nil {
		log.Print(err)
		return scheduleStruct, err
	}
	scheduleStruct.start = date.Add(startDuration)

	endDuration, err := time.ParseDuration(record[2])
	if err != nil {
		log.Print(err)
		return scheduleStruct, err
	}
	scheduleStruct.end = date.Add(endDuration)

	return scheduleStruct, nil
}

func bandToStruct(bandStruct bandCSVFormat, members map[int]model.Member, locations map[int]model.Location) (band model.Band, err error) {
	band.ID = bandStruct.id
	band.Name = bandStruct.name

	bandType, err := model.SetBandType(bandStruct.bandTypeID)
	if err != nil {
		return band, err
	}
	band.BandType = bandType

	err = model.CheckValidLocationID(bandStruct.locID, locations)
	if err != nil {
		return band, err
	}
	band.DesireLocationID = bandStruct.locID

	bandIDs := []int{
		bandStruct.member1ID,
		bandStruct.member2ID,
		bandStruct.member3ID,
		bandStruct.member4ID,
		bandStruct.member5ID,
		bandStruct.member6ID,
	}

	isOK, err := model.CheckValidMembers(bandIDs, members)

	if !isOK {
		return band, fmt.Errorf("登録されていないメンバーが存在します。 バンド名: %s", bandStruct.name)
	}

	bandMembers := []model.Member{}
	for _, id := range bandIDs {
		if id == 0 {
			continue
		}
		bandMembers = append(bandMembers, members[id])
	}

	band.Members = bandMembers

	return band, nil
}

// ImportBand is to import bands from csv
// TODO: write test
func ImportBand(fileName string, members map[int]model.Member, locations map[int]model.Location) (bands []model.Band, err error) {
	if !fileExists(fileName) {
		return bands, fmt.Errorf("存在しないfileNameです")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	csvReader := csv.NewReader(file)
	csvReader.Read()
	for {
		record, err := csvReader.Read()
		if err != nil && err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			break
		}

		bandStruct, err := applyBandSliceToStruct(record)
		if err != nil {
			log.Fatal(err)
			break
		}

		band, err := bandToStruct(bandStruct, members, locations)
		if err != nil {
			log.Fatal(err)
			break
		}

		bands = append(bands, band)
	}

	return bands, nil
}

// ImportMember is to import members from csv
func ImportMember(fileName string) (map[int]model.Member, error) {
	var members map[int]model.Member
	members = make(map[int]model.Member)

	if !fileExists(fileName) {
		return members, fmt.Errorf("存在しないfileNameです")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return members, err
	}

	csvReader := csv.NewReader(file)
	csvReader.Read()
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			log.Println(err)
			return members, err
		}

		memberStruct, err := applyMemberSliceToStruct(record)
		if err != nil {
			log.Println(err)
			return members, err
		}

		member := model.Member{
			ID:   memberStruct.ID,
			Name: memberStruct.name,
		}

		members[member.ID] = member
	}

	return members, nil
}

// ImportSchedule is to import schedule from csv
func ImportSchedule(fileName string, locations map[int]model.Location) ([]model.Schedule, error) {
	schedules := []model.Schedule{}

	if !fileExists(fileName) {
		return schedules, fmt.Errorf("存在しないfileNameです")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return schedules, err
	}

	csvReader := csv.NewReader(file)
	csvReader.Read()
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			log.Println(err)
			return schedules, err
		}

		scheduleStruct, err := applyScheduleSliceToStruct(record)
		if err != nil {
			log.Println(err)
			return schedules, err
		}

		for id := range locations {
			schedule := model.Schedule{
				Day:        scheduleStruct.date,
				Start:      scheduleStruct.start,
				End:        scheduleStruct.end,
				LocationID: id,
			}

			schedules = append(schedules, schedule)
		}
	}

	return schedules, nil
}

// ImportImpossibleTime is to import impossibleTime from csv
func ImportImpossibleTime(fileName string, bands []model.Band, schedules []model.Schedule) ([]model.Band, error) {
	if !fileExists(fileName) {
		return bands, fmt.Errorf("存在しないfileNameです")
	}

	file, err := os.Open(fileName)
	if err != nil {
		return bands, err
	}

	csvReader := csv.NewReader(file)
	csvReader.Read()
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil && err != io.EOF {
			log.Println(err)
			return bands, err
		}

		impossibleTimeStruct, err := applyImpossibleTimeSliceToStruct(record)
		if err != nil {
			log.Println(err)
			return bands, err
		}

		impossibleTime := model.ImpossibleTime{
			Date:  impossibleTimeStruct.date,
			Start: impossibleTimeStruct.start,
			End:   impossibleTimeStruct.end,
		}

		bands, err = model.AddImpossibleTimeFromID(impossibleTimeStruct.groupID, bands, impossibleTime)
		if err != nil {
			log.Println(err)
			return bands, err
		}
	}

	return bands, nil
}
