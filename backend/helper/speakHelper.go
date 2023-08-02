package helper

var displayShortMap = map[string]string{
	"晓晓": "zh-CN-XiaoxiaoNeural",
	"晓伊": "zh-CN-XiaoyiNeural",
	"云健": "zh-CN-YunjianNeural",
	"云希": "zh-CN-YunxiNeural",
	"云夏": "zh-CN-YunxiaNeural",
	"云扬": "zh-CN-YunyangNeural",
}

// get short name by display name
func SpeakerShortName(display string) string {
	display, ok := displayShortMap[display]
	if ok {
		return display
	}
	return ""
}
