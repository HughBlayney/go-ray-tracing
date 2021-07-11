package vectors

import (
	"testing"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/utils"
)

func TestVector_Magnitude(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Positive components (no Z)",
			fields: fields{3.0, 4.0, 0.0},
			want:   5.0,
		},
		{
			name:   "Positive components (no X)",
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
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			if got := v.Magnitude(); !(utils.Close_enough(got, tt.want)) {
				t.Errorf("Vector.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVector_Normalise(t *testing.T) {
	type fields struct {
		X float64
		Y float64
		Z float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name:   "Positive components (no Z)",
			fields: fields{3.0, 4.0, 0.0},
			want:   1.0,
		},
		{
			name:   "Positive components (no X)",
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
			name:   "MiXed positive & negative components",
			fields: fields{3.0, -4.0, 0.0},
			want:   1.0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Vector{
				X: tt.fields.X,
				Y: tt.fields.Y,
				Z: tt.fields.Z,
			}
			v.Normalise()
			if got := v.Magnitude(); !(utils.Close_enough(got, tt.want)) {
				t.Errorf("After normalisation, Vector.Magnitude() = %v, want %v", got, tt.want)
			}
		})
	}
}
