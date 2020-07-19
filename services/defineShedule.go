package services

import (
	"FestivalSchedule/model"
	"time"
)

func getUNRegisterdSchedule(schedule model.Schedule) (time.Time, error) {
	return time.Now(), nil
}

func searchMatchedBand(
	targetTime time.Time,
	anotherEvents []model.Event,
	bands []model.Band,
	members map[int]model.Member,
	locations map[int]model.Location,
) (model.Band, error) {
	return bands[0], nil
}

func getEvents(schedule model.Schedule, isStreet bool) []model.Event {
	if isStreet {
		return schedule.StEvents
	}
	return schedule.CafeEvents
}

func addEvent(events *[]model.Event, band model.Band, targetTime time.Time) error {
	//
	return nil
}

func addTimeForCodeSetting(schedule model.Schedule, band model.Band, locations map[int]model.Location) error {
	return nil
}

func clearTimeForCodeSetup(schedule *model.Schedule) error {
	schedule.TimeFromBeforeCodeRollUP = 0
	return nil
}

func addCodeSetup(evnets *[]model.Event) error {
	return nil
}

func isNeedCodeSetting(schedule model.Schedule, locations map[int]model.Location) bool {
	return true
}

func defineSchedule(schedule *model.Schedule, bands []model.Band, members map[int]model.Member, locations map[int]model.Location) error {
	isStreets := [2]bool{true, false}
	// TODO: ループ処理を追加する必要がある
	for _, isStreet := range isStreets {
		//  未登録のスケジュールを取得
		targetTime, err := getUNRegisterdSchedule(*schedule)
		if err != nil {
			return err
		}

		anotherEvents := getEvents(*schedule, isStreet)

		//  当てはまるバンド検索
		targetBand, err := searchMatchedBand(targetTime, anotherEvents, bands, members, locations)
		if err != nil {

		}

		// scheduleにevent追加
		targetEvents := getEvents(*schedule, !isStreet)
		err = addEvent(&targetEvents, targetBand, targetTime)
		if err != nil {

		}

		// コード巻き取り時間を追加
		addTimeForCodeSetting(*schedule, targetBand, locations)

		// 対象bandのisMapped追加
		targetBand.IsMapped = true

		// コード巻き取りが必要
		if isNeedCodeSetting(*schedule, locations) {
			err = addCodeSetup(&targetEvents)
			if err != nil {

			}

			clearTimeForCodeSetup(schedule)
		}
	}

	return nil
}

// DefineSchedules は全てのスケジュールを決定するビジネスロジックです。
func DefineSchedules(schedules []model.Schedule, bands []model.Band, members map[int]model.Member, locations map[int]model.Location) ([]model.Schedule, error) {
	for i := range schedules {
		err := defineSchedule(&schedules[i], bands, members, locations)
		if err != nil {
			return schedules, err
		}
	}
	return schedules, nil
}
