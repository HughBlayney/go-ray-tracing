package objects

import (
	"testing"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/rays"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/utils"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

func TestSphere_CollideDistances(t *testing.T) {
	type fields struct {
		Radius float64
		Center vectors.Vector
	}
	type args struct {
		test_ray rays.Ray
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []float64
	}{
		{
			name:   "Middle of sphere, 2 collisions",
			fields: fields{Radius: 1.0, Center: vectors.Vector{X: 0.0, Y: 10.0, Z: 10.0}},
			args: args{test_ray: rays.Ray{
				Origin:    &vectors.Vector{X: 0.0, Y: 0.0, Z: 10.0},
				Direction: &vectors.Vector{X: 0.0, Y: 1.0, Z: 0.0}}},
			want: []float64{9.0, 11.0},
		},
		{
			name:   "Edge of sphere, 1 collision",
			fields: fields{Radius: 1.0, Center: vectors.Vector{X: 0.0, Y: 10.0, Z: 10.0}},
			args: args{test_ray: rays.Ray{
				Origin:    &vectors.Vector{X: 0.0, Y: 0.0, Z: 11.0},
				Direction: &vectors.Vector{X: 0.0, Y: 1.0, Z: 0.0}}},
			want: []float64{10.0},
		},
		{
			name:   "Miss sphere, 0 collisions",
			fields: fields{Radius: 1.0, Center: vectors.Vector{X: 0.0, Y: 10.0, Z: 10.0}},
			args: args{test_ray: rays.Ray{
				Origin:    &vectors.Vector{X: 0.0, Y: 0.0, Z: 12.0},
				Direction: &vectors.Vector{X: 0.0, Y: 1.0, Z: 0.0}}},
			want: []float64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere{
				Radius: tt.fields.Radius,
				Center: tt.fields.Center,
			}
			if got := s.CollideDistances(tt.args.test_ray); !utils.Slice_close_enough(got, tt.want) {
				t.Errorf("Sphere.CollideDistances() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_Normal(t *testing.T) {
	type fields struct {
		Radius float64
		Center vectors.Vector
	}
	type args struct {
		surface_point *vectors.Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *vectors.Vector
	}{
		{
			name:   "Centered at origin",
			fields: fields{Radius: 1.0, Center: vectors.Vector{X: 0.0, Y: 0.0, Z: 0.0}},
			args:   args{surface_point: &vectors.Vector{X: 0.0, Y: 1.0, Z: 0.0}},
			want:   &vectors.Vector{X: 0.0, Y: 1.0, Z: 0.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere{
				Radius: tt.fields.Radius,
				Center: tt.fields.Center,
			}
			if got := s.Normal(tt.args.surface_point); !got.CloseTo(tt.want) {
				t.Errorf("Sphere.Normal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSphere_Reflect(t *testing.T) {
	type fields struct {
		Radius float64
		Center vectors.Vector
	}
	type args struct {
		incoming_ray          rays.Ray
		point_of_intersection *vectors.Vector
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   rays.Ray
	}{
		{
			name:   "Reflect off origin",
			fields: fields{Radius: 1.0, Center: vectors.Vector{X: 0.0, Y: -1.0, Z: 0.0}},
			args: args{
				incoming_ray: rays.MakeRay(
					&vectors.Vector{
						X: -1.0, Y: 1.0,
					},
					&vectors.Vector{
						X: 1.0, Y: -1.0,
					},
				),
				point_of_intersection: &vectors.Vector{
					X: 0.0, Y: 0.0,
				},
			},
			want: rays.MakeRay(
				&vectors.Vector{
					X: 0.0, Y: 0.0,
				},
				&vectors.Vector{
					X: 1.0, Y: 1.0,
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sphere{
				Radius: tt.fields.Radius,
				Center: tt.fields.Center,
			}
			if got := s.Reflect(tt.args.incoming_ray, tt.args.point_of_intersection); !got.CloseTo(tt.want) {
				t.Errorf("Sphere.Reflect() = %s, want %s", got.Print(), tt.want.Print())
			}
		})
	}
}
