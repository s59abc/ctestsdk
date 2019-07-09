package ctydat

import (
	"ctestsdk/adif"
	"ctestsdk/geo"
	"ctestsdk/spot"
	"reflect"
	"testing"
)

func Test_get(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name  string
		args  args
		want  spot.CtyDta
		want1 bool
	}{
		{
			name: "EmptyKey",
			args: args{
				key: "",
			},
			want: spot.CtyDta{
				CountryName:   "",
				PrimaryPrefix: "",
				AliasPrefix:   "",
				Continent:     0,
				CqZone:        0,
				ItuZone:       0,
				LatLon: geo.LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				TimeOffset: "",
			},
			want1: false,
		},
		{
			name: "WrongKey-1",
			args: args{
				key: "x",
			},
			want: spot.CtyDta{
				CountryName:   "",
				PrimaryPrefix: "",
				AliasPrefix:   "",
				Continent:     0,
				CqZone:        0,
				ItuZone:       0,
				LatLon: geo.LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				TimeOffset: "",
			},
			want1: false,
		},
		{
			name: "WrongKey-2",
			args: args{
				key: ":",
			},
			want: spot.CtyDta{
				CountryName:   "",
				PrimaryPrefix: "",
				AliasPrefix:   "",
				Continent:     0,
				CqZone:        0,
				ItuZone:       0,
				LatLon: geo.LatLonDeg{
					Lat: 0,
					Lon: 0,
				},
				TimeOffset: "",
			},
			want1: false,
		},

		{
			name: "S51DS",
			args: args{
				key: "S51DS",
			},
			want: spot.CtyDta{
				CountryName:   "Slovenia",
				PrimaryPrefix: "S5",
				AliasPrefix:   "S5",
				Continent:     adif.EU,
				CqZone:        adif.CQZONE15,
				ItuZone:       adif.ITUZONE28,
				LatLon: geo.LatLonDeg{
					Lat: 46.0000,
					Lon: -14.0000,
				},
				TimeOffset: "-1.0",
			},
			want1: true,
		},

		{
			name: "9A/S51DS/P",
			args: args{
				key: "9A/S51DS/P",
			},
			want: spot.CtyDta{
				CountryName:   "Croatia",
				PrimaryPrefix: "9A",
				AliasPrefix:   "9A",
				Continent:     adif.EU,
				CqZone:        adif.CQZONE15,
				ItuZone:       adif.ITUZONE28,
				LatLon: geo.LatLonDeg{
					Lat: 45.1800,
					Lon: -15.3000,
				},
				TimeOffset: "-1.0",
			},
			want1: true,
		},

		{
			name: "KH2/AC2AI",
			args: args{
				key: "KH2/AC2AI",
			},
			want: spot.CtyDta{
				CountryName:   "Guam",
				PrimaryPrefix: "KH2",
				AliasPrefix:   "KH2",
				Continent:     adif.OC,
				CqZone:        adif.CQZONE27,
				ItuZone:       adif.ITUZONE64,
				LatLon: geo.LatLonDeg{
					Lat: 13.3700,
					Lon: -144.7000,
				},
				TimeOffset: "-10.0",
			},
			want1: true,
		},

		{
			name: "AC2AI/KH2",
			args: args{
				key: "AC2AI/KH2",
			},
			want: spot.CtyDta{
				CountryName:   "Guam",
				PrimaryPrefix: "KH2",
				AliasPrefix:   "KH2",
				Continent:     adif.OC,
				CqZone:        adif.CQZONE27,
				ItuZone:       adif.ITUZONE64,
				LatLon: geo.LatLonDeg{
					Lat: 13.3700,
					Lon: -144.7000,
				},
				TimeOffset: "-10.0",
			},
			want1: true,
		},

		{
			name: "I20000X",
			args: args{
				key: "I20000X",
			},
			want: spot.CtyDta{
				CountryName:   "Italy",
				PrimaryPrefix: "I",
				AliasPrefix:   "I2",
				Continent:     adif.EU,
				CqZone:        adif.CQZONE15,
				ItuZone:       adif.ITUZONE28,
				LatLon: geo.LatLonDeg{
					Lat: 45.4700,
					Lon: -9.2000,
				},
				TimeOffset: "-1.0",
			},
			want1: true,
		},

		{
			name: "3DA0RS",
			args: args{
				key: "3DA0RS",
			},
			want: spot.CtyDta{
				CountryName:   "Swaziland",
				PrimaryPrefix: "3DA",
				AliasPrefix:   "3DA",
				Continent:     adif.AF,
				CqZone:        adif.CQZONE38,
				ItuZone:       adif.ITUZONE57,
				LatLon: geo.LatLonDeg{
					Lat: -26.6500,
					Lon: -31.4800,
				},
				TimeOffset: "-2.0",
			},
			want1: true,
		},

		{
			name: "K4X",
			args: args{
				key: "K4X",
			},
			want: spot.CtyDta{
				CountryName:   "United States",
				PrimaryPrefix: "K",
				AliasPrefix:   "K4",
				Continent:     adif.NA,
				CqZone:        adif.CQZONE5,
				ItuZone:       adif.ITUZONE8,
				LatLon: geo.LatLonDeg{
					Lat: 33.1800,
					Lon: 82.2300,
				},
				TimeOffset: "5.0",
			},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := get(tt.args.key)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("get() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("get() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
