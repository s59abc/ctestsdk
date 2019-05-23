package freq

import "testing"

func TestGetBand(t *testing.T) {
	type args struct {
		kHz string
	}
	tests := []struct {
		name    string
		args    args
		want    Band
		wantErr bool
	}{
		{
			name: "160M-1",
			args: args{
				kHz: "1800.0",
			},
			want:    M160,
			wantErr: false,
		},
		{
			name: "80M-1",
			args: args{
				kHz: "3500.0",
			},
			want:    M80,
			wantErr: false,
		},
		{
			name: "40M-1",
			args: args{
				kHz: "   7000.0   ",
			},
			want:    M40,
			wantErr: false,
		},
		{
			name: "20M-1",
			args: args{
				kHz: "14000.0   ",
			},
			want:    M20,
			wantErr: false,
		},
		{
			name: "15M-1",
			args: args{
				kHz: " 21000.0",
			},
			want:    M15,
			wantErr: false,
		},
		{
			name: "10M-1",
			args: args{
				kHz: "28000.0",
			},
			want:    M10,
			wantErr: false,
		},

		{
			name: "error-1",
			args: args{
				kHz: "",
			},
			want:    UNKNOWN,
			wantErr: true,
		},

		{
			name: "error-2",
			args: args{
				kHz: "3500,0",
			},
			want:    UNKNOWN,
			wantErr: true,
		},
		{
			name: "error-3",
			args: args{
				kHz: "3.500.0",
			},
			want:    UNKNOWN,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBand(tt.args.kHz)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetBand() = %v, want %v", got, tt.want)
			}
		})
	}
}
