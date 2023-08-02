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
	type args struct {
		voice    string
		contents []string
	}
	tests := []struct {
		name string
		sp   *SpeakerHandler
		args args
	}{
		{
			name: "test-1",
			sp:   &SpeakerHandler{},
			args: args{
				voice: "晓晓",
				contents: []string{
					"第一卷、剑之起源 第一章 我想要回去",
					"光明纪元8888年。",
					"东道真洲，北海帝国，风语行省。",
					"云梦城，省立第三初级剑士学院。",
					"风和日丽，万物生长 ，微风习习。",
					"初夏，正值云梦城一年之内气温最舒适的时节。",
					"[[云健]]“不能生气，不能生气。”",
					"[[云健]]“他脑子有病，我又没病。”",
					`[[云健]]“不能和一个有脑疾的纨绔一般见识。”`,
					`以脾气火爆闻名第三初级学院资深教习丁三石，在心中一遍遍地默念，努力让自己无视这个全城最大的纨绔败家子，继续上课。`,
					`其他学员看到‘暴躁教习’丁三石如此‘忍气吞声’的样子，不由得都忍俊不禁，想笑也不敢笑出声。`,
					`只是老师和同学们不知道的是，林北辰并不是脑疾犯了在‘看手’。`,
					`而是在看手机。`,
					`一部除了林北辰自己之外，其他人绝对看不见的智能手机。`,

					`[[云扬]]“坑爹啊。”`,
					`此时的林北辰，正在心里嗷嚎。`,
					`他这是造了什么孽啊。`,
					`只不过是在马路上拉住了一个闯红灯差点儿被大卡车撞死的神经病，结果就被这个自称是‘死神’的家伙，强行塞了这部没有品牌LOGO的智能手机，再然后他就……灵魂穿越了！`,
					`来到了这个叫做东道真洲的奇怪世界。`,
					`成为了北海帝国‘十大名将’之一的‘战天侯’林近南的嫡子。`,
					`[[云健]]“好了，刚才已经为大家解析了【初等玄气凝练术】的完整版，现在我们休息一刻钟，然后继续上课。”`,
					`老教习丁三石端起杯子喝了一口水，润了润嗓子。`,

					`[[云健]]“大家都知道，三日之后，就是咱们学院的年中大比了，这次大比的重要性，不用我再多强调吧？好了，提前预告一下，下一节课，我为大家准备了一节精选课程，是我的独门秘技【基础剑术近身三连】。”`,
					`老教习丁三石说着，目光又看到了林北辰。`,
					`见这个纨绔子弟，还是一副魂游天外的样子，他不由得失望地摇摇头。`,

					`[[云健]]“林北辰，下节课你要好好听，这一套【基础剑术近身三连】，最适合你这样基础差的学员。”`,
					`丁三时忍不住特意多说了一句。`,
					`然而林北辰依旧是木若呆鸡毫无反应。`,
					`老教习一脸无语地转身出了教室。`,
					`对于老教习的怨念，林北辰丝毫不以为意。`,
					`他对这个世界，缺乏认同感。`,
					`对于自己的新身份，也是毫无代入感。`,
					`现在的他，一门心思想的，是如何回到原来的世界中去。`,
					`因此对于什么狗屁年中大比，什么晋升和前途，什么【初等玄气凝练术】、什么【基础剑术近身三连】，都特么一边玩蛋去吧。`,
					"他继续默默地研究手机。",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sp.SpeakSample(tt.args.voice, tt.args.contents)
		})
	}
}
