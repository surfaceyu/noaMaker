package backend

import (
	"github.com/surfaceyu/noaMaker/backend/dao"
	"github.com/surfaceyu/noaMaker/backend/driver"
	"github.com/surfaceyu/noaMaker/backend/models"
)

func init() {
	bindHandlers = append(bindHandlers, &BookHandler{})
}

type BookHandler struct {
}

// 搜索书本
func (b *BookHandler) SearchBook(name string, siteIndex int) []models.Book {
	sites := dao.GetBookSites()
	// 判断site是否存在 不存在则返回
	if siteIndex < 0 || siteIndex >= len(sites) {
		return nil
	}
	site := sites[siteIndex]
	// 从 site中搜索书籍
	// 返回书籍列表
	return driver.SearchBook(name, site.SearchUrl)
}

func (b *BookHandler) SearchBookByRule(name string, siteRule models.BookSiteUri) []models.Book {
	// 从 site中搜索书籍
	// 返回书籍列表
	return driver.SearchBook(name, siteRule)
}

// 搜索书本章节列表
func (b *BookHandler) SearchChapter(book models.Book, site models.BookSite) []models.Chapter {
	return driver.SearchChapter(book, site)
}

// 搜索书本章节内容
func (b *BookHandler) SearchContent(chapter models.Chapter, site models.BookSite) models.Content {
	return driver.SearchContent(chapter, site)
}
