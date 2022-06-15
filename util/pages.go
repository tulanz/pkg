package util

// 分页信息
type Page struct {
	PageNo     int         `json:"page_no"`
	PageSize   int         `json:"page_size"`
	TotalPage  int         `json:"total_page"`
	TotalCount int         `json:"total_count"`
	FirstPage  bool        `json:"first_page"`
	LastPage   bool        `json:"last_page"`
	List       interface{} `json:"list"`
}

// PageUtil生成分页结构工具函数
func PageUtil(count int, pageNo int, pageSize int, list ...interface{}) Page {
	if pageSize == 0 {
		pageSize = 5
	}
	tp := count / pageSize
	if count%pageSize > 0 || count == 0 {
		tp = count/pageSize + 1
	}

	page := Page{
		PageNo:     pageNo,
		PageSize:   pageSize,
		TotalPage:  tp,
		TotalCount: count,
		FirstPage:  pageNo == 1,
		LastPage:   pageNo == tp,
	}
	if list != nil && len(list) > 0 {
		page.List = list[0]
	}
	return page
}
