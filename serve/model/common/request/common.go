package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page         int    `json:"page" form:"page"`            // 页码
	PageSize     int    `json:"pageSize" form:"pageSize"`    // 每页大小
	Keyword      string `json:"keyword" form:"keyword"`      //关键字
	Authority_Id int64  `json:"authority_id" form:"keyword"` //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

type GetByName struct {
	Name string `json:"name" form:"id"` // 主键ID
}

type GetByNcrID struct {
	NcrID      int64  `json:"ncr_id" form:"ncr_id"`
	ReportName string `json:"reportname" form:"id"` // 主键ID
}

func (r *GetById) Uint() uint {
	return uint(r.ID)
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId uint `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}
