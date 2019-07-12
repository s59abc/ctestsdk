package spot

import (
	"ctestsdk/adif"
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
				Dx:       "",
				De:       "",
				Freq:     "",
				Raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-001",
			args: args{
				rawData: "XY De",
			},
			want: Data{
				Dx:       "",
				De:       "",
				Freq:     "",
				Raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-002",
			args: args{
				rawData: "DX De ",
			},
			want: Data{
				Dx:       "",
				De:       "",
				Freq:     "",
				Raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-003",
			args: args{
				rawData: "DX De  S",
			},
			want: Data{
				Dx:       "",
				De:       "",
				Freq:     "",
				Raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
			},
			wantErr: true,
		},
		{
			name: "Not-a-Spot-004",
			args: args{
				rawData: "DX De  S59ABC",
			},
			want: Data{
				Dx:       "",
				De:       "",
				Freq:     "",
				Raw:      "",
				DxQTH:    geo.QTH{},
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
			},
			wantErr: true,
		},
		{
			name: "Spot-001",
			args: args{
				rawData: "DX De S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
			},
			want: Data{
				Dx:       "S50A",
				De:       "S50ARX",
				Freq:     "3502.8",
				Raw:      "DX De S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				Comments: "CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				Source:   "",
				IsRbn:    true,
				BAND:     adif.Band80M,
				MODE:     adif.CW,
				DxQTH:    geo.QTH{},
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
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
			name: "POC-1",
			args: args{
				rawData: "DX De S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				source:  "",
			},
			want: Data{
				Dx:       "S50A",
				De:       "S50ARX",
				Freq:     "3502.8",
				Raw:      "DX De S50ARX-#:     3502.8  S50A        CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				Comments: "CW    14 dB  26 WPM  CQ       1118Z IO73\a\a",
				Source:   "",
				IsRbn:    true,
				BAND:     adif.Band80M,
				MODE:     adif.CW,
				DxQTH:    geo.QTH{},
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
			},
			wantErr: false,
		},

		{
			name: "POC-2",
			args: args{
				rawData: "DX de KM3T-2-#:   1823.3  K0RF           CW    14 dB  25 WPM  CQ      1035Z",
				source:  "",
			},
			want: Data{
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
				DxCtyDta: CtyDta{},
				DeQTH:    geo.QTH{},
				DeCtyDta: CtyDta{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSpot(tt.args.rawData, tt.args.source)
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
