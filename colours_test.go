package gocolournamer

import (
	"reflect"
	"testing"
)

func Test_populate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			"Populated array test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			populate()
			if colours[0].b != 44 {
				t.Errorf("populate() failed to populate array")
			}
		})
	}
}

func TestToNearestColour(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name      string
		args      args
		wantNames Named
		wantErr   bool
	}{
		{
			"Valid white (6) hex test - no hash",
			args{
				"FFFFFF",
			},
			Named{
				Hex:    "FFFFFF",
				Colour: "White",
				Hue:    "White",
				Huehex: "FFFFFF",
			},
			false,
		},
		{
			"Valid white (3) hex test - no hash",
			args{
				"FFF",
			},
			Named{
				Hex:    "FFFFFF",
				Colour: "White",
				Hue:    "White",
				Huehex: "FFFFFF",
			},
			false,
		},
		{
			"Valid white (3) hex test - hash",
			args{
				"#FFFFFF",
			},
			Named{
				Hex:    "FFFFFF",
				Colour: "White",
				Hue:    "White",
				Huehex: "FFFFFF",
			},
			false,
		},
		{
			"Valid white (6) hex test - hash",
			args{
				"#FFFFFF",
			},
			Named{
				Hex:    "FFFFFF",
				Colour: "White",
				Hue:    "White",
				Huehex: "FFFFFF",
			},
			false,
		},
		{
			"Valid Colour - slightly off match",
			args{
				"7D9D73",
			},
			Named{
				Hex:    "7D9D73",
				Colour: "Amulet",
				Hue:    "Green",
				Huehex: "008000",
			},
			false,
		},
		{
			"Invalid length hex",
			args{
				"7D972",
			},
			Named{},
			true,
		},
		{
			"Invalid chracter hex",
			args{
				"7D9Q72",
			},
			Named{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNames, err := ToNearestColour(tt.args.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNearestColour() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotNames, tt.wantNames) {
				t.Errorf("ToNearestColour() = %v, want %v", gotNames, tt.wantNames)
			}
		})
	}
}

func Test_validHex(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name    string
		args    args
		wantH   string
		wantErr bool
	}{
		{
			"Valid Colour - 6 hash",
			args{
				"#7D9D72",
			},
			"7D9D72",
			false,
		},
		{
			"Valid Colour - 6 no hash",
			args{
				"7D9D72",
			},
			"7D9D72",
			false,
		},
		{
			"Valid Colour - 3 no hash",
			args{
				"888",
			},
			"888888",
			false,
		},
		{
			"Invalid length hex",
			args{
				"7D972",
			},
			"",
			true,
		},
		{
			"Invalid chracter hex",
			args{
				"7D9Q72",
			},
			"",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotH, err := validHex(tt.args.hex)
			if (err != nil) != tt.wantErr {
				t.Errorf("validHex() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotH != tt.wantH {
				t.Errorf("validHex() = %v, want %v", gotH, tt.wantH)
			}
		})
	}
}

func Test_rgb(t *testing.T) {
	type args struct {
		hex string
	}
	tests := []struct {
		name  string
		args  args
		wantR int64
		wantG int64
		wantB int64
	}{
		{
			"Standard RGB Test white",
			args{
				"FFFFFF",
			},
			255,
			255,
			255,
		},
		{
			"Standard RGB Test random",
			args{
				"74F2a8",
			},
			116,
			242,
			168,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotR, gotG, gotB := rgb(tt.args.hex)
			if gotR != tt.wantR {
				t.Errorf("rgb() gotR = %v, want %v", gotR, tt.wantR)
			}
			if gotG != tt.wantG {
				t.Errorf("rgb() gotG = %v, want %v", gotG, tt.wantG)
			}
			if gotB != tt.wantB {
				t.Errorf("rgb() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_hsl(t *testing.T) {
	type args struct {
		ri int64
		gi int64
		bi int64
	}
	tests := []struct {
		name  string
		args  args
		wantH float64
		wantS float64
		wantL float64
	}{
		{
			"White",
			args{
				255,
				255,
				255,
			},
			0,
			0.0,
			100,
		},
		{
			"Random",
			args{
				116,
				242,
				168,
			},
			145,
			82.9,
			70.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotH, gotS, gotL := hsl(tt.args.ri, tt.args.gi, tt.args.bi)
			if gotH != tt.wantH {
				t.Errorf("hsl() gotH = %v, want %v", gotH, tt.wantH)
			}
			if gotS != tt.wantS {
				t.Errorf("hsl() gotS = %v, want %v", gotS, tt.wantS)
			}
			if gotL != tt.wantL {
				t.Errorf("hsl() gotL = %v, want %v", gotL, tt.wantL)
			}
		})
	}
}

func Test_shade(t *testing.T) {
	type args struct {
		shadename string
	}
	tests := []struct {
		name    string
		args    args
		wantHex string
	}{
		{
			"White",
			args{
				"White",
			},
			"FFFFFF",
		},
		{
			"Blue",
			args{
				"Blue",
			},
			"0000FF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotHex := shade(tt.args.shadename); gotHex != tt.wantHex {
				t.Errorf("shade() = %v, want %v", gotHex, tt.wantHex)
			}
		})
	}
}
