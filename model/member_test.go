package model

import (
	"testing"
)

func TestMember_CheckDoubleCheck(t *testing.T) {
	type fields struct {
		name  string
		bands []Band
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			member := Member{
				Name:  tt.fields.name,
				Bands: tt.fields.bands,
			}
			if got := member.CheckDoubleCheck(); got != tt.want {
				t.Errorf("Member.CheckDoubleCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
