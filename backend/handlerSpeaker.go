package backend

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/samber/lo"
	"github.com/surfaceyu/edge-tts-go/edgeTTS"
	"github.com/surfaceyu/noaMaker/backend/helper"
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

func (sp *SpeakerHandler) SpeakSample() {
	f, err := os.Open("./out/tmp.mp3")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	var wg sync.WaitGroup
	wg.Add(1)
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		wg.Done()
	})))

	wg.Wait()
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

func (sp *SpeakerHandler) Text2Speech(voice string, rate int, vol int, path string, contents []string) {
	args := edgeTTS.Args{
		Voice:      helper.SpeakerShortName(voice),
		WriteMedia: lo.If(path != "", path).Else("./out/tmp.mp3"),
		Rate:       lo.If(rate < 0, fmt.Sprintf("%d%%", rate)).Else(fmt.Sprintf("+%d%%", rate)),
		Volume:     lo.If(vol < 0, fmt.Sprintf("%d%%", vol)).Else(fmt.Sprintf("+%d%%", vol)),
	}
	tts := edgeTTS.NewTTS(args)
	for _, v := range contents {
		speaker, text := sp.splitSpeaker(v)
		tts.AddTextWithVoice(text, speaker)
	}
	tts.Speak()
}
