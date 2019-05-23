package spot

import "testing"

func TestGetMode(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    Mode
		wantErr bool
	}{
		{
			name: "error",
			args: args{
				input: "abc",
			},
			want:    0,
			wantErr: true,
		},
		{
			name: "cw",
			args: args{
				input: "  cW  ",
			},
			want:    CW,
			wantErr: false,
		},
		{
			name: "empty",
			args: args{
				input: "",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMode(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetMode() = %v, want %v", got, tt.want)
			}
		})
	}
}
