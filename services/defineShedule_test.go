package services

import (
	"FestivalSchedule/model"
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
					Start: time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
					End:   time.Date(2020, 5, 20, 10, 5, 0, 0, time.UTC),
					Band: model.Band{
						ID:   1,
						Name: "",
						Members: []model.Member{
							sampleMember,
						},
						// 希望出演場所
						DesireLocationID: 1,
						BandType:         bandType,
						WantCafePlayTime: 10,
						WantStPlayTime:   5,
					},
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
						WantCafePlayTime: 10,
						WantStPlayTime:   5,
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
				t.Errorf("Hogefunc differs: (-got +want)\n%s", diff)
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("DefineSchedules() = %v, want %v", got, tt.want)
			// }
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
