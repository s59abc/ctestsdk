package ctydat

import (
	"ctestsdk/adif"
	"ctestsdk/geo"
	"ctestsdk/spot"
	"github.com/golang/geo/s2"
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
					Lon: 14.0000,
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
					Lon: 15.3000,
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
					Lon: 144.7000,
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
					Lon: 144.7000,
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
					Lon: 9.2000,
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
					Lon: 31.4800,
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
					Lon: -82.2300,
				},
				TimeOffset: "5.0",
			},
			want1: true,
		},

		{
			name: "S51DS-1",
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
					Lon: 14.0000,
				},
				TimeOffset: "-1.0",
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

func TestAddCtyDat(t *testing.T) {
	type args struct {
		spot spot.Data
	}
	tests := []struct {
		name string
		args args
		want spot.Data
	}{
		{
			name: "Test-01",
			args: args{
				spot: spot.Data{
					Dx:       "S59ABC",
					De:       "S50ARX",
					Freq:     "3502.8",
					Raw:      "DX De S50ARX-#:     3502.8  S59ABC      CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
					Comments: "CW    14 dB  26 WPM  CQ       1118Z JN76\a\a",
					Source:   "",
					IsRbn:    true,
					BAND:     adif.Band80M,
					MODE:     adif.CW,
					DxQTH:    geo.QTH{},
					DxCtyDta: spot.CtyDta{
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
					DeQTH: geo.QTH{},
					DeCtyDta: spot.CtyDta{
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
				},
			},
			want: spot.Data{
				Dx:       "S59ABC",
				De:       "S50ARX",
				Freq:     "3502.8",
				Raw:      "DX De S50ARX-#:     3502.8  S59ABC      CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				Comments: "CW    14 dB  26 WPM  CQ       1118Z JN76\a\a",
				Source:   "",
				IsRbn:    true,
				BAND:     adif.Band80M,
				MODE:     adif.CW,
				DxQTH: geo.QTH{
					Loc: "JN76AA",
					LatLon: geo.LatLonDeg{
						Lat: 46,
						Lon: 14,
					},
					LatLng: s2.LatLng{
						Lat: 0.8028514559173916,
						Lng: 0.24434609527920614,
					},
				},
				DxCtyDta: spot.CtyDta{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0000,
						Lon: 14.0000,
					},
					TimeOffset: "-1.0",
				},
				DeQTH: geo.QTH{
					Loc: "JN65TW",
					LatLon: geo.LatLonDeg{
						Lat: 45.937499996666666,
						Lon: 13.625000003333334,
					},
					LatLng: s2.LatLng{
						Lat: 0.8017606250767175,
						Lng: 0.23780111064240506,
					},
				},
				DeCtyDta: spot.CtyDta{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0000,
						Lon: 14.0000,
					},
					TimeOffset: "-1.0",
				},
			},
		},

		{
			name: "Test-02",
			args: args{
				spot: spot.Data{
					Dx:       "K0RF",
					De:       "KM3T",
					Freq:     "1823.3",
					Raw:      "DX de KM3T-2-#:   1823.3  K0RF           CW    14 dB  25 WPM  CQ      1035Z",
					Comments: "CW    14 dB  25 WPM  CQ      1035Z",
					Source:   "",
					IsRbn:    true,
					BAND:     adif.Band160M,
					MODE:     adif.CW,
					DxQTH:    geo.QTH{},
					DxCtyDta: spot.CtyDta{
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
					DeQTH: geo.QTH{},
					DeCtyDta: spot.CtyDta{
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
				},
			},
			want: spot.Data{
				Dx:       "K0RF",
				De:       "KM3T",
				Freq:     "1823.3",
				Raw:      "DX de KM3T-2-#:   1823.3  K0RF           CW    14 dB  25 WPM  CQ      1035Z",
				Comments: "CW    14 dB  25 WPM  CQ      1035Z",
				Source:   "",
				IsRbn:    true,
				BAND:     adif.Band160M,
				MODE:     adif.CW,
				DxQTH: geo.QTH{
					Loc: "EN11TB",
					LatLon: geo.LatLonDeg{
						Lat: 41.08,
						Lon: -96.4,
					},
					LatLng: s2.LatLng{
						Lat: 0.716981,
						Lng: -1.682497,
					},
				},
				DxCtyDta: spot.CtyDta{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "K0",
					Continent:     adif.NA,
					CqZone:        adif.CQZONE4,
					ItuZone:       adif.ITUZONE7,
					LatLon: geo.LatLonDeg{
						Lat: 41.0800,
						Lon: -96.4000,
					},
					TimeOffset: "6.0",
				},
				DeQTH: geo.QTH{
					Loc: "FN42ET",
					LatLon: geo.LatLonDeg{
						Lat: 42.812499996666666,
						Lon: -71.62499999666667,
					},
					LatLng: s2.LatLng{
						Lat: 0.747219,
						Lng: -1.250092,
					},
				},
				DeCtyDta: spot.CtyDta{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "KM3",
					Continent:     adif.NA,
					CqZone:        adif.CQZONE5,
					ItuZone:       adif.ITUZONE8,
					LatLon: geo.LatLonDeg{
						Lat: 39.9800,
						Lon: -76.8800,
					},
					TimeOffset: "5.0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			//			if got := AddCtyDat(tt.args.spot); !reflect.DeepEqual(got, tt.want) {
			if got := AddCtyDat(tt.args.spot); !spot.DataEqual(got, tt.want) {
				t.Errorf("AddCtyDat() = %s \n ----- W A N T ----- \n %s", got.String(), tt.want.String())
			}
		})
	}
}
