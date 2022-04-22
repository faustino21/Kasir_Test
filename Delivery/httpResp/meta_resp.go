package httpResp

type MetaResp struct {
	Total int `json:"total"`
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

func NewMetaResp(total, limit, skip int) MetaResp {
	return MetaResp{
		total, limit, skip,
	}
}
