package api

import "mangosteen/config/queries"

type CreateTagReq struct {
	Sign string `json:"sign" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateTagRes struct {
	Resource queries.Tag `json:"resource"`
}

type DeleteTagRes struct {
	Resource queries.Tag `json:"resource"`
}

type TagUpdateReq struct {
	Id   int32  `json:"id" binding:"required"`
	Sign string `json:"sign"`
	Name string `json:"name"`
}

type TagUpdateRes struct {
	Resource queries.Tag `json:"resource"`
}

type FindTagRes struct {
	Resource queries.Tag `json:"resource"`
}
type TagListReq struct{}

type TagListRes struct {
	Resource []queries.Tag `json:"resource"`
}
