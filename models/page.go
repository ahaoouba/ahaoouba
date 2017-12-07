package models

import (
	"math"
	"net/url"
	"strconv"
)

const (
	PAGE_KEY    = "p"
	PAGE_MAX_NO = 5                 //页码显示最大数量
	PAGE_OFFSET = PAGE_MAX_NO/3 + 1 //页码显示左右偏移量
)

//分页结构
type Page struct {
	PageNo     int    //当前页号
	PageSize   int    //每页条数
	TotalPage  int    //总页数
	TotalCount int    //总数据量
	url        string //页面链接
	pageRanges []int
}

func NewPage(total int, pageNo int, pageSize int, url string) *Page {
	tp := total / pageSize
	if total%pageSize > 0 {
		tp = total/pageSize + 1
	}
	return &Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: total, url: url}
}

//是否有上一页
func (p *Page) HasPreview() bool {
	if p.TotalPage > 0 {
		return p.PageNo > 1
	}
	return false

}

//是否有下一页
func (p *Page) HasNext() bool {
	if p.TotalPage > 0 {
		return p.TotalPage > p.PageNo
	}
	return false
}

// 获取起始页数
func (p *Page) StartNum() int {
	if p.PageNo > 0 && p.TotalPage > 0 {
		return (p.PageNo-1)*p.PageSize + 1
	}
	if p.TotalPage == 0 {
		return 0
	}
	return 1
}

// 获取结尾页数
func (p *Page) EndNum() int {
	var en = p.PageSize
	if p.PageNo > 0 {
		en = p.PageNo * p.PageSize
	}
	if en > p.TotalCount {
		return p.TotalCount
	}

	return en
}

// 获取分页显示数组
func (p *Page) Pages() []int {
	maxNo := math.Min(float64(PAGE_MAX_NO), float64(p.TotalPage))
	p.pageRanges = make([]int, int(maxNo))
	if PAGE_MAX_NO >= p.TotalPage {
		for i, _ := range p.pageRanges {
			p.pageRanges[i] = i + 1
		}
	} else if p.TotalPage > PAGE_MAX_NO {
		// 1 2 3 4 5 6 7 8 9
		//当前页，总页数，显示页码
		switch {
		case p.PageNo <= PAGE_OFFSET:
			for i, _ := range p.pageRanges {
				p.pageRanges[i] = i + 1
			}
		case p.PageNo+PAGE_OFFSET < p.TotalPage:
			for i, _ := range p.pageRanges {
				p.pageRanges[i] = p.PageNo - PAGE_OFFSET + i
			}
		default:
			for i, _ := range p.pageRanges {
				p.pageRanges[i] = p.TotalPage - PAGE_MAX_NO + 1 + i
			}
		}

	}

	return p.pageRanges
}

//生成跳转页面链接
func (p *Page) Link(page int) string {
	link, _ := url.ParseRequestURI(p.url)
	values := link.Query()
	values.Set(PAGE_KEY, strconv.Itoa(page))
	values.Del("token")
	link.RawQuery = values.Encode()
	return link.String()
}

// 上一页跳转页面
func (p *Page) PreviewLink() string {
	return p.Link(p.PageNo - 1)
}

// 上一页跳转页面
func (p *Page) NextLink() string {
	return p.Link(p.PageNo + 1)
}

//// 首页跳转页面
//func (p *Page) FirstLink() string {
//	return p.Link(1)
//}
