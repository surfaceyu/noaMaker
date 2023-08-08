package dao

import "github.com/surfaceyu/noaMaker/backend/models"

const bookSitesKey = "bookSites"

var bookSites []models.BookSite

func initBookSites() {
	db.View(bookSitesKey, &bookSites)
}

func saveBookSites() error {
	return db.Update(bookSitesKey, bookSites)
}

func GetBookSites() []models.BookSite {
	return bookSites
}

func AddBookSites(bs models.BookSite) error {
	// 判断bookSites []models.BookSite 中是否有bs
	for _, site := range bookSites {
		if site.Name == bs.Name {
			return nil // 如果已经存在，直接返回nil，表示操作成功
		}
	}
	// 如果bookSites中不存在bs，则将其插入数组中
	bookSites = append(bookSites, bs)
	// 保存进入数据库
	return saveBookSites()
}

func EditBookSites(bs models.BookSite) error {
	// 判断bookSites []models.BookSite 中是否有bs
	// 有site.Name == bs.Name 则将site赋值为bs
	for i, site := range bookSites {
		if site.Name == bs.Name {
			// 将site赋值为bs
			bookSites[i] = bs
			break
		}
	}
	// 保存进入数据库
	return saveBookSites()
}

func DelBookSites(bs models.BookSite) error {
	// 判断bookSites []models.BookSite 中是否有bs
	// 如果bookSites中存在bs，则将其删除
	index := -1
	for i, site := range bookSites {
		if site.Name == bs.Name {
			index = i
			break
		}
	}
	if index != -1 {
		// 从bookSites中删除bs
		bookSites = append(bookSites[:index], bookSites[index+1:]...)
	}
	// 保存进入数据库
	return saveBookSites()
}
