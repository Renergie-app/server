package graph

import (
	"renergie-server/graph/model"
	"testing"
)

func TestPercentageWithOrientationAndAngle(t *testing.T) {
	type args struct {
		orientation model.Orientation
		angle       int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "East with angle of 0", args: args{
			orientation: model.OrientationEast,
			angle:       0,
		}, want: 93.0},
		{name: "East with angle of 30", args: args{
			orientation: model.OrientationEast,
			angle:       30,
		}, want: 90.0},
		{name: "East with angle of 60", args: args{
			orientation: model.OrientationEast,
			angle:       60,
		}, want: 78.0},
		{name: "East with angle of 90", args: args{
			orientation: model.OrientationEast,
			angle:       90,
		}, want: 55.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PercentageWithOrientationAndAngle(tt.args.orientation, tt.args.angle); got != tt.want {
				t.Errorf("PercentageWithOrientationAndAngle() = %v, want %v", got, tt.want)
			}
		})
	}
}
