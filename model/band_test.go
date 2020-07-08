package model

import (
	"reflect"
	"testing"
)

func Test_setBandType(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name    string
		args    args
		want    bandType
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				num: 1,
			},
			want: bandType{
				id:   1,
				name: "本バンド",
			},
			wantErr: false,
		},
		{
			name: "FALSE",
			args: args{
				num: 4,
			},
			want:    bandType{},
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
