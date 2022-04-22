package httpReq

type PageReq struct {
	limit int `json:"limit"`
	skip  int `json:"skip"`
}

func NewPageReq(limit, skip int) PageReq {
	return PageReq{
		limit: limit,
		skip:  skip,
	}
}
