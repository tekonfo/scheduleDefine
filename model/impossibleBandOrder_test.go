package model

import "testing"

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
