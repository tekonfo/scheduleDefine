package model

import (
	"reflect"
	"testing"
)

func TestImpossibleBandOrder_IsExistBandOrder(t *testing.T) {
	bandID := 4

	impossibleBandOrders := []ImpossibleBandOrder{
		{
			1, 2, 3, 4, 5,
		},
		{
			1, 2, 3, 4,
		},
	}

	type args struct {
		bandID               int
		impossibleBandOrders []ImpossibleBandOrder
	}
	tests := []struct {
		name             string
		currentBandOrder ImpossibleBandOrder
		args             args
		want             bool
	}{
		{
			name: "false",
			args: args{
				bandID:               bandID,
				impossibleBandOrders: impossibleBandOrders,
			},
			currentBandOrder: ImpossibleBandOrder{
				1,
			},
			want: false,
		},
		{
			name: "true",
			args: args{
				bandID:               bandID,
				impossibleBandOrders: impossibleBandOrders,
			},
			currentBandOrder: ImpossibleBandOrder{
				1, 2, 3,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.currentBandOrder.IsExistBandOrder(tt.args.bandID, tt.args.impossibleBandOrders); got != tt.want {
				t.Errorf("ImpossibleBandOrder.IsExistBandOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestImpossibleBandOrder_DeleteBandOrder(t *testing.T) {
	tests := []struct {
		name             string
		currentBandOrder ImpossibleBandOrder
		want             ImpossibleBandOrder
		wantErr          bool
	}{
		{
			name:             "true",
			currentBandOrder: ImpossibleBandOrder{1, 2, 3, 4, 5},
			want:             ImpossibleBandOrder{1, 2, 3, 4},
		},
		{
			name:             "except",
			currentBandOrder: ImpossibleBandOrder{},
			wantErr:          true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.currentBandOrder.DeleteBandOrder()
			if (err != nil) != tt.wantErr {
				t.Errorf("ImpossibleBandOrder.DeleteBandOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImpossibleBandOrder.DeleteBandOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
