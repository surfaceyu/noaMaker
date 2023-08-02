import { Boot, IDomEditor, IButtonMenu } from '@wangeditor/editor'
import { DomEditor } from '@wangeditor/core'
import { log } from 'console'

function withParagraph<T extends IDomEditor>(editor: T): T { 
    const { insertBreak } = editor // 获取当前 editor API
    const newEditor = editor
    // 重写 insertBreak 换行
    newEditor.insertBreak = () => {
        const selectedNode = DomEditor.getSelectedNodeByType(newEditor, 'paragraph')
        if (selectedNode != null) {
            // 选中了 paragraph ，则在 cell 内换行
            newEditor.insertText('\n')
            return
        }
        // 未选中 paragraph ，默认的换行
        insertBreak()
        return
    }
    // 返回 newEditor ，重要！
    return newEditor
}
Boot.registerPlugin(withParagraph)

class SpeakerButtonMenu implements IButtonMenu {
    title: string
    tag: string
    constructor(title: string) {
        this.title = title // 自定义菜单标题
        // this.iconSvg = '<svg>...</svg>' // 可选
        this.tag = 'button'
    }

    // 获取菜单执行时的 value ，用不到则返回空 字符串或 false
    getValue(editor: IDomEditor): string | boolean {
        // 去除[[]]中间的字符串
        // 如 [[晓晓]]注：食量参考来源是家里参加过知青下乡的长辈。
        // 去除[[晓晓]]
        return editor.getSelectionText().replace(/\[\[.*?\]\]/g, '');
    }

    // 菜单是否需要激活（如选中加粗文本，“加粗”菜单会激活），用不到则返回 false
    isActive(editor: IDomEditor): boolean {
        return false
    }

    // 菜单是否需要禁用（如选中 H1 ，“引用”菜单被禁用），用不到则返回 false
    isDisabled(editor: IDomEditor): boolean {
        return false
    }

    // 点击菜单时触发的函数
    exec(editor: IDomEditor, value: string | boolean) {
        if (this.isDisabled(editor)) return
        // 如果value是字符串 执行editor.insertText(value)
        // value 即 this.value(editor) 的返回值
        if (typeof value === 'string') {
            editor.insertText(`\n[[${this.title}]]${value}`);
        }
    }
}

// 晓晓-女-暖 晓伊-女- 云健-男-体育影视 云希-男-暖 云夏-男-少年 云扬-男-专业流利
export const speakers = ['晓晓', '晓伊', '云健', '云希', '云夏', '云扬']

speakers.forEach(name => {
    Boot.registerMenu({
        key: name,
        factory() {
            return new SpeakerButtonMenu(name)
        },
    })
})