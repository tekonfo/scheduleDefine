package model

import (
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_setBandType(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    BandType
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				num: 1,
			},
			want: BandType{
				ID:   1,
				Name: "本バンド",
			},
			wantErr: false,
		},
		{
			name: "FALSE",
			args: args{
				num: 4,
			},
			want:    BandType{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SetBandType(tt.args.num)
			if (err != nil) != tt.wantErr {
				t.Errorf("setBandType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("setBandType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBand_AddBandIsMapped(t *testing.T) {
	want := Band{
		ID:       1,
		IsMapped: true,
	}

	type fields struct {
		ID                               int
		Name                             string
		Email                            string
		Members                          []Member
		ImpossibleTimes                  []ImpossibleTime
		DesireLocationID                 int
		BandType                         BandType
		IsMultiPlay                      bool
		WantPrayTime                     map[int]int
		IsHavingMemberAttendingMainStage bool
		IsWantPracticeWithMachine        bool
		IsMapped                         bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    Band
		wantErr bool
	}{
		{
			name: "ok",
			fields: fields{
				ID:       1,
				IsMapped: false,
			},
			want: want,
		},
		{
			name: "except",
			fields: fields{
				ID:       1,
				IsMapped: true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			band := Band{
				ID:                               tt.fields.ID,
				Name:                             tt.fields.Name,
				Email:                            tt.fields.Email,
				Members:                          tt.fields.Members,
				ImpossibleTimes:                  tt.fields.ImpossibleTimes,
				DesireLocationID:                 tt.fields.DesireLocationID,
				BandType:                         tt.fields.BandType,
				IsMultiPlay:                      tt.fields.IsMultiPlay,
				WantPrayTime:                     tt.fields.WantPrayTime,
				IsHavingMemberAttendingMainStage: tt.fields.IsHavingMemberAttendingMainStage,
				IsWantPracticeWithMachine:        tt.fields.IsWantPracticeWithMachine,
				IsMapped:                         tt.fields.IsMapped,
			}
			got, err := band.AddBandIsMapped()
			if (err != nil) != tt.wantErr {
				t.Errorf("Band.AddBandIsMapped() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if diff := cmp.Diff(got, tt.want); diff != "" {
					t.Errorf("Band.AddBandIsMapped() got differs: (-got +want)\n%s", diff)
				}
			}
		})
	}
}

func TestBand_IsImpossibleTime(t *testing.T) {
	type fields struct {
		ID                               int
		Name                             string
		Email                            string
		Members                          []Member
		ImpossibleTimes                  []ImpossibleTime
		DesireLocationID                 int
		BandType                         BandType
		IsMultiPlay                      bool
		WantPrayTime                     map[int]int
		IsHavingMemberAttendingMainStage bool
		IsWantPracticeWithMachine        bool
		IsMapped                         bool
	}
	type args struct {
		targetStartTime time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true",
			fields: fields{
				WantPrayTime: map[int]int{
					1: 10,
				},
				DesireLocationID: 1,
				ImpossibleTimes: []ImpossibleTime{
					{
						Start: time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
						End:   time.Date(2020, 5, 20, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				targetStartTime: time.Date(2020, 5, 20, 12, 0, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "true2",
			fields: fields{
				WantPrayTime: map[int]int{
					1: 10,
				},
				DesireLocationID: 1,
				ImpossibleTimes: []ImpossibleTime{
					{
						Start: time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
						End:   time.Date(2020, 5, 20, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				targetStartTime: time.Date(2020, 5, 20, 9, 50, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "true3",
			fields: fields{
				WantPrayTime: map[int]int{
					1: 10,
				},
				DesireLocationID: 1,
				ImpossibleTimes: []ImpossibleTime{
					{
						Start: time.Date(2020, 5, 20, 10, 0, 0, 0, time.UTC),
						End:   time.Date(2020, 5, 20, 12, 0, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				targetStartTime: time.Date(2020, 5, 20, 12, 50, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			band := Band{
				ID:                               tt.fields.ID,
				Name:                             tt.fields.Name,
				Email:                            tt.fields.Email,
				Members:                          tt.fields.Members,
				ImpossibleTimes:                  tt.fields.ImpossibleTimes,
				DesireLocationID:                 tt.fields.DesireLocationID,
				BandType:                         tt.fields.BandType,
				IsMultiPlay:                      tt.fields.IsMultiPlay,
				WantPrayTime:                     tt.fields.WantPrayTime,
				IsHavingMemberAttendingMainStage: tt.fields.IsHavingMemberAttendingMainStage,
				IsWantPracticeWithMachine:        tt.fields.IsWantPracticeWithMachine,
				IsMapped:                         tt.fields.IsMapped,
			}
			if got := band.IsImpossibleTime(tt.args.targetStartTime); got != tt.want {
				t.Errorf("Band.IsImpossibleTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
