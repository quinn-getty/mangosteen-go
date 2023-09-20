package api

import (
	"mangosteen/config/queries"
	"time"
)

type CreateItemReq struct {
	Amount     int32        `json:"amount" binding:"required"`
	Kind       queries.Kind `json:"kind" binding:"required"`
	HappenedAt time.Time    `json:"happenedAt" binding:"required"`
	TagIds     []int32      `json:"tagIds" binding:"required"`
}

type CreateItemRes struct {
	Resource queries.Item `json:"resource"`
}

type ItemGetListRes struct {
	Resourses []queries.Item `json:"resourses"`
	Pager     Pager          `json:"pager"`
	Income    int32          `json:"inCome"`
	Expenses  int32          `json:"expenses"`
}

type ItemGetListReq struct {
	Current         int32     `json:"current" binding:"required"`
	Size            int32     `json:"size" binding:"required"`
	HappenedAtBegin time.Time `json:"happenedAtBegin" `
	HappenedAtEnd   time.Time `json:"happenedAtEnd" `
}

type ItemGetSummaryReq struct {
	HappenedAtBegin time.Time    `json:"happenedAtBegin" binding:"required"`
	HappenedAtEnd   time.Time    `json:"happenedAtEnd" binding:"required"`
	Kind            queries.Kind `json:"kind" binding:"required"`
	// GroupBy         string       `json:"groupBy" binding:"required"`
}

type ItemGetSummaryRes struct {
	Groups []ItemGetSummaryResGroups `json:"groups"`
	Total  int32                     `json:"total"`
}

type ItemGetSummaryResGroups struct {
	HappenedAt string `json:"happenedAt"`
	Amount     int32  `json:"amount"`
}
