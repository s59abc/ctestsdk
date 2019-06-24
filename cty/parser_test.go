package cty

import (
	"ctestsdk/geo"
	"ctestsdk/testdata"
	"reflect"
	"testing"
)

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
			name: "empty record",
			args: args{
				ctyDatRecord: "",
			},
			wantDtaRecords: nil,
			wantErr:        true,
		},
		{
			name: "too short record",
			args: args{
				ctyDatRecord: "Slovenia:                 15:  28:  EU: ",
			},
			wantDtaRecords: nil,
			wantErr:        true,
		},
		{
			name: "Slovenia",
			args: args{
				ctyDatRecord: testdata.Slovenia,
			},
			wantDtaRecords: []Dta{
				{
					countryName: "Slovenia",
					prefix:      "S5",
					continent:   EU,
					cqZone:      CQZONE15,
					ituZone:     ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName: "Slovenia",
					prefix:      "S51LGT/LH",
					continent:   EU,
					cqZone:      CQZONE15,
					ituZone:     ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName: "Slovenia",
					prefix:      "S52AL/YL",
					continent:   EU,
					cqZone:      CQZONE15,
					ituZone:     ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName: "Slovenia",
					prefix:      "S52L/LH",
					continent:   EU,
					cqZone:      CQZONE15,
					ituZone:     ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName: "Slovenia",
					prefix:      "S58U/LH",
					continent:   EU,
					cqZone:      CQZONE15,
					ituZone:     ITUZONE28,
					latLon: geo.LatLonDeg{
						Lat: 46.0,
						Lon: -14.0,
					},
					timeOffset: "-1.0",
				},
				{
					countryName: "Slovenia",
					prefix:      "S59HIJ/LH",
					continent:   EU,
					cqZone:      CQZONE15,
					ituZone:     ITUZONE28,
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
