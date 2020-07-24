package services

import (
	"FestivalSchedule/handler"
	"FestivalSchedule/model"
	"log"
	"reflect"
	"testing"
)

const (
	memberCSV         = "../csv/member.csv"
	bandCSV           = "../csv/band.csv"
	scheduleCSV       = "../csv/schedule.csv"
	impossibleTimeCSV = "../csv/impossible.csv"
)

func TestDefineSchedules(t *testing.T) {
	members, err := handler.ImportMember(memberCSV)
	if err != nil {
		log.Fatal(err)
	}
	locations := model.InitializeLocation()
	bands, err := handler.ImportBand(bandCSV, members, locations)
	if err != nil {
		log.Fatal(err)
	}
	schedules, _ := handler.ImportSchedule(scheduleCSV, locations)
	bands, err = handler.ImportImpossibleTime(impossibleTimeCSV, bands, schedules)
	if err != nil {
		log.Fatal(err)
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
				schedules: schedules,
			},
			want: schedules,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DefineSchedules(tt.args.schedules, tt.args.bands, tt.args.members, tt.args.locations)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefineSchedules() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefineSchedules() = %v, want %v", got, tt.want)
			}
		})
	}
}
