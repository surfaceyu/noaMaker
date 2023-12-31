<script lang="ts" setup>
import { onBeforeUnmount, reactive, ref, shallowRef, watch } from 'vue';
import { Editor} from '@wangeditor/editor-for-vue'
import { insertNewlineBeforeFourSpaces, splitStringByBr } from '@scripts/strings'
import { speakers } from '@scripts/editorHelper'
import { SpeakSample, DeleteSpeakSample, Text2Speak, Text2Speech } from "@wailsjs/go/backend/SpeakerHandler";

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

// 按钮组 分割 试听 合成
const buttonsText = [
    {
        custom: "分割",
    },
    {
        custom: "试听",
        loading: "试听中",
    },
    {
        custom: "合成",
        loading: "合成中",
    },
]
// {type, text, loading, onClick}
const buttons = reactive([
    { type: 'primary', text: buttonsText[0].custom, loading: false, onClick: handleSplit },
    { type: 'info', text: buttonsText[1].custom, loading: false, onClick: handleAudioSample },
    { type: 'success', text: buttonsText[2].custom, loading: false, onClick: handleAudioSynthesizing },
])


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
    buttons[1].loading = true
    buttons[1].text = buttonsText[1].loading!
    const sampleText = splitStringByBr(editorRef.value.getText()).filter((it, index) => it != "" && index <= 1)
    await Text2Speech(speakers[0], 0, 0, "", sampleText)
    await SpeakSample()
    buttons[1].loading = false
    buttons[1].text = buttonsText[1].custom
}

async function handleAudioSynthesizing() {
    if (!props.name || !props.chapter) {
        return
    }
    buttons[2].loading = true
    buttons[2].text = buttonsText[2].loading!
    await Text2Speak(speakers[0], props.name, props.chapter, splitStringByBr(editorRef.value.getText()))
    buttons[2].loading = false
    buttons[2].text = buttonsText[2].custom
}
</script>

<template>
    <el-row :gutter="20">
        <el-col class="mbt-10" :span="8" :offset="16">
            <el-button v-for="it in buttons" :type="it.type" size="default" :loading="it.loading" @click="it.onClick">{{ it.text }}</el-button>
            
            <!-- <el-button type="primary" size="default" @click="handleSplit">分割</el-button>
            <el-button type="info" size="default" @click="handleAudioSample">试听</el-button>
            <el-button type="success" size="default" @click="handleAudioSynthesizing">合成</el-button> -->
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