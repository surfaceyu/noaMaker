package helper

import "testing"

func TestSpeakerShortName(t *testing.T) {
	type args struct {
		display string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-1",
			args: args{
				display: "晓晓",
			},
			want: "zh-CN-XiaoxiaoNeural",
		},
		{
			name: "test-2",
			args: args{
				display: "晓1晓",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpeakerShortName(tt.args.display); got != tt.want {
				t.Errorf("SpeakerShortName() = %v, want %v", got, tt.want)
			}
		})
	}
}
