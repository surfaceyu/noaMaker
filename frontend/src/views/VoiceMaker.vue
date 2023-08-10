<script lang="ts" setup>
import { reactive, ref, shallowRef } from 'vue';
import { Editor} from '@wangeditor/editor-for-vue'

import { speakers } from '@scripts/editorHelper'
import { Text2Speech } from "@wailsjs/go/backend/SpeakerHandler";

const speakersLabels = ['晓晓(女声)', '晓伊(女童声)', '云健(男声)', '云希(男声)', '云夏(男童声)', '云扬(男声)']

// 按钮组 分割 试听 合成
const buttonsText = [
    {
        custom: "试听",
        loading: "试听中",
    },
    {
        custom: "合成",
        loading: "合成中",
    },
    {
        custom: "查看",
    },
]
// {type, text, loading, onClick}
const buttons = reactive([
    { type: 'warning', text: buttonsText[0].custom, loading: false, onClick: null },
    { type: 'info', text: buttonsText[1].custom, loading: false, onClick: handleAudioSynthesizing },
    { type: 'success', text: buttonsText[2].custom, loading: false, onClick: onLookUp },
])

const editorConfig = {
    placeholder: '暂无内容',
    hoverbarKeys: {
        'text': {
            // 清空 text 元素的 hoverbar
            menuKeys: [],
        },
        'image': {
            // 清空 image 元素的 hoverbar
            menuKeys: [],
        },
        'link': {
            // 清空 link 元素的 hoverbar
            menuKeys: [],
        },
    },
}

const voice = ref(speakers[0])
const voiceRate = ref(0)
const voiceVolume = ref(0)
const editorRef = shallowRef()
const editContent = ref("")
const result = ref("Default")

function transVoiceNum(num: number) {
    return num >= 0 ? `+${num}%` : `${num}%`
}

function handleCreated(editor: any) {
    editorRef.value = editor // 记录 editor 实例，重要！
}

async function handleAudioSynthesizing() {
    buttons[1].loading = true
    buttons[1].text = buttonsText[1].loading!
    await Text2Speech(voice.value, voiceRate.value, voiceVolume.value, "", ["你好 你好 你好 你好"])
    buttons[1].loading = false
    buttons[1].text = buttonsText[1].custom
}

function onLookUp() {    
    result.value = `${voice.value}  ${transVoiceNum(voiceRate.value)}  ${transVoiceNum(voiceVolume.value)}`
}

</script>

<template>
    <h5 class="mb-2 h5-title">音频制作</h5>
    <div class="radius-border border-width" :style="{ borderRadius: 'var(--el-border-radius-small)' }">
        <el-row class="mt-2">
            <el-col :span="16">
                <el-row :gutter="3" class="pl-1">
                    <el-col :span="5">
                        <el-select v-model="voice" placeholder="Select">
                            <el-option class="ta-l" v-for="(item, index) in speakers" :key="item" :label="speakersLabels[index]"
                                :value="item" />
                        </el-select>
                    </el-col>
                    <el-col :span="10">
                        <div class="slider-demo-block">
                            <span class="demonstration">语速 {{ voiceRate }}</span>
                            <el-slider v-model="voiceRate" :min=-100 :max=200 :step=5 />
                        </div>
                    </el-col>
                    <el-col :span="9">
                        <div class="slider-demo-block">
                            <span class="demonstration">音调 {{ voiceVolume }}</span>
                            <el-slider v-model="voiceVolume" :min=-50 :max=50 />
                        </div>
                    </el-col>
                </el-row>

            </el-col>
            <el-col :span="8">
                <el-button-group class="ml-4">
                    <el-button v-for="it in buttons" :type="it.type" :loading="it.loading"
                        @click="it.onClick">{{
                            it.text }}</el-button>
                </el-button-group>
            </el-col>
        </el-row>
    </div>
    <div class="mt-4 border-width" style="border: 1px solid #ccc">
        <Editor style="height: 400px; overflow-y: hidden; text-align: left;" v-model="editContent"
            :defaultConfig="editorConfig" mode="default" @onCreated="handleCreated" />
    </div>
</template>

<style scoped>
.h5-title {
    text-align: left;
    padding-left: 20px;
    font-size: large;
    color: #e7ac08;
}

.radius-border {
    height: 3rem;
    border: 1px solid var(--el-border-color);
    border-radius: 0;
    margin-top: 20px;
}

.border-width {
    width: 90%;
    margin-left: 5%;
}

.slider-demo-block {
    display: flex;
    align-items: center;
}

.slider-demo-block .el-slider {
    margin-top: 0;
    margin-left: 12px;
}

.slider-demo-block .demonstration {
    font-size: 14px;
    color: var(--el-text-color-secondary);
    line-height: 44px;
    flex: 1;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    margin-bottom: 0;
}

.slider-demo-block .demonstration+.el-slider {
    flex: 0 0 70%;
}
</style>
