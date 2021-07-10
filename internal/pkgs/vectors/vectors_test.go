package vectors

import (
	"math"
	"testing"
)

const tolerance float64 = 0.0001

func close_enough(x, y float64) bool {
	return math.Abs(x-y) < tolerance
}

func TestVector_Magnitude(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Positive components (no z)",
			fields: fields{3.0, 4.0, 0.0},
			want:   5.0,
		},
		{
			name:   "Positive components (no x)",
			fields: fields{0.0, 3.0, 4.0},
			want:   5.0,
		},
		{
			name:   "Zero vector",
			fields: fields{0.0, 0.0, 0.0},
			want:   0.0,
		},
		{
			name:   "Negative components",
			fields: fields{-3.0, -4.0, 0.0},
			want:   5.0,
		},
		{
			name:   "Mixed positive & negative components",
			fields: fields{3.0, -4.0, 0.0},
			want:   5.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			if got := v.Magnitude(); !(close_enough(got, tt.want)) {
				t.Errorf("Vector.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Normalise(t *testing.T) {
	type fields struct {
		x float64
		y float64
		z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Positive components (no z)",
			fields: fields{3.0, 4.0, 0.0},
			want:   1.0,
		},
		{
			name:   "Positive components (no x)",
			fields: fields{0.0, 3.0, 4.0},
			want:   1.0,
		},
		{
			name:   "Zero vector",
			fields: fields{0.0, 0.0, 0.0},
			want:   0.0,
		},
		{
			name:   "Negative components",
			fields: fields{-3.0, -4.0, 0.0},
			want:   1.0,
		},
		{
			name:   "Mixed positive & negative components",
			fields: fields{3.0, -4.0, 0.0},
			want:   1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				x: tt.fields.x,
				y: tt.fields.y,
				z: tt.fields.z,
			}
			v.Normalise()
			if got := v.Magnitude(); !(close_enough(got, tt.want)) {
				t.Errorf("After normalisation, Vector.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
