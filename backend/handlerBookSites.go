package backend

import (
	"github.com/surfaceyu/myproject/backend/dao"
	"github.com/surfaceyu/myproject/backend/models"
)

func init() {
	bindHandlers = append(bindHandlers, &BookSitesHandler{})
}

type Book struct {
	Name   string            `json:"name"`
	Author string            `json:"author"`
	Desc   string            `json:"desc"`
	Sites  []models.BookSite `json:"sites"`
}

type BookSitesHandler struct {
}

// GetSites 返回所有书籍网站的列表
func (g *BookSitesHandler) GetSites() []models.BookSite {
	return dao.GetBookSites()
}

// AddSites 添加一个书籍网站到列表中
// 参数：bs BookSite - 要添加的书籍网站
func (g *BookSitesHandler) AddSites(bs models.BookSite) {
	if err := dao.AddBookSites(bs); err == nil {
		// 执行成功 更新前端
	}
}

func (g *BookSitesHandler) EditSites(bs models.BookSite) error {
	return dao.EditBookSites(bs)
}

// DelSites 从列表中删除一个书籍网站
// 参数：bs BookSite - 要删除的书籍网站
func (g *BookSitesHandler) DelSites(bs models.BookSite) {
	if err := dao.DelBookSites(bs); err == nil {
		// 执行成功 更新前端
	}
}

func (g *BookSitesHandler) GetBooks() []Book {
	return []Book{
		{Name: "大师兄不在家", Author: "咸鱼", Desc: "一本测试的小说", Sites: []models.BookSite{
			{Uri: "www.hhhhhzzz.com"},
		}},
		{Name: "小师妹不在家", Author: "宝剑", Desc: "宝剑亚甲基地方还等哈来算的话拉丁化拉客的火辣的你拉省的回答克里斯蒂哈拉省的啦DHL", Sites: []models.BookSite{
			{Uri: "www.aaaaaaaaaaaaaa.com"},
			{Uri: "www.ppppppppppppp.com"},
		}},
		{Name: "师父去哪了", Author: "咸鱼", Desc: "快乐与哦辽宁路六年看来你健康噜噜噜农科路快乐和健康陆军预备役如图古", Sites: []models.BookSite{
			{Uri: "www.bbbbbbbbbbbb.com"},
			{Uri: "www.cccccccccccc.com"},
			{Uri: "www.dddddddddddd.com"},
			{Uri: "www.eeeeeeeeeeeee.com"},
		}},
	}
}
