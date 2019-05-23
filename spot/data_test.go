package spot

import (
	"ctestsdk/cty"
	"ctestsdk/geo"
	"ctestsdk/testdata"
	"reflect"
	"testing"
)

func TestNewSpot(t *testing.T) {
	type args struct {
		rawData string
	}
	tests := []struct {
		name    string
		args    args
		want    Data
		wantErr bool
	}{
		{
			name: "Empty",
			args: args{
				rawData: "",
			},
			want: Data{
				dx:       "",
				de:       "",
				freq:     "",
				raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: cty.Dta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: cty.Dta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-001",
			args: args{
				rawData: "XY de",
			},
			want: Data{
				dx:       "",
				de:       "",
				freq:     "",
				raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: cty.Dta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: cty.Dta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-002",
			args: args{
				rawData: "DX de ",
			},
			want: Data{
				dx:       "",
				de:       "",
				freq:     "",
				raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: cty.Dta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: cty.Dta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-003",
			args: args{
				rawData: "DX de  S",
			},
			want: Data{
				dx:       "",
				de:       "",
				freq:     "",
				raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: cty.Dta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: cty.Dta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-004",
			args: args{
				rawData: "DX de  S59ABC",
			},
			want: Data{
				dx:       "",
				de:       "",
				freq:     "",
				raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: cty.Dta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: cty.Dta{},
			},
			wantErr: true,
		},
		{
			name: "Spot-001",
			args: args{
				rawData: "DX de S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
			},
			want: Data{
				dx:       "",
				de:       "",
				freq:     "",
				raw:      "DX de S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				DxQTH:    geo.QTH{},
				DxCtyDta: cty.Dta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: cty.Dta{},
			},
			wantErr: false,
		},

		{
			name: "Spot-002",
			args: args{
				rawData: testdata.Spots[0],
			},
			want: Data{
				dx:       "",
				de:       "",
				freq:     "",
				raw:      testdata.Spots[0],
				BAND:     0,
				MODE:     0,
				DxQTH:    geo.QTH{},
				DxCtyDta: cty.Dta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: cty.Dta{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSpot(tt.args.rawData)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSpot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpot() = %v, want %v", got, tt.want)
			}
		})
	}
}
