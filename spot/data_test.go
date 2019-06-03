package spot

import (
	"ctestsdk/cty"
	"ctestsdk/geo"
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
				dx:       "S50A",
				de:       "S50ARX",
				freq:     "3502.8",
				raw:      "DX de S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				comments: "CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				Source:   "",
				IsRbn:    true,
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
			got, err := NewSpot(tt.args.rawData, "")
			//			got, err := NewSpot(tt.args.rawData, "data_test")
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

func TestNewSpotPoc(t *testing.T) {
	type args struct {
		rawData string
		source  string
	}
	tests := []struct {
		name    string
		args    args
		want    Data
		wantErr bool
	}{
		{
			name: "",
			args: args{
				rawData: "DX de S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				source:  "",
			},
			want: Data{
				dx:       "S50A",
				de:       "S50ARX",
				freq:     "3502.8",
				raw:      "DX de S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				comments: "CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				Source:   "",
				IsRbn:    true,
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
			got, err := NewSpotPoc(tt.args.rawData, tt.args.source)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSpotPoc() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpotPoc() = %v, want %v", got, tt.want)
			}
		})
	}
}
