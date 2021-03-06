package services

import (
	"FestivalSchedule/model"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

const (
	memberCSV         = "../csv/member.csv"
	bandCSV           = "../csv/band.csv"
	scheduleCSV       = "../csv/schedule.csv"
	impossibleTimeCSV = "../csv/impossible.csv"
)

func TestDefineSchedules(t *testing.T) {
	locations := model.InitializeLocation()

	schedules := []model.Schedule{
		{
			Day:        time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC),
			Start:      time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
			End:        time.Date(2020, 5, 20, 18, 0, 0, 0, time.UTC),
			LocationID: 1,
		},
	}

	sampleMember := model.Member{
		ID: 1,
	}

	bandType, _ := model.SetBandType(1)

	wantSchedule := []model.Schedule{
		{
			Day:        time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC),
			Start:      time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
			End:        time.Date(2020, 5, 20, 18, 0, 0, 0, time.UTC),
			LocationID: 1,
			Events: []model.Event{
				{
					Start:  time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
					End:    time.Date(2020, 5, 20, 10, 5, 0, 0, time.UTC),
					BandID: 1,
				},
			},
		},
	}

	type args struct {
		schedules []model.Schedule
		bands     []model.Band
		members   map[int]model.Member
		locations map[int]model.Location
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Schedule
		wantErr bool
	}{
		{
			name: "sample-A",
			args: args{
				locations: locations,
				bands: []model.Band{
					{
						ID:   1,
						Name: "",
						Members: []model.Member{
							sampleMember,
						},
						// 希望出演場所
						DesireLocationID: 1,
						BandType:         bandType,
						WantPrayTime: map[int]int{
							1: 10,
							2: 5,
						},
					},
				},
				schedules: schedules,
			},
			want: wantSchedule,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefineSchedules(tt.args.schedules, tt.args.bands, tt.args.members, tt.args.locations)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefineSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DefineSchedules differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_existUnplayBand(t *testing.T) {
	type args struct {
		bands []model.Band
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "true",
			args: args{
				bands: []model.Band{
					{
						ID:       1,
						IsMapped: true,
					},
					{
						ID:       2,
						IsMapped: false,
					},
				},
			},
			want: true,
		},
		{
			name: "false",
			args: args{
				bands: []model.Band{
					{
						ID:       1,
						IsMapped: true,
					},
					{
						ID:       2,
						IsMapped: true,
					},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := existUnplayBand(tt.args.bands); got != tt.want {
				t.Errorf("existUnplayBand() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isTryAllOrders(t *testing.T) {
	bands := []model.Band{
		{
			ID: 1,
		},
		{
			ID: 2,
		},
		{
			ID: 3,
		},
		{
			ID: 4,
		},
		{
			ID: 5,
		},
	}

	type args struct {
		impossibleBandOrders []model.ImpossibleBandOrder
		bands                []model.Band
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "false",
			args: args{
				impossibleBandOrders: []model.ImpossibleBandOrder{{1, 2, 3, 4}},
				bands:                bands,
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				impossibleBandOrders: []model.ImpossibleBandOrder{
					{1, 2, 3, 4},
					{2},
					{5, 4, 3},
					{5},
				},
				bands: bands,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTryAllOrders(tt.args.impossibleBandOrders, tt.args.bands); got != tt.want {
				t.Errorf("isTryAllOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_existEmptySchedule(t *testing.T) {
	locations := model.InitializeLocation()

	emptySchedule := model.Schedule{
		End: time.Date(2020, 5, 20, 18, 0, 0, 0, time.UTC),
		Events: []model.Event{
			{
				End: time.Date(2020, 5, 20, 17, 0, 0, 0, time.UTC),
			},
		},
		LocationID: 1,
	}

	notEmptySchedule := model.Schedule{
		End: time.Date(2020, 5, 20, 18, 0, 0, 0, time.UTC),
		Events: []model.Event{
			{
				End: time.Date(2020, 5, 20, 17, 56, 0, 0, time.UTC),
			},
		},
		LocationID: 1,
	}

	type args struct {
		schedules []model.Schedule
		locations map[int]model.Location
	}
	tests := []struct {
		name    string
		args    args
		want    model.Schedule
		wantErr bool
	}{
		{
			name: "empty schedule is exist",
			args: args{
				schedules: []model.Schedule{
					emptySchedule,
				},
				locations: locations,
			},
			want: emptySchedule,
		},
		{
			name: "empty schedule is not exist",
			args: args{
				schedules: []model.Schedule{
					notEmptySchedule,
				},
				locations: locations,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := existEmptySchedule(tt.args.schedules, tt.args.locations)
			if (err != nil) != tt.wantErr {
				t.Errorf("existEmptySchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("existEmptySchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getAnotherSchedule(t *testing.T) {

	aSchedule := model.Schedule{
		Day:        time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC),
		LocationID: 1,
	}

	bSchedule := model.Schedule{
		Day:        time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC),
		LocationID: 2,
	}

	type args struct {
		schedules      []model.Schedule
		targetSchedule model.Schedule
	}
	tests := []struct {
		name    string
		args    args
		want    model.Schedule
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				schedules: []model.Schedule{
					aSchedule, bSchedule,
				},
				targetSchedule: aSchedule,
			},
			want: bSchedule,
		},
		{
			name: "ok",
			args: args{
				schedules: []model.Schedule{
					aSchedule, bSchedule,
				},
				targetSchedule: bSchedule,
			},
			want: aSchedule,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getAnotherSchedule(tt.args.schedules, tt.args.targetSchedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("getAnotherSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getAnotherSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getUNRegisterdSchedule(t *testing.T) {
	aSchedule := model.Schedule{
		Day:        time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC),
		LocationID: 1,
		Events: []model.Event{
			{
				End: time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
			},
		},
	}

	type args struct {
		schedule model.Schedule
	}
	tests := []struct {
		name    string
		args    args
		want    time.Time
		wantErr bool
	}{
		{
			name: "true",
			args: args{
				schedule: aSchedule,
			},
			want: time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getUNRegisterdSchedule(tt.args.schedule)
			if (err != nil) != tt.wantErr {
				t.Errorf("getUNRegisterdSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getUNRegisterdSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addEvent(t *testing.T) {
	member := model.Member{
		ID: 1,
	}

	band := model.Band{
		ID: 1,
		Members: []model.Member{
			member,
		},
	}
	playTime := 10
	targetTime := time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC)
	events := []model.Event{}

	want := []model.Event{
		{
			Start:      time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
			End:        time.Date(2020, 5, 20, 11, 05, 0, 0, time.UTC),
			BandID:     1,
			LocationID: 1,
		},
	}

	wantBand := model.Band{
		ID: 1,
		Members: []model.Member{
			{
				ID: 1,
				Events: []model.Event{
					{
						Start:      time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
						End:        time.Date(2020, 5, 20, 11, 05, 0, 0, time.UTC),
						BandID:     1,
						LocationID: 1,
					},
				},
			},
		},
	}

	type args struct {
		events     []model.Event
		playTime   int
		band       model.Band
		targetTime time.Time
		locationID int
	}
	tests := []struct {
		name     string
		args     args
		want     []model.Event
		wantBand model.Band
		wantErr  bool
	}{
		{
			name: "ok",
			args: args{
				band:       band,
				playTime:   playTime,
				targetTime: targetTime,
				events:     events,
				locationID: 1,
			},
			want:     want,
			wantBand: wantBand,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotBand, err := addEvent(tt.args.events, tt.args.locationID, tt.args.playTime, tt.args.band, tt.args.targetTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("addEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("addEvent differs: (-got +want)\n%s", diff)
			}
			if diff := cmp.Diff(gotBand, tt.wantBand); diff != "" {
				t.Errorf("addEvent gotBand differs: (-got +want)\n%s", diff)
			}
		})
	}
}

func Test_addTimeForCodeSetting(t *testing.T) {
	schedule := model.Schedule{
		TimeFromBeforeCodeRollUP: 80,
		Events: []model.Event{
			{
				Start: time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
				End:   time.Date(2020, 5, 20, 11, 00, 0, 0, time.UTC),
			},
		},
	}

	playTime := 10

	want := model.Schedule{
		TimeFromBeforeCodeRollUP: 0,
		Events: []model.Event{
			{
				Start: time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
				End:   time.Date(2020, 5, 20, 11, 00, 0, 0, time.UTC),
			},
			{
				Start:         time.Date(2020, 5, 20, 11, 00, 0, 0, time.UTC),
				End:           time.Date(2020, 5, 20, 11, 10, 0, 0, time.UTC),
				IsCodeSetting: true,
			},
		},
	}

	type args struct {
		schedule model.Schedule
		playTime int
	}
	tests := []struct {
		name    string
		args    args
		want    model.Schedule
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				schedule: schedule,
				playTime: playTime,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := addTimeForCodeSetting(tt.args.schedule, tt.args.playTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("addTimeForCodeSetting() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("addTimeForCodeSetting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatchedBand(t *testing.T) {
	type args struct {
		targetTime           time.Time
		locationID           int
		anotherEvents        []model.Event
		bands                []model.Band
		members              map[int]model.Member
		locations            map[int]model.Location
		currentBandOrder     model.ImpossibleBandOrder
		impossibleBandOrders []model.ImpossibleBandOrder
	}
	tests := []struct {
		name    string
		args    args
		want    model.Band
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := searchMatchedBand(tt.args.targetTime, tt.args.locationID, tt.args.anotherEvents, tt.args.bands, tt.args.members, tt.args.locations, tt.args.currentBandOrder, tt.args.impossibleBandOrders)
			if (err != nil) != tt.wantErr {
				t.Errorf("searchMatchedBand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchMatchedBand() = %v, want %v", got, tt.want)
			}
		})
	}
}
