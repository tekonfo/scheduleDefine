package services

import (
	"FestivalSchedule/model"
	"errors"
	"fmt"
	"time"
)

// CODEROLLUPINTERVAL はコード巻きの間隔
// これを超えるとコード巻きをしなければならない
const CODEROLLUPINTERVAL = 90

// RollbackError はロールバックが発生される場合に返されるエラーです。
type RollbackError struct {
	Message string
	Code    int
}

func (err *RollbackError) Error() string {
	return fmt.Sprintf("rollback error= %s [code=%d]", err.Message, err.Code)
}

func getUNRegisterdSchedule(schedule model.Schedule) (time.Time, error) {
	lastEvent := schedule.Events[len(schedule.Events)-1]
	return lastEvent.End, nil
}

func searchMatchedBand(
	targetTime time.Time,
	locationID int,
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

		// 希望しているLocationIDと今回のLocationがあっているか
		if !band.IsMatchLocation(locationID) {
			continue
		}

		// impossibleBandOrderに含まれていない
		if currentBandOrder.IsExistBandOrder(band.ID, impossibleBandOrders) {
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

// bandにIsMapped追加
// eventsにevent追加
func addEvent(events []model.Event, locationID int, playTime int, band *model.Band, targetTime time.Time) ([]model.Event, error) {
	event := model.Event{
		Start:      targetTime,
		End:        targetTime.Add(time.Minute * time.Duration(playTime)),
		Band:       *band,
		LocationID: locationID,
	}
	events = append(events, event)

	for _, member := range band.Members {
		member.Events = append(member.Events, event)
	}

	return events, nil
}

func addTimeForCodeSetting(schedule model.Schedule, playTime int) (model.Schedule, error) {
	schedule.TimeFromBeforeCodeRollUP += playTime

	// コード巻きイベント追加
	if schedule.TimeFromBeforeCodeRollUP >= CODEROLLUPINTERVAL {

		lastEvent := schedule.Events[len(schedule.Events)-1]
		lastEventEnd := lastEvent.End

		event := model.Event{
			Start:         lastEventEnd,
			End:           lastEventEnd.Add(time.Minute * time.Duration(playTime)),
			IsCodeSetting: true,
		}

		schedule.Events = append(schedule.Events, event)

		schedule.TimeFromBeforeCodeRollUP -= CODEROLLUPINTERVAL
	}

	return schedule, nil
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

// defineScheduleの実行でscheduleが一つ or 一つ + コード巻き取り が埋まる
func defineSchedule(
	schedule model.Schedule,
	otherSchedule model.Schedule,
	bands []model.Band,
	members map[int]model.Member,
	locations map[int]model.Location,
	currentBandOrder model.ImpossibleBandOrder,
	impossibleBandOrders []model.ImpossibleBandOrder,
) error {

	//  未登録のスケジュールを取得
	targetTime, err := getUNRegisterdSchedule(schedule)
	if err != nil {
		return err
	}

	//  当てはまるバンド検索
	targetBand, err := searchMatchedBand(targetTime, *&schedule.LocationID, otherSchedule.Events, bands, members, locations, currentBandOrder, impossibleBandOrders)
	if err != nil {
		// rollback
		return errors.New("please rollback")
	}

	playTime := targetBand.WantPrayTime[schedule.LocationID]

	// scheduleにevent追加
	schedule.Events, err = addEvent(schedule.Events, schedule.LocationID, playTime, &targetBand, targetTime)
	if err != nil {
		return err
	}

	_, err = targetBand.AddBandIsMapped()
	if err != nil {
		return err
	}

	// コード巻き取り時間を追加
	// ここで必要ならEventも追加してしまっている
	schedule, err = addTimeForCodeSetting(schedule, playTime)
	if err != nil {
		return err
	}

	// 対象bandのisMapped追加
	targetBand.IsMapped = true

	return nil
}

// 末尾のスケジュールの最後が埋まっているかどうかを確認する
func existEmptySchedule(schedules []model.Schedule, locations map[int]model.Location) (model.Schedule, error) {
	for _, schedule := range schedules {
		if len(schedule.Events) == 0 {
			return schedule, nil
		}

		lastEvent := schedule.Events[len(schedule.Events)-1]

		// スケジュールを埋めれない最小の時間
		targetLocation := locations[schedule.LocationID]
		minimumTime := targetLocation.PlayTimes[0]

		estimateNextEndTime := lastEvent.End.Add(time.Minute * time.Duration(minimumTime))

		if estimateNextEndTime.Unix() <= schedule.End.Unix() {
			return schedule, nil
		}
	}

	return schedules[0], errors.New("could not find empty schedule")
}

func existUnplayBand(bands []model.Band) bool {
	for _, band := range bands {
		if !band.IsMapped {
			return true
		}
	}
	return false
}

// bands配列の、一番末尾のbandのみのbandordersが存在すれば、
// 全パターンを試したことになるため、それで判定
func isTryAllOrders(impossibleBandOrders []model.ImpossibleBandOrder, bands []model.Band) bool {
	lastBand := bands[len(bands)-1]
	targetBandOrder := []int{lastBand.ID}
	for _, order := range impossibleBandOrders {
		if order.Deepequal(targetBandOrder) {
			return true
		}
	}
	return false
}

func getAnotherSchedule(schedules []model.Schedule, targetSchedule model.Schedule) (model.Schedule, error) {
	targetLocationID := targetSchedule.LocationID
	targetDay := targetSchedule.Day

	for _, schedule := range schedules {
		if schedule.Day.Unix() == targetDay.Unix() && schedule.LocationID != targetLocationID {
			return schedule, nil
		}
	}

	return model.Schedule{}, errors.New("no such a schedule")
}

// DefineSchedules は全てのスケジュールを決定するビジネスロジックです。
func DefineSchedules(schedules []model.Schedule, bands []model.Band, members map[int]model.Member, locations map[int]model.Location) ([]model.Schedule, error) {
	impossibleBandOrders := []model.ImpossibleBandOrder{}
	currentBandOrder := model.ImpossibleBandOrder{}
	for {
		if existUnplayBand(bands) {
			fmt.Println("success!!")
			return schedules, nil
		}

		// 順序をすべて試した
		if isTryAllOrders(impossibleBandOrders, bands) {
			fmt.Println("fail: could not find match pattern")
			return schedules, errors.New("could not find match pattern")
		}

		// 未登録のスケジュールが存在する
		emptySchedule, err := existEmptySchedule(schedules, locations)
		if err != nil {
			return schedules, errors.New("all shcedules is mapped")
		}

		// cafeならst, stならcafeのスケジュールを取得
		anotherSchedule, err := getAnotherSchedule(schedules, emptySchedule)
		if err != nil {
			return schedules, err
		}

		// scheduleを一件決定
		err = defineSchedule(emptySchedule, anotherSchedule, bands, members, locations, currentBandOrder, impossibleBandOrders)
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
