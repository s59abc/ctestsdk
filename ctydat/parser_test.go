package ctydat

import (
	"ctestsdk/adif"
	"ctestsdk/geo"
	"ctestsdk/spot"
	"ctestsdk/testdata"
	"reflect"
	"testing"
)

func Test_parseCtyDatRecordErrorCases(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "emptyRecord",
			args: args{
				ctyDatRecord: "",
			},
			wantCtyDatList: nil,
			wantErr:        true,
		},
		{
			name: "tooShorRecord",
			args: args{
				ctyDatRecord: "SloveniaCtyDat:                 15:  28:  EU: ",
			},
			wantCtyDatList: nil,
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordSlovenia(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantDtaRecords []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "SloveniaCtyDat",
			args: args{
				ctyDatRecord: testdata.SloveniaCtyDat,
			},
			wantDtaRecords: []spot.CtyDta{
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S51LGT/LH",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S52AL/YL",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S52L/LH",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S58U/LH",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S59HIJ/LH",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: 14.0,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDtaRecords, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotDtaRecords, tt.wantDtaRecords) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotDtaRecords, tt.wantDtaRecords)
			}
		})
	}
}

func Test_parseCtyDatRecordSweden(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "SwedenCtyDat",
			args: args{
				ctyDatRecord: testdata.SwedenCtyDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "7S",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "8S",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SL",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "8S8ODEN/MM",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE40,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SK7RN/LH",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM7AAL/S",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
			},

			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordAfricanItaly(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "AfricanItalyCtyDat",
			args: args{
				ctyDatRecord: testdata.AfricanItalyCtyDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IG9",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE33,
					ItuZone:       adif.ITUZONE37,
					LatLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IH9",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE33,
					ItuZone:       adif.ITUZONE37,
					LatLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IO9Y",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE33,
					ItuZone:       adif.ITUZONE37,
					LatLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "African Italy",
					PrimaryPrefix: "IG9",
					AliasPrefix:   "IY9A",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE33,
					ItuZone:       adif.ITUZONE37,
					LatLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: 12.67,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordYemen(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "YemenCtyDat",
			args: args{
				ctyDatRecord: testdata.YemenCtyDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "Yemen",
					PrimaryPrefix: "7O",
					AliasPrefix:   "7O",
					Continent:     adif.AS,
					CqZone:        adif.CQZONE21,
					ItuZone:       adif.ITUZONE39,
					LatLon: geo.LatLonDeg{
						Lat: 15.65,
						Lon: 48.12,
					},
					TimeOffset: "-3.0",
				},
				{
					CountryName:   "Yemen",
					PrimaryPrefix: "7O",
					AliasPrefix:   "7O2A",
					Continent:     adif.AS,
					CqZone:        adif.CQZONE37,
					ItuZone:       adif.ITUZONE48,
					LatLon: geo.LatLonDeg{
						Lat: 15.65,
						Lon: 48.12,
					},
					TimeOffset: "-3.0",
				},
				{
					CountryName:   "Yemen",
					PrimaryPrefix: "7O",
					AliasPrefix:   "7O6T",
					Continent:     adif.AS,
					CqZone:        adif.CQZONE37,
					ItuZone:       adif.ITUZONE48,
					LatLon: geo.LatLonDeg{
						Lat: 15.65,
						Lon: 48.12,
					},
					TimeOffset: "-3.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordPeter1Island(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "Peter 1 Island",
			args: args{
				ctyDatRecord: testdata.Peter1IslandCtyDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "Peter 1 Island",
					PrimaryPrefix: "3Y/p",
					AliasPrefix:   "3Y/p",
					Continent:     adif.SA,
					CqZone:        adif.CQZONE12,
					ItuZone:       adif.ITUZONE72,
					LatLon: geo.LatLonDeg{
						Lat: -68.77,
						Lon: -90.58,
					},
					TimeOffset: "4.0",
				},
				{
					CountryName:   "Peter 1 Island",
					PrimaryPrefix: "3Y/p",
					AliasPrefix:   "3Y0X",
					Continent:     adif.SA,
					CqZone:        adif.CQZONE12,
					ItuZone:       adif.ITUZONE72,
					LatLon: geo.LatLonDeg{
						Lat: -68.77,
						Lon: -90.58,
					},
					TimeOffset: "4.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordBouvet(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "BouvetCtyDat",
			args: args{
				ctyDatRecord: testdata.BouvetCtyDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y/b",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE38,
					ItuZone:       adif.ITUZONE67,
					LatLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y/ZS6GCM",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE38,
					ItuZone:       adif.ITUZONE67,
					LatLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y0C",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE38,
					ItuZone:       adif.ITUZONE67,
					LatLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Bouvet",
					PrimaryPrefix: "3Y/b",
					AliasPrefix:   "3Y0E",
					Continent:     adif.AF,
					CqZone:        adif.CQZONE38,
					ItuZone:       adif.ITUZONE67,
					LatLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: 3.38,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}
func Test_parseCtyDatRecordSloveniaCtyWtDat(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "Slovenia Cty Wt Mod",
			args: args{
				ctyDatRecord: testdata.SloveniaWtModDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "Slovenia",
					PrimaryPrefix: "S5",
					AliasPrefix:   "S5",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE15,
					ItuZone:       adif.ITUZONE28,
					LatLon: geo.LatLonDeg{
						Lat: 46.00,
						Lon: 14.00,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordSwedenCtyWtDat(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "Sweden Cty Wt Mod",
			args: args{
				ctyDatRecord: testdata.SwedenWtModDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},

				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "7S",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: 14.57,
					},
					TimeOffset: "-1.0",
				},
				{
					CountryName:   "Sweden",
					PrimaryPrefix: "SM",
					AliasPrefix:   "SM7",
					Continent:     adif.EU,
					CqZone:        adif.CQZONE14,
					ItuZone:       adif.ITUZONE18,
					LatLon: geo.LatLonDeg{
						Lat: 55.58,
						Lon: 13.10,
					},
					TimeOffset: "-1.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_parseCtyDatRecordUnitedStatesCtyWtDat(t *testing.T) {
	type args struct {
		ctyDatRecord string
	}
	tests := []struct {
		name           string
		args           args
		wantCtyDatList []spot.CtyDta
		wantErr        bool
	}{
		{
			name: "USA Cty Wt Mod",
			args: args{
				ctyDatRecord: testdata.UnitedStatesWtModDat,
			},
			wantCtyDatList: []spot.CtyDta{
				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "K",
					Continent:     adif.NA,
					CqZone:        adif.CQZONE5,
					ItuZone:       adif.ITUZONE8,
					LatLon: geo.LatLonDeg{
						Lat: 37.53,
						Lon: -91.67,
					},
					TimeOffset: "5.0",
				},

				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "AA",
					Continent:     adif.NA,
					CqZone:        adif.CQZONE5,
					ItuZone:       adif.ITUZONE8,
					LatLon: geo.LatLonDeg{
						Lat: 37.53,
						Lon: -91.67,
					},
					TimeOffset: "5.0",
				},

				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "K5ZD",
					Continent:     adif.NA,
					CqZone:        adif.CQZONE5,
					ItuZone:       adif.ITUZONE8,
					LatLon: geo.LatLonDeg{
						Lat: 42.27,
						Lon: -71.37,
					},
					TimeOffset: "5.0",
				},

				{
					CountryName:   "United States",
					PrimaryPrefix: "K",
					AliasPrefix:   "AD1C",
					Continent:     adif.NA,
					CqZone:        adif.CQZONE4,
					ItuZone:       adif.ITUZONE7,
					LatLon: geo.LatLonDeg{
						Lat: 39.52,
						Lon: -105.20,
					},
					TimeOffset: "6.0",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCtyDatList, err := parseCtyDatRecord(tt.args.ctyDatRecord)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCtyDatList, tt.wantCtyDatList) {
				t.Errorf("parseCtyDatRecord() = %v, want %v", gotCtyDatList, tt.wantCtyDatList)
			}
		})
	}
}

func Test_removeComments(t *testing.T) {
	type args struct {
		ctyDatRecords string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "comments-1",
			args: args{
				ctyDatRecords: testdata.TestInput1,
			},
			want: testdata.TestOutput1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeComments(tt.args.ctyDatRecords); got != tt.want {
				t.Errorf("removeComments() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCtyDatRecordsGo(t *testing.T) {
	type args struct {
		ctyDatRecords string
	}
	tests := []struct {
		name      string
		args      args
		wantMsize int
		wantErr   bool
	}{
		{
			name: "CtyDatRecords",
			args: args{
				ctyDatRecords: testdata.CtyDatRecords,
			},
			wantMsize: 21314,
			wantErr:   false,
		},
		{
			name: "CtyWtModDatRecords",
			args: args{
				ctyDatRecords: testdata.CtyWtModDatRecords,
			},
			wantMsize: 26991,
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotMsize, err := parseCtyDatRecordsForTest(tt.args.ctyDatRecords)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecordsForTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMsize != tt.wantMsize {
				t.Errorf("parseCtyDatRecordsForTest() = %v, want %v", gotMsize, tt.wantMsize)
			}
		})
	}
}
