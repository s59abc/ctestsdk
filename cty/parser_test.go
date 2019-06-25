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
				ctyDatRecord: "Slovenia:                 15:  28:  EU: ",
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
			name: "Slovenia",
			args: args{
				ctyDatRecord: testdata.Slovenia,
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
			name: "Sweden",
			args: args{
				ctyDatRecord: testdata.Sweden,
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
			name: "AfricanItaly",
			args: args{
				ctyDatRecord: testdata.AfricanItaly,
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
			name: "Yemen",
			args: args{
				ctyDatRecord: testdata.Yemen,
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
				ctyDatRecord: testdata.Peter1Island,
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
			name: "Bouvet",
			args: args{
				ctyDatRecord: testdata.Bouvet,
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
