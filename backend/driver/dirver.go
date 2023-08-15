package driver

import (
	"regexp"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/surfaceyu/noaMaker/backend/models"
)

func extractHref(uri string, reg string) string {
	re := regexp.MustCompile(reg)
	match := re.FindStringSubmatch(uri)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func SearchBook(name string, rule models.BookSiteUri) []models.Book {
	var results []models.Book
	// uri 将 rule.Uri 中的"{{key}}" 替换为name
	replacedUri := strings.ReplaceAll(rule.Uri, "{{key}}", name)
	// 使用colly爬取uri
	c := colly.NewCollector()
	// 设置回调函数，在爬取到指定的HTML元素时进行处理
	c.OnHTML(rule.BookList, func(e *colly.HTMLElement) {
		// 按照rule中的选择器 获取BookList Author BookName BookUri BookUri BookId
		author := e.ChildText(rule.Author)
		bookName := e.ChildText(rule.BookName)
		bookUri := e.ChildAttr(rule.BookUri, "href")
		bookId := extractHref(bookUri, rule.BookId)
		book := models.Book{
			Author:   author,
			BookName: bookName,
			BookId:   bookId,
		}
		if book.BookName != "" {
			results = append(results, book)
		}
	})

	// 发起请求，开始爬取网页数据
	c.Visit(replacedUri)
	return results
}

func SearchChapter(book models.Book, site models.BookSite) []models.Chapter {
	var results []models.Chapter
	// uri 将 site.ChapterUrl.Uri 中的"{{bookId}}" 替换为 book.BookId
	replacedUri := strings.ReplaceAll(site.ChapterUrl.Uri, "{{bookId}}", book.BookId)
	// 使用colly爬取uri
	c := colly.NewCollector()
	// 设置回调函数，在爬取到指定的HTML元素时进行处理
	c.OnHTML(site.ChapterUrl.ChapterList, func(e *colly.HTMLElement) {
		// 按照site.ChapterUrl中的选择器 获取 ChapterName ChapterId
		chapterName := e.ChildText(site.ChapterUrl.ChapterName)
		bookUri := e.ChildAttr(site.ChapterUrl.ChapterUri, "href")
		chapterId := extractHref(bookUri, site.ChapterUrl.ChapterId)
		chapter := models.Chapter{
			BookId:      book.BookId,
			ChapterId:   chapterId,
			ChapterName: chapterName,
		}
		if chapter.ChapterId != "" {
			results = append(results, chapter)
		}
	})

	// 发起请求，开始爬取网页数据
	c.Visit(replacedUri)
	return results
}

func SearchContent(chapter models.Chapter, site models.BookSite) models.Content {
	var results models.Content
	// uri 将 site.ContentUrl.Uri 中的"{{bookId}}" 替换为 chapter.BookId
	replacedUri := strings.ReplaceAll(site.ContentUrl.Uri, "{{bookId}}", chapter.BookId)
	replacedUri = strings.ReplaceAll(replacedUri, "{{chapterId}}", chapter.ChapterId)
	// 使用colly爬取uri
	c := colly.NewCollector()
	// 设置回调函数，在爬取到指定的HTML元素时进行处理
	c.OnHTML(site.ContentUrl.Content, func(e *colly.HTMLElement) {
		book := models.Content{
			BookId:    chapter.BookId,
			ChapterId: chapter.ChapterId,
			Content:   e.Text,
		}
		if book.Content != "" {
			results = book
		}
	})

	// 发起请求，开始爬取网页数据
	c.Visit(replacedUri)
	return results
}
