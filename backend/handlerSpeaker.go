package backend

import (
	"fmt"
	"os"
	"strings"

	"github.com/surfaceyu/edge-tts-go/edgeTTS"
	"github.com/surfaceyu/myproject/backend/helper"
)

func init() {
	bindHandlers = append(bindHandlers, &SpeakerHandler{})
}

type SpeakerHandler struct {
}

// 获取[[]]中间的字符串 并分割
// 如 [[晓晓]]注：食量参考来源是家里参加过知青下乡的长辈。
// 返回 晓晓 和 注：食量参考来源是家里参加过知青下乡的长辈。
func (sp *SpeakerHandler) splitSpeaker(content string) (string, string) {
	startIndex := strings.Index(content, "[[")
	endIndex := strings.Index(content, "]]")
	if startIndex == -1 || endIndex == -1 {
		return "", content
	}
	speaker := helper.SpeakerShortName(content[startIndex+2 : endIndex])
	annotation := content[endIndex+2:]
	return speaker, annotation
}

func (sp *SpeakerHandler) SpeakSample(voice string, contents []string) {
	args := edgeTTS.Args{
		Voice:      helper.SpeakerShortName(voice),
		WriteMedia: "./sample.mp3",
	}
	tts := edgeTTS.NewTTS(args)
	for _, v := range contents {
		speaker, text := sp.splitSpeaker(v)
		fmt.Println(speaker, text)
		tts.AddTextWithVoice(text, speaker)
	}
	tts.Speak()
}

// 删除试听的mp3文件 "./sample.mp3"
func (sp *SpeakerHandler) DeleteSpeakSample() {
	_ = os.Remove("./sample.mp3")
}

func (sp *SpeakerHandler) Text2Speak(voice string, name string, title string, contents []string) {
	args := edgeTTS.Args{
		Voice:      helper.SpeakerShortName(voice),
		WriteMedia: fmt.Sprintf("./out/%s/%s.mp3", name, title),
	}
	tts := edgeTTS.NewTTS(args)
	for _, v := range contents {
		speaker, text := sp.splitSpeaker(v)
		tts.AddTextWithVoice(text, speaker)
	}
	tts.Speak()
}
