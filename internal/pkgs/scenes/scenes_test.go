package scenes

import (
	"image/color"
	"reflect"
	"testing"

	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/lights"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/materials"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/objects"
	"github.com/HughBlayney/go-ray-tracing/internal/pkgs/vectors"
)

func Test_computeDiffuseSpecular(t *testing.T) {
	type args struct {
		Diffuse_const    [3]float64
		Specular_const   [3]float64
		alpha            float64
		light_position   *vectors.Vector
		surface_position *vectors.Vector
		surface_normal   *vectors.Vector
		viewer_direction *vectors.Vector
	}
	tests := []struct {
		name  string
		args  args
		want  [3]float64
		want1 [3]float64
	}{
		{
			name: "Direct dot",
			args: args{
				Diffuse_const:    [3]float64{1.0, 1.0, 1.0},
				Specular_const:   [3]float64{1.0, 1.0, 1.0},
				alpha:            0.5,
				light_position:   &vectors.Vector{0.0, 1.0, 0.0},
				surface_position: &vectors.Vector{0.0, 0.0, 0.0},
				surface_normal:   &vectors.Vector{0.0, 1.0, 0.0},
				viewer_direction: &vectors.Vector{0.0, 1.0, 0.0},
			},
			want:  [3]float64{1.0, 1.0, 1.0},
			want1: [3]float64{1.0, 1.0, 1.0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var objs []objects.Object
			got, got1 := computeDiffuseSpecular(tt.args.Diffuse_const, tt.args.Specular_const, tt.args.alpha, tt.args.light_position, tt.args.surface_position, tt.args.surface_normal, tt.args.viewer_direction, objs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("computeDiffuseSpecular() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("computeDiffuseSpecular() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_clipFloat(t *testing.T) {
	type args struct {
		float_val float64
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		{
			name: "No clip",
			args: args{float_val: 100.0},
			want: uint8(100),
		},
		{
			name: "Clip",
			args: args{float_val: 256.0},
			want: uint8(255),
		},
		{
			name: "Huge clip",
			args: args{float_val: 10000000000.0},
			want: uint8(255),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := clipFloat(tt.args.float_val); got != tt.want {
				t.Errorf("clipFloat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ComputePhong(t *testing.T) {
	type fields struct {
		Color           color.RGBA
		Specular_const  float64
		Diffuse_const   float64
		Ambient_const   float64
		Shininess_const float64
		Matte_const     float64
	}
	type args struct {
		lights           []lights.Light
		Ambient_color    color.RGBA
		surface_position *vectors.Vector
		surface_normal   *vectors.Vector
		viewer_direction *vectors.Vector
	}
	tests := []struct {
		name             string
		fields           fields
		args             args
		wantIllumination color.RGBA
	}{
		{
			name: "Direct dot",
			fields: fields{
				Color:           color.RGBA{0xff, 0xff, 0xff, 0xff},
				Specular_const:  0.5,
				Diffuse_const:   0.5,
				Ambient_const:   0.5,
				Shininess_const: 0.005,
				Matte_const:     0.5,
			},
			args: args{
				lights:           []lights.Light{{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Intensity: 1.0, Position: vectors.Vector{0.0, 1.0, 0.0}}},
				Ambient_color:    color.RGBA{0.0, 0.0, 0.0, 0.0},
				surface_position: &vectors.Vector{0.0, 0.0, 0.0},
				surface_normal:   &vectors.Vector{0.0, 1.0, 0.0},
				viewer_direction: &vectors.Vector{0.0, 1.0, 0.0},
			},
			wantIllumination: color.RGBA{0xff, 0xff, 0xff, 0xff},
		},
		{
			name: "Direct eclipse of light",
			fields: fields{
				Color:           color.RGBA{0xff, 0xff, 0xff, 0xff},
				Specular_const:  0.5,
				Diffuse_const:   0,
				Ambient_const:   0,
				Shininess_const: 0.005,
				Matte_const:     0.5,
			},
			args: args{
				lights:           []lights.Light{{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Intensity: 1.0, Position: vectors.Vector{1.0, -1.0, 0.0}}},
				Ambient_color:    color.RGBA{0.0, 0.0, 0.0, 0.0},
				surface_position: &vectors.Vector{0.0, 0.0, 0.0},
				surface_normal:   &vectors.Vector{-1.0, -1.0, 0.0},
				viewer_direction: &vectors.Vector{0.0, 1.0, 0.0},
			},
			wantIllumination: color.RGBA{0, 0, 0, 0xff},
		},
		{
			name: "Indirect eclipse of light",
			fields: fields{
				Color:           color.RGBA{0xff, 0xff, 0xff, 0xff},
				Specular_const:  0.5,
				Diffuse_const:   0,
				Ambient_const:   0,
				Shininess_const: 0.005,
				Matte_const:     0.5,
			},
			args: args{
				lights:           []lights.Light{{Color: color.RGBA{0xff, 0xff, 0xff, 0xff}, Intensity: 1.0, Position: vectors.Vector{0.0, -1.0, 0.0}}},
				Ambient_color:    color.RGBA{0.0, 0.0, 0.0, 0.0},
				surface_position: &vectors.Vector{1.0, 0.0, 0.0},
				surface_normal:   &vectors.Vector{1.0, 1.0, 0.0},
				viewer_direction: &vectors.Vector{0.0, 1.0, 0.0},
			},
			wantIllumination: color.RGBA{0, 0, 0, 0xff},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var objs []objects.Object
			m := materials.MakeMaterial(
				tt.fields.Color,
				tt.fields.Diffuse_const,
				tt.fields.Specular_const,
				tt.fields.Ambient_const,
				tt.fields.Shininess_const,
				tt.fields.Matte_const,
			)
			if gotIllumination := ComputePhong(m, tt.args.lights, objs, tt.args.Ambient_color, tt.args.surface_position, tt.args.surface_normal, tt.args.viewer_direction); !reflect.DeepEqual(gotIllumination, tt.wantIllumination) {
				t.Errorf("Material.ComputePhong() = %v, want %v", gotIllumination, tt.wantIllumination)
			}
		})
	}
}
