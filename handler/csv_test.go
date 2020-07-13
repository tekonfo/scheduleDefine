package handler

import (
	"FestivalSchedule/model"
	"reflect"
	"testing"
	"time"
)

func Test_applyBandSliceToStruct(t *testing.T) {
	okRecord := []string{
		"LiuLiu",
		"1",
		"2",
		"3",
		"4",
		"5",
		"6",
		"",
		"2",
		"1",
	}
	falseRecord := []string{
		"airy",
		"1",
		"2",
		"3",
		"4",
		"",
		"",
		"",
		"1",
		"1",
	}
	type args struct {
		record []string
	}
	tests := []struct {
		name    string
		args    args
		want    bandCSVFormat
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				record: okRecord,
			},
			want: bandCSVFormat{
				name:       "LiuLiu",
				member1ID:  1,
				member2ID:  2,
				member3ID:  3,
				member4ID:  4,
				member5ID:  5,
				member6ID:  6,
				member7ID:  0,
				memberNum:  6,
				locID:      2,
				bandTypeID: 1,
			},
			wantErr: false,
		},
		{
			name: "OK",
			args: args{
				record: falseRecord,
			},
			want: bandCSVFormat{
				name:       "airy",
				member1ID:  1,
				member2ID:  2,
				member3ID:  3,
				member4ID:  4,
				member5ID:  0,
				member6ID:  0,
				member7ID:  0,
				memberNum:  4,
				locID:      1,
				bandTypeID: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyBandSliceToStruct(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyBandSliceToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyBandSliceToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_bandToStruct(t *testing.T) {
	members := map[int]model.Member{
		1: {ID: 1, Name: "kondo"},
		2: {ID: 2, Name: "hamamoto"},
	}
	locations := model.InitializeLocation()

	okStruct := bandCSVFormat{
		name:       "LiuLiu",
		member1ID:  1,
		member2ID:  2,
		member3ID:  0,
		member4ID:  0,
		member5ID:  0,
		member6ID:  0,
		member7ID:  0,
		memberNum:  2,
		locID:      2,
		bandTypeID: 1,
	}

	bandType, _ := model.SetBandType(1)
	bandMember := []model.Member{members[1], members[2]}

	type args struct {
		bandStruct bandCSVFormat
		members    map[int]model.Member
		locations  map[int]model.Location
	}
	tests := []struct {
		name     string
		args     args
		wantBand model.Band
		wantErr  bool
	}{
		{
			name: "OK",
			args: args{
				bandStruct: okStruct,
				members:    members,
				locations:  locations,
			},
			wantBand: model.Band{
				ID:               1,
				Members:          bandMember,
				BandType:         bandType,
				IsMultiPlay:      false,
				DesireLocationID: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBand, err := bandToStruct(tt.args.bandStruct, tt.args.members, tt.args.locations)
			if (err != nil) != tt.wantErr {
				t.Errorf("bandToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotBand, tt.wantBand) {
				t.Errorf("bandToStruct() = %v, want %v", gotBand, tt.wantBand)
			}
		})
	}
}

func Test_applyMemberSliceToStruct(t *testing.T) {
	type args struct {
		record []string
	}
	tests := []struct {
		name    string
		args    args
		want    memberCSVFormat
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyMemberSliceToStruct(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyMemberSliceToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyMemberSliceToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImportMember(t *testing.T) {
	want := map[int]model.Member{
		1: {ID: 1, Name: "taisho"},
		2: {ID: 2, Name: "haruki"},
		3: {ID: 3, Name: "katsuya"},
		4: {ID: 4, Name: "hinako"},
		5: {ID: 5, Name: "rino"},
		6: {ID: 6, Name: "miran"},
	}

	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    map[int]model.Member
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				fileName: "../csv/test/member.csv",
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImportMember(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportMember() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImportMember() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImportSchedule(t *testing.T) {
	const scheduleCSV = "../csv/test/schedule.csv"

	want := []model.Schedule{
		{
			time.Date(2020, 5, 20, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
			time.Date(2020, 5, 20, 18, 0, 0, 0, time.UTC),
			nil,
		},
		{
			time.Date(2020, 5, 21, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 5, 21, 10, 0, 0, 0, time.UTC),
			time.Date(2020, 5, 21, 18, 0, 0, 0, time.UTC),
			nil,
		},
		{
			time.Date(2020, 5, 22, 0, 0, 0, 0, time.UTC),
			time.Date(2020, 5, 22, 10, 0, 0, 0, time.UTC),
			time.Date(2020, 5, 22, 18, 0, 0, 0, time.UTC),
			nil,
		},
	}

	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		want    []model.Schedule
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				fileName: scheduleCSV,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImportSchedule(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportSchedule() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImportSchedule() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_applyImpossibleTimeSliceToStruct(t *testing.T) {
	record := []string{"1", "2020-05-22", "10h00m", "13h00m"}

	want := impossibleCSVFormat{
		groupID: 1,
		date:    time.Date(2020, 5, 22, 0, 0, 0, 0, time.UTC),
		start:   time.Date(2020, 5, 22, 10, 0, 0, 0, time.UTC),
		end:     time.Date(2020, 5, 22, 13, 0, 0, 0, time.UTC),
	}

	type args struct {
		record []string
	}
	tests := []struct {
		name    string
		args    args
		want    impossibleCSVFormat
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				record: record,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := applyImpossibleTimeSliceToStruct(tt.args.record)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyImpossibleTimeSliceToStruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyImpossibleTimeSliceToStruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
