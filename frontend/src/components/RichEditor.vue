<script lang="ts" setup>
import { onBeforeUnmount, ref, shallowRef, watch } from 'vue';
import { Editor} from '@wangeditor/editor-for-vue'
import { insertNewlineBeforeFourSpaces, splitStringByBr } from '@scripts/strings'
import { speakers } from '@scripts/editorHelper'
import { SpeakSample, DeleteSpeakSample, Text2Speak } from "@wailsjs/go/backend/SpeakerHandler";

const props = defineProps({
    name: String,
    chapter: String,
    data: String,
})

const editorRef = shallowRef()
const sampleAudio = ref<HTMLAudioElement>()

const editorConfig = {
    placeholder: '暂无内容',
    hoverbarKeys: {
        'text': {
            // 清空 text 元素的 hoverbar
            menuKeys: [...speakers],
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
    MENU_CONF: {}
}

const handleCreated = (editor: any) => {
    editorRef.value = editor // 记录 editor 实例，重要！
}

const editContent = ref(props.data)
watch(props, (data) => {
    editContent.value = data.data
})

onBeforeUnmount(() => {
    const editor = editorRef.value
    if (editor == null) return
    editor.destroy()
})

function handleSplit() {
    if (editContent.value) {
        editContent.value = insertNewlineBeforeFourSpaces(editContent.value)
    }
}

async function handleAudioSample() {
    console.log(splitStringByBr(editorRef.value.getSelectionText()))
    // await SpeakSample(speakers[0], splitStringByBr(editorRef.value.getSelectionText()))
    // // sampleAudio 类型是HTMLAudioElement
    // if (sampleAudio.value) {        
    //     // 停止播放
    //     sampleAudio.value.pause(); 
    //     sampleAudio.value.currentTime = 0; 
    // }
    // // sampleAudio 加载"../../../sample.mp3" 并播放
    // sampleAudio.value = new Audio("../assets/sample.mp3"); 
    // sampleAudio.value.play(); 
}

async function handleAudioSynthesizing() {
    // console.log(splitStringByBr(editorRef.value.getText()));
    if (!props.name || !props.chapter) {
        return
    }
    await Text2Speak(speakers[0], props.name, props.chapter, splitStringByBr(editorRef.value.getText()))
}
</script>

<template>
    <el-row :gutter="20">
        <el-col class="mbt-10" :span="8" :offset="16">
            <el-button type="primary" size="default" @click="handleSplit">分割</el-button>
            <el-button type="info" size="default" @click="handleAudioSample">试听</el-button>
            <el-button type="success" size="default" @click="handleAudioSynthesizing">合成</el-button>
        </el-col>
    </el-row>

    <div style="border: 1px solid #ccc">
        <!-- <Toolbar style="border-bottom: 1px solid #ccc" :editor="editorRef" :defaultConfig="toolbarConfig" mode="default" /> -->
        <Editor style="height: 400px; overflow-y: hidden; text-align: left;" v-model="editContent"
            :defaultConfig="editorConfig" mode="default" @onCreated="handleCreated" />
    </div>
</template>

<style scoped>
.mbt-10 {
    margin-top: 10px;
    margin-bottom: 10px;
}
</style>