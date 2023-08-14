package backend

import (
	"testing"
)

func TestSpeakerHandler_splitSpeaker(t *testing.T) {
	type args struct {
		content string
	}
	tests := []struct {
		name  string
		g     *SpeakerHandler
		args  args
		want  string
		want1 string
	}{
		{
			name: "test-1",
			g:    &SpeakerHandler{},
			args: args{
				content: "[[晓晓]]注：食量参考来源是家里参加过知青下乡的长辈。",
			},
			want:  "zh-CN-XiaoxiaoNeural",
			want1: "注：食量参考来源是家里参加过知青下乡的长辈。",
		},
		{
			name: "test-2",
			g:    &SpeakerHandler{},
			args: args{
				content: "[[晓晓2]]注：食量参考来源是家里参加过知青下乡的长辈。",
			},
			want:  "",
			want1: "注：食量参考来源是家里参加过知青下乡的长辈。",
		},
		{
			name: "test-3",
			g:    &SpeakerHandler{},
			args: args{
				content: "注：食量参考来源是家里参加过知青下乡的长辈。",
			},
			want:  "",
			want1: "注：食量参考来源是家里参加过知青下乡的长辈。",
		},
		{
			name: "test-4",
			g:    &SpeakerHandler{},
			args: args{
				content: "晓晓2]]注：食量参考来源是家里参加过知青下乡的长辈。",
			},
			want:  "",
			want1: "晓晓2]]注：食量参考来源是家里参加过知青下乡的长辈。",
		},
		{
			name: "test-5",
			g:    &SpeakerHandler{},
			args: args{
				content: "[[晓晓2注：食量参考来源是家里参加过知青下乡的长辈。",
			},
			want:  "",
			want1: "[[晓晓2注：食量参考来源是家里参加过知青下乡的长辈。",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.g.splitSpeaker(tt.args.content)
			if got != tt.want {
				t.Errorf("SpeakerHandler.splitSpeaker() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SpeakerHandler.splitSpeaker() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSpeakerHandler_SpeakSample(t *testing.T) {
	tests := []struct {
		name string
		sp   *SpeakerHandler
	}{
		{
			name: "test-1",
			sp:   &SpeakerHandler{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sp.SpeakSample()
		})
	}
}
