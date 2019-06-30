package cty

import (
	"ctestsdk/geo"
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
		wantCtyDatList []Dta
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
		wantDtaRecords []Dta
		wantErr        bool
	}{
		{
			name: "SloveniaCtyDat",
			args: args{
				ctyDatRecord: testdata.SloveniaCtyDat,
			},
			wantDtaRecords: []Dta{
				{
					countryName:   "Slovenia",
					primaryPrefix: "S5",
					aliasPrefix:   "S5",
					continent:     EU,
					cqZone:        CQZONE15,
					ituZone:       ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Slovenia",
					primaryPrefix: "S5",
					aliasPrefix:   "S51LGT/LH",
					continent:     EU,
					cqZone:        CQZONE15,
					ituZone:       ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Slovenia",
					primaryPrefix: "S5",
					aliasPrefix:   "S52AL/YL",
					continent:     EU,
					cqZone:        CQZONE15,
					ituZone:       ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Slovenia",
					primaryPrefix: "S5",
					aliasPrefix:   "S52L/LH",
					continent:     EU,
					cqZone:        CQZONE15,
					ituZone:       ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Slovenia",
					primaryPrefix: "S5",
					aliasPrefix:   "S58U/LH",
					continent:     EU,
					cqZone:        CQZONE15,
					ituZone:       ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Slovenia",
					primaryPrefix: "S5",
					aliasPrefix:   "S59HIJ/LH",
					continent:     EU,
					cqZone:        CQZONE15,
					ituZone:       ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "SwedenCtyDat",
			args: args{
				ctyDatRecord: testdata.SwedenCtyDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "SM",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},

				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "7S",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "8S",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "SL",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},

				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "8S8ODEN/MM",
					continent:     EU,
					cqZone:        CQZONE40,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "SK7RN/LH",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "SM7AAL/S",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "AfricanItalyCtyDat",
			args: args{
				ctyDatRecord: testdata.AfricanItalyCtyDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "African Italy",
					primaryPrefix: "IG9",
					aliasPrefix:   "IG9",
					continent:     AF,
					cqZone:        CQZONE33,
					ituZone:       ITUZONE37,
					latLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: -12.67,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "African Italy",
					primaryPrefix: "IG9",
					aliasPrefix:   "IH9",
					continent:     AF,
					cqZone:        CQZONE33,
					ituZone:       ITUZONE37,
					latLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: -12.67,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "African Italy",
					primaryPrefix: "IG9",
					aliasPrefix:   "IO9Y",
					continent:     AF,
					cqZone:        CQZONE33,
					ituZone:       ITUZONE37,
					latLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: -12.67,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "African Italy",
					primaryPrefix: "IG9",
					aliasPrefix:   "IY9A",
					continent:     AF,
					cqZone:        CQZONE33,
					ituZone:       ITUZONE37,
					latLon: geo.LatLonDeg{
						Lat: 35.67,
						Lon: -12.67,
					},
					timeOffset: "-1.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "YemenCtyDat",
			args: args{
				ctyDatRecord: testdata.YemenCtyDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "Yemen",
					primaryPrefix: "7O",
					aliasPrefix:   "7O",
					continent:     AS,
					cqZone:        CQZONE21,
					ituZone:       ITUZONE39,
					latLon: geo.LatLonDeg{
						Lat: 15.65,
						Lon: -48.12,
					},
					timeOffset: "-3.0",
				},
				{
					countryName:   "Yemen",
					primaryPrefix: "7O",
					aliasPrefix:   "7O2A",
					continent:     AS,
					cqZone:        CQZONE37,
					ituZone:       ITUZONE48,
					latLon: geo.LatLonDeg{
						Lat: 15.65,
						Lon: -48.12,
					},
					timeOffset: "-3.0",
				},
				{
					countryName:   "Yemen",
					primaryPrefix: "7O",
					aliasPrefix:   "7O6T",
					continent:     AS,
					cqZone:        CQZONE37,
					ituZone:       ITUZONE48,
					latLon: geo.LatLonDeg{
						Lat: 15.65,
						Lon: -48.12,
					},
					timeOffset: "-3.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "Peter 1 Island",
			args: args{
				ctyDatRecord: testdata.Peter1IslandCtyDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "Peter 1 Island",
					primaryPrefix: "3Y/p",
					aliasPrefix:   "3Y/p",
					continent:     SA,
					cqZone:        CQZONE12,
					ituZone:       ITUZONE72,
					latLon: geo.LatLonDeg{
						Lat: -68.77,
						Lon: 90.58,
					},
					timeOffset: "4.0",
				},
				{
					countryName:   "Peter 1 Island",
					primaryPrefix: "3Y/p",
					aliasPrefix:   "3Y0X",
					continent:     SA,
					cqZone:        CQZONE12,
					ituZone:       ITUZONE72,
					latLon: geo.LatLonDeg{
						Lat: -68.77,
						Lon: 90.58,
					},
					timeOffset: "4.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "BouvetCtyDat",
			args: args{
				ctyDatRecord: testdata.BouvetCtyDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "Bouvet",
					primaryPrefix: "3Y/b",
					aliasPrefix:   "3Y/b",
					continent:     AF,
					cqZone:        CQZONE38,
					ituZone:       ITUZONE67,
					latLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: -3.38,
					},
					timeOffset: "-1.0",
				},

				{
					countryName:   "Bouvet",
					primaryPrefix: "3Y/b",
					aliasPrefix:   "3Y/ZS6GCM",
					continent:     AF,
					cqZone:        CQZONE38,
					ituZone:       ITUZONE67,
					latLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: -3.38,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Bouvet",
					primaryPrefix: "3Y/b",
					aliasPrefix:   "3Y0C",
					continent:     AF,
					cqZone:        CQZONE38,
					ituZone:       ITUZONE67,
					latLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: -3.38,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Bouvet",
					primaryPrefix: "3Y/b",
					aliasPrefix:   "3Y0E",
					continent:     AF,
					cqZone:        CQZONE38,
					ituZone:       ITUZONE67,
					latLon: geo.LatLonDeg{
						Lat: -54.42,
						Lon: -3.38,
					},
					timeOffset: "-1.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "Slovenia Cty Wt Mod",
			args: args{
				ctyDatRecord: testdata.SloveniaWtModDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "Slovenia",
					primaryPrefix: "S5",
					aliasPrefix:   "S5",
					continent:     EU,
					cqZone:        CQZONE15,
					ituZone:       ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.00,
						Lon: -14.00,
					},
					timeOffset: "-1.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "Sweden Cty Wt Mod",
			args: args{
				ctyDatRecord: testdata.SwedenWtModDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "SM",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},

				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "7S",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 61.20,
						Lon: -14.57,
					},
					timeOffset: "-1.0",
				},
				{
					countryName:   "Sweden",
					primaryPrefix: "SM",
					aliasPrefix:   "SM7",
					continent:     EU,
					cqZone:        CQZONE14,
					ituZone:       ITUZONE18,
					latLon: geo.LatLonDeg{
						Lat: 55.58,
						Lon: -13.10,
					},
					timeOffset: "-1.0",
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
		wantCtyDatList []Dta
		wantErr        bool
	}{
		{
			name: "USA Cty Wt Mod",
			args: args{
				ctyDatRecord: testdata.UnitedStatesWtModDat,
			},
			wantCtyDatList: []Dta{
				{
					countryName:   "United States",
					primaryPrefix: "K",
					aliasPrefix:   "K",
					continent:     NA,
					cqZone:        CQZONE5,
					ituZone:       ITUZONE8,
					latLon: geo.LatLonDeg{
						Lat: 37.53,
						Lon: 91.67,
					},
					timeOffset: "5.0",
				},

				{
					countryName:   "United States",
					primaryPrefix: "K",
					aliasPrefix:   "AA",
					continent:     NA,
					cqZone:        CQZONE5,
					ituZone:       ITUZONE8,
					latLon: geo.LatLonDeg{
						Lat: 37.53,
						Lon: 91.67,
					},
					timeOffset: "5.0",
				},

				{
					countryName:   "United States",
					primaryPrefix: "K",
					aliasPrefix:   "K5ZD",
					continent:     NA,
					cqZone:        CQZONE5,
					ituZone:       ITUZONE8,
					latLon: geo.LatLonDeg{
						Lat: 42.27,
						Lon: 71.37,
					},
					timeOffset: "5.0",
				},

				{
					countryName:   "United States",
					primaryPrefix: "K",
					aliasPrefix:   "AD1C",
					continent:     NA,
					cqZone:        CQZONE4,
					ituZone:       ITUZONE7,
					latLon: geo.LatLonDeg{
						Lat: 39.52,
						Lon: 105.20,
					},
					timeOffset: "6.0",
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

func Test_parseCtyDatRecordsOld(t *testing.T) {
	type args struct {
		ctyDatRecords string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CtyDatRecordsOld",
			args: args{
				ctyDatRecords: testdata.CtyDatRecords,
			},
			wantErr: false,
		},
		{
			name: "CtyWtModDatRecordsOld",
			args: args{
				ctyDatRecords: testdata.CtyWtModDatRecords,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseCtyDatRecordsOld(tt.args.ctyDatRecords); (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecordsOld() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseCtyDatRecords(t *testing.T) {
	type args struct {
		ctyDatRecords string
	}
	tests := []struct {
		name     string
		args     args
		wantSize int
		wantErr  bool
	}{
		{
			name: "CtyDatRecords",
			args: args{
				ctyDatRecords: testdata.CtyDatRecords,
			},
			wantSize: 21314,
			wantErr:  false,
		},
		{
			name: "CtyWtModDatRecords",
			args: args{
				ctyDatRecords: testdata.CtyWtModDatRecords,
			},
			wantSize: 26991,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSize, err := parseCtyDatRecords(tt.args.ctyDatRecords)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSize != tt.wantSize {
				t.Errorf("parseCtyDatRecords() = %v, want %v", gotSize, tt.wantSize)
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
			gotMsize, err := parseCtyDatRecordsGo(tt.args.ctyDatRecords)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCtyDatRecordsGo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotMsize != tt.wantMsize {
				t.Errorf("parseCtyDatRecordsGo() = %v, want %v", gotMsize, tt.wantMsize)
			}
		})
	}
}
