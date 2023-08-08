package driver

import (
	"testing"

	"github.com/surfaceyu/noaMaker/backend/models"
)

var bookSite = models.BookSite{
	Name: "笔趣阁",
	Uri:  "111111",
	SearchUrl: models.BookSiteUri{
		Uri:      "https://www.xbiquwx.la/modules/article/search.php?searchkey={{key}}",
		BookList: ".grid tbody > tr",
		Author:   ".odd:nth-child(3)",
		BookName: "td.odd > a",
		BookUri:  "td.odd > a[href]",
		BookId:   "\\/([^\\/]+)\\/$",
	},
	ChapterUrl: models.BookSiteChapterUri{
		Uri:         "https://www.xbiquwx.la/{{bookId}}/",
		ChapterList: "#list dd",
		ChapterName: "a",
		ChapterUri:  "a[href]",
		ChapterId:   "^(.+)\\.html$",
	},
	ContentUrl: models.BookSiteContentUri{
		Uri:     "https://www.xbiquwx.la/{{bookId}}/{{chapterId}}.html",
		Content: "#content",
	},
}

func TestSearchBook(t *testing.T) {
	type args struct {
		name string
		site models.BookSite
	}
	tests := []struct {
		name string
		args args
		want []models.Book
	}{
		{
			name: "Positive test",
			args: args{
				name: "剑仙",
				site: bookSite,
			},
			want: []models.Book{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchBook(tt.args.name, tt.args.site)
			if len(got) <= 0 {
				t.Errorf("Search() = %v, want %v", got, tt.want)
			} else {
				t.Logf("Search() = %v", got)
			}
		})
	}
}

func TestSearchChapter(t *testing.T) {
	type args struct {
		book models.Book
		site models.BookSite
	}
	tests := []struct {
		name string
		args args
		want []models.Chapter
	}{
		{
			name: "Positive test",
			args: args{
				book: models.Book{
					Author:   "懒源",
					BookName: "剑仙",
					BookId:   "125_125767",
				},
				site: bookSite,
			},
			want: []models.Chapter{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchChapter(tt.args.book, tt.args.site)
			if len(got) <= 0 {
				t.Errorf("SearchChapter() = %v, want %v", got, tt.want)
			} else {
				t.Logf("SearchChapter() = %v", got)
			}
		})
	}
}

func TestSearchContent(t *testing.T) {
	type args struct {
		chapter models.Chapter
		site    models.BookSite
	}
	tests := []struct {
		name string
		args args
		want models.Content
	}{
		{
			name: "Positive test",
			args: args{
				chapter: models.Chapter{
					BookId:      "125_125767",
					ChapterId:   "32646339",
					ChapterName: "第一百五十八章：太始之息",
				},
				site: bookSite,
			},
			want: models.Content{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SearchContent(tt.args.chapter, tt.args.site)
			t.Logf("SearchContent() = %v", got)
		})
	}
}
