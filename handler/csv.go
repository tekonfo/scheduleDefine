package handler

import (
	"FestivalSchedule/model"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type bandCSVFormat struct {
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

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// 情報がないmemberXは0になる
func applyBandSliceToStruct(record []string) (bandCSVFormat, error) {
	var bandStruct bandCSVFormat
	var memberNum int
	bandStruct.name = record[0]
	if record[1] != "" {
		member1, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member1ID = member1
		memberNum++
	}

	if record[2] != "" {
		member2, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member2ID = member2
		memberNum++
	}

	if record[3] != "" {
		member3, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member3ID = member3
		memberNum++
	}

	if record[4] != "" {
		member4, err := strconv.Atoi(record[4])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member4ID = member4
		memberNum++
	}

	if record[5] != "" {
		member5, err := strconv.Atoi(record[5])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member5ID = member5
		memberNum++
	}

	if record[6] != "" {
		member6, err := strconv.Atoi(record[6])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member6ID = member6
		memberNum++
	}

	if record[7] != "" {
		member7, err := strconv.Atoi(record[7])
		if err != nil {
			log.Fatal(err)
			return bandStruct, err
		}
		bandStruct.member7ID = member7
		memberNum++
	}

	bandStruct.memberNum = memberNum

	locID, err := strconv.Atoi(record[8])
	if err != nil {
		log.Fatal(err)
		return bandStruct, err
	}
	bandStruct.locID = locID

	bandTypeID, err := strconv.Atoi(record[9])
	if err != nil {
		log.Fatal(err)
		return bandStruct, err
	}
	bandStruct.bandTypeID = bandTypeID

	return bandStruct, nil
}

func bandToStruct(bandStruct bandCSVFormat, members map[int]model.Member, locations map[int]model.Location) (band model.Band, err error) {

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
