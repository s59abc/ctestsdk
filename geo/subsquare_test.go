package geo

import (
	"reflect"
	"testing"
)

func TestSubsquareEncode(t *testing.T) {
	type args struct {
		lld LatLonDeg
	}
	tests := []struct {
		name  string
		args  args
		want  field
		want1 square
		want2 subsquare
	}{
		{
			name: "encode-S59ABC-JN76TO",
			args: args{
				lld: LatLonDeg{
					lat: 46.60333,
					lon: 15.62333,
				},
			},
			want: field{
				decoded: LatLonDeg{
					lat: 40,
					lon: 0,
				},
				encoded: latLonChar{
					latChar: 78,
					lonChar: 74,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					lat: 6,
					lon: 14,
				},
				encoded: latLonChar{
					latChar: 54,
					lonChar: 55,
				},
			},
			want2: subsquare{
				decoded: LatLonDeg{
					lat: 35, //minutes
					lon: 95, //minutes
				},
				encoded: latLonChar{
					latChar: 79, // O
					lonChar: 84, // T
				},
			},
		},

		{
			name: "encode-K1TTT-FN32LL",
			args: args{
				lld: LatLonDeg{
					lat: 42.4662,
					lon: -73.0232,
				},
			},
			want: field{
				decoded: LatLonDeg{
					lat: 40,
					lon: -80,
				},
				encoded: latLonChar{
					latChar: 78, //N
					lonChar: 70, //F
				},
			},
			want1: square{
				decoded: LatLonDeg{
					lat: 2,
					lon: 6,
				},
				encoded: latLonChar{
					latChar: 50, //2
					lonChar: 51, //3
				},
			},
			want2: subsquare{
				decoded: LatLonDeg{
					lat: 27.5, //minutes
					lon: 55,   //minutes
				},
				encoded: latLonChar{
					latChar: 76, // L
					lonChar: 76, // L
				},
			},
		},

		{
			name: "encode-PS2T-GG58WG",
			args: args{
				lld: LatLonDeg{
					lat: -21.7487,
					lon: -48.1268,
				},
			},
			want: field{
				decoded: LatLonDeg{
					lat: -30,
					lon: -60,
				},
				encoded: latLonChar{
					latChar: 71,
					lonChar: 71,
				},
			},
			want1: square{
				decoded: LatLonDeg{
					lat: 8,
					lon: 10,
				},
				encoded: latLonChar{
					latChar: 56,
					lonChar: 53,
				},
			},
			want2: subsquare{
				decoded: LatLonDeg{
					lat: 15,
					lon: 110,
				},
				encoded: latLonChar{
					latChar: 71,
					lonChar: 87,
				},
			},
		},

		{
			name: "encode-ZM4T-RF80LQ",
			args: args{
				lld: LatLonDeg{
					lat: -39.3125,
					lon: 176.9583333,
				},
			},
			want: field{
				decoded: LatLonDeg{
					lat: -40,
					lon: 160,
				},
				encoded: latLonChar{
					latChar: 70, // F
					lonChar: 82, // R
				},
			},
			want1: square{
				decoded: LatLonDeg{
					lat: 0,
					lon: 16,
				},
				encoded: latLonChar{
					latChar: 48, // 0
					lonChar: 56, // 8
				},
			},
			want2: subsquare{
				decoded: LatLonDeg{
					lat: 40,
					lon: 55,
				},
				encoded: latLonChar{
					latChar: 81, // L
					lonChar: 76, // Q
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := subsquareEncode(tt.args.lld)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fieldEncode() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("squareEncode() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("subsquareEncode() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestSubsquareDecode(t *testing.T) {
	type args struct {
		llc latLonChar
	}
	tests := []struct {
		name string
		args args
		want subsquare
	}{
		{
			name: "decode-zero",
			args: args{
				llc: latLonChar{
					latChar: 0,
					lonChar: 0,
				},
			},
			want: subsquare{
				decoded: LatLonDeg{
					lat: 0,
					lon: 0,
				},
				encoded: latLonChar{
					latChar: 0,
					lonChar: 0,
				},
			},
		},

		{
			name: "decode-AA",
			args: args{
				llc: latLonChar{
					latChar: 65,
					lonChar: 65,
				},
			},
			want: subsquare{
				decoded: LatLonDeg{
					lat: 0,
					lon: 0,
				},
				encoded: latLonChar{
					latChar: 65,
					lonChar: 65,
				},
			},
		},

		{
			name: "decode-XX",
			args: args{
				llc: latLonChar{
					latChar: 88,
					lonChar: 88,
				},
			},
			want: subsquare{
				decoded: LatLonDeg{
					lat: 57.5,
					lon: 115,
				},
				encoded: latLonChar{
					latChar: 88,
					lonChar: 88,
				},
			},
		},

		{
			name: "decode-S59ABC-JN76TO",
			args: args{
				llc: latLonChar{
					latChar: 79, // O
					lonChar: 84, // T
				},
			},
			want: subsquare{
				decoded: LatLonDeg{
					lat: 35, //minutes
					lon: 95, //minutes
				},
				encoded: latLonChar{
					latChar: 79, // O
					lonChar: 84, // T
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsquareDecode(tt.args.llc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subsquareDecode() = %v, want %v", got, tt.want)
			}
		})
	}
}
