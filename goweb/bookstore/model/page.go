package model

//Page 结构体
type Page struct {
	Books       []*Book //每页查询出来的的图书存放的切片
	PageNo      int64   //当前页
	PageSize    int64   //每页显示的条数
	TotalPageNo int64   //总页数，通过计算得到
	TotalRecord int64   //总记录数
	MinPrice    string  //最小金额
	MaxPrice    string  //最大金额
	IsLogin     bool    //是否登录
	UserName    string  //登录用户名

}

//IsHasPrev 判断是否有上一页
func (page *Page) IsHasPrev() bool {
	return page.PageNo > 1
}

//IsHasNext 判断是否有上一页
func (page *Page) IsHasNext() bool {
	return page.PageNo < page.TotalPageNo
}

//GetPrevPageNo 获取上一页
func (page *Page) GetPrevPageNo() int64 {
	if page.IsHasPrev() {
		return page.PageNo - 1
	} else {
		return page.PageNo
	}
}

//GetNextPageNo 获取下一页
func (page *Page) GetNextPageNo() int64 {
	if page.IsHasNext() {
		return page.PageNo + 1
	} else {
		return page.PageNo
	}
}
