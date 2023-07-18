package models

type BookSite struct {
	Name       string             `json:"name"`       // 书源-名称
	Uri        string             `json:"uri"`        // 书源-搜索链接
	SearchUrl  BookSiteUri        `json:"searchUrl"`  // 书源-搜索解析
	ChapterUrl BookSiteChapterUri `json:"chapterUrl"` // 书源-章节解析
	ContentUrl BookSiteContentUri `json:"contentUrl"` // 书源-内容解析
}

type BookSiteUri struct {
	Uri      string `json:"uri"`      // 书本搜索链接
	BookList string `json:"bookList"` // 书本解析-列表
	Author   string `json:"author"`   // 书本解析-作者
	BookName string `json:"bookName"` // 书本解析-名称
	BookUri  string `json:"bookUri"`  // 书本解析-所有章节链接
	BookId   string `json:"bookId"`   // 书本解析-书本Id
}

type BookSiteChapterUri struct {
	Uri         string `json:"uri"`
	ChapterList string `json:"chapterList"` // 章节列表
	ChapterName string `json:"chapterName"` // 章节名称
	ChapterUri  string `json:"chapterUrl"`  // 章节链接
	ChapterId   string `json:"chapterId"`   // 章节Id
}

type BookSiteContentUri struct {
	Uri     string `json:"uri"`
	Content string `json:"content"` // 章节内容
}
