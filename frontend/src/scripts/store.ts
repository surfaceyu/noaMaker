import { reactive, ref } from "vue";
import { models } from '@wailsjs/go/models';
import { GetSites } from '@backend/BookSitesHandler';

// 编辑界面
export namespace editRule {
    export const bookData = ref<models.Book>()
    export const chapterData = ref<models.Chapter>()
    export const contentData = ref<models.Content>()
}

export const bookSiteRules = ref<models.BookSite[]>([]);
export const editRuleRow = ref(new models.BookSite())

export async function useRule() {
    bookSiteRules.value = await GetSites()
    return bookSiteRules
}