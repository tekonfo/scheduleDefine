package model

import (
	"testing"
	"time"
)

func TestCheckValidMembers(t *testing.T) {
	memberA := Member{
		ID:   1,
		Name: "hoge",
	}
	memberB := Member{
		ID:   2,
		Name: "taro",
	}
	members := map[int]Member{
		1: memberA,
		2: memberB,
	}

	type args struct {
		ids     []int
		members map[int]Member
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				ids:     []int{1, 2},
				members: members,
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "False",
			args: args{
				ids:     []int{3},
				members: members,
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckValidMembers(tt.args.ids, tt.args.members)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckValidMembers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckValidMembers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMember_IsViolateTimeRule(t *testing.T) {
	type fields struct {
		ID      int
		Name    string
		BandIDs []int
		Events  []Event
	}
	type args struct {
		startTime time.Time
		endTime   time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "true: violate starttime",
			fields: fields{
				Events: []Event{
					{
						Start: time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
						End:   time.Date(2020, 5, 20, 11, 05, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				startTime: time.Date(2020, 5, 20, 11, 04, 0, 0, time.UTC),
				endTime:   time.Date(2020, 5, 20, 11, 14, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "true: violate movetime",
			fields: fields{
				Events: []Event{
					{
						Start: time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
						End:   time.Date(2020, 5, 20, 11, 05, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				startTime: time.Date(2020, 5, 20, 11, 15, 0, 0, time.UTC),
				endTime:   time.Date(2020, 5, 20, 11, 25, 0, 0, time.UTC),
			},
			want: true,
		},
		{
			name: "false",
			fields: fields{
				Events: []Event{
					{
						Start: time.Date(2020, 5, 20, 10, 55, 0, 0, time.UTC),
						End:   time.Date(2020, 5, 20, 11, 05, 0, 0, time.UTC),
					},
				},
			},
			args: args{
				startTime: time.Date(2020, 5, 20, 11, 20, 0, 0, time.UTC),
				endTime:   time.Date(2020, 5, 20, 11, 30, 0, 0, time.UTC),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			member := Member{
				ID:      tt.fields.ID,
				Name:    tt.fields.Name,
				BandIDs: tt.fields.BandIDs,
				Events:  tt.fields.Events,
			}
			if got := member.IsViolateTimeRule(tt.args.startTime, tt.args.endTime); got != tt.want {
				t.Errorf("Member.IsViolateTimeRule() = %v, want %v", got, tt.want)
			}
		})
	}
}
