package services

import (
	"FestivalSchedule/model"
	"errors"
	"fmt"
	"time"
)

// RollbackError はロールバックが発生される場合に返されるエラーです。
type RollbackError struct {
	Message string
	Code    int
}

func (err *RollbackError) Error() string {
	return fmt.Sprintf("rollback error= %s [code=%d]", err.Message, err.Code)
}

func getUNRegisterdSchedule(schedule model.Schedule) (time.Time, error) {
	return time.Now(), nil
}

func searchMatchedBand(
	targetTime time.Time,
	anotherEvents []model.Event,
	bands []model.Band,
	members map[int]model.Member,
	locations map[int]model.Location,
	currentBandOrder model.ImpossibleBandOrder,
	impossibleBandOrders []model.ImpossibleBandOrder,
) (model.Band, error) {
	for _, band := range bands {
		// 一回も歌っていないバンドである
		if band.IsPlay() {
			continue
		}

		// impossibleBandOrderに含まれていない
		if currentBandOrder.IsExistBandOrder(band, impossibleBandOrders) {
			continue
		}

		// 不可能時間でない
		if band.IsImpossibleTime() {
			continue
		}

		for _, member := range band.Members {
			// 時間の制約チェック
			if member.IsViolateTimeRule(targetTime) {
				continue
			}
		}

		// 条件達成
		return band, nil
	}
	return bands[0], errors.New("could not find band")
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

// defineSchedule　の実行でscheduleが一つ or 一つ + コード巻き取り が埋まる
func defineSchedule(
	schedule *model.Schedule,
	otherSchedule *model.Schedule,
	bands []model.Band,
	members map[int]model.Member,
	locations map[int]model.Location,
	currentBandOrder model.ImpossibleBandOrder,
	impossibleBandOrders []model.ImpossibleBandOrder,
) error {

	//  未登録のスケジュールを取得
	targetTime, err := getUNRegisterdSchedule(*schedule)
	if err != nil {
		return err
	}

	//  当てはまるバンド検索
	targetBand, err := searchMatchedBand(targetTime, otherSchedule.Events, bands, members, locations, currentBandOrder, impossibleBandOrders)
	if err != nil {
		// rollback
		return errors.New("please rollback")
	}

	// scheduleにevent追加
	err = addEvent(&schedule.Events, targetBand, targetTime)
	if err != nil {

	}

	// コード巻き取り時間を追加
	addTimeForCodeSetting(*schedule, targetBand, locations)

	// 対象bandのisMapped追加
	targetBand.IsMapped = true

	// コード巻き取りが必要
	if isNeedCodeSetting(*schedule, locations) {
		err = addCodeSetup(&schedule.Events)
		if err != nil {

		}

		clearTimeForCodeSetup(schedule)
	}

	return nil
}

func existEmptySchedule(schedules []model.Schedule) (bool, model.Schedule) {
	return true, schedules[0]
}

func existUnplayBand() bool {
	return true
}

func isTryAllOrders() bool {
	return true
}

func getAnotherSchedule(schedules []model.Schedule, targetSchedule model.Schedule) model.Schedule {
	return schedules[1]
}

// DefineSchedules は全てのスケジュールを決定するビジネスロジックです。
func DefineSchedules(schedules []model.Schedule, bands []model.Band, members map[int]model.Member, locations map[int]model.Location) ([]model.Schedule, error) {
	impossibleBandOrders := []model.ImpossibleBandOrder{}
	currentBandOrder := model.ImpossibleBandOrder{}
	for {
		if existUnplayBand() {
			fmt.Println("success!!")
			return schedules, nil
		}

		// 順序をすべて試した
		if isTryAllOrders() {
			fmt.Println("fail: could not find match pattern")
			return schedules, errors.New("could not find match pattern")
		}

		// 未登録のスケジュールが存在する
		if b, emptySchedule := existEmptySchedule(schedules); b {
			anotherSchedule := getAnotherSchedule(schedules, emptySchedule)
			err := defineSchedule(&emptySchedule, &anotherSchedule, bands, members, locations, currentBandOrder, impossibleBandOrders)
			if err != nil {
				switch e := err.(type) {
				case *RollbackError:
					currentBandOrder, err = currentBandOrder.DeleteBandOrder()
					if err != nil {
						return schedules, err
					}
					currentBandOrder.AddImpossibleBandOrders(impossibleBandOrders)
					continue
				default:
					return schedules, e
				}
			}
		}
	}
}
