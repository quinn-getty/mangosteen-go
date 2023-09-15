package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
}

func (ctrl *ItemController) RegisterRouter(rg *gin.RouterGroup) {
	item := rg.Group("/item")
	item.POST("", ctrl.Create)
	item.GET("", ctrl.getList)
}

type CreateItemReq struct {
	Amount     int32        `json:"amount" binding:"required"`
	Kind       queries.Kind `json:"kind" binding:"required"`
	HappenedAt time.Time    `json:"happenedAt" binding:"required"`
	TagIds     []int32      `json:"tagIds" binding:"required"`
}

type CreateItemRes struct {
	Resource queries.Item `json:"resource"`
}

// CreateItem godoc
//
//	@Summary		创建item
//	@Description	创建item
//	@Security		Bearer
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateItemReq	true	"body参数"
//	@Success		200		{object}	CreateItemRes
//	@Router			/item [post]
func (ctrl *ItemController) Create(c *gin.Context) {
	req := CreateItemReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Print("入参错误", err)
		c.String(http.StatusUnprocessableEntity, "参数错误")
		return
	}

	user, _ := middleware.GetMe(c)
	q := database.NewQuery()
	item, err := q.CreateItem(c, queries.CreateItemParams{
		UserID:     user.ID,
		Amount:     req.Amount,
		Kind:       req.Kind,
		HappenedAt: req.HappenedAt,
		TagIds:     req.TagIds,
	})

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateItemRes{
		Resource: item,
	})

}

func (ctrl *ItemController) Delete(c *gin.Context) {}
func (ctrl *ItemController) Update(c *gin.Context) {}
func (ctrl *ItemController) Get(c *gin.Context)    {}

type Pager struct {
	Total   int64 `json:"total"`
	Current int32 `json:"current" binding:"required"`
	Size    int32 `json:"size" binding:"required"`
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

// ItemList godoc
//
//	@Summary		item List
//	@Description	item List
//	@Security		Bearer
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			query	query		ItemGetListReq	true	"query参数"
//	@Success		200		{object}	ItemGetListRes
//	@Router			/item [get]
func (ctrl *ItemController) getList(c *gin.Context) {
	params := ItemGetListReq{}
	user, ok := middleware.GetMe(c)

	currentStr, _ := c.GetQuery("current")
	if current, err := strconv.Atoi(currentStr); err == nil {
		if current == 0 {
			params.Current = 1
		} else {
			params.Current = int32(current)
		}
	}

	sizeStr, _ := c.GetQuery("size")
	if size, err := strconv.Atoi(sizeStr); err == nil {
		params.Size = int32(size)
	}

	happenedAtBeginStr, _ := c.GetQuery("happenedAtBegin")
	if happenedAtBeginStr == "" {
		params.HappenedAtBegin = time.Now().AddDate(-100, 0, 0)
	} else if happenedAtBegin, err := time.Parse(time.RFC3339, happenedAtBeginStr); err != nil {
		params.HappenedAtBegin = happenedAtBegin
	} else {
		params.HappenedAtBegin = time.Now().AddDate(-100, 0, 0)
	}

	happenedAtEndStr, _ := c.GetQuery("happenedAtEnd")
	if happenedAtEndStr == "" {
		params.HappenedAtEnd = time.Now().AddDate(0, 0, 1)
	} else if happenedAtEnd, err := time.Parse(time.RFC3339, happenedAtEndStr); err != nil {
		params.HappenedAtEnd = happenedAtEnd
	} else {
		params.HappenedAtEnd = time.Now().AddDate(0, 0, 1)
	}

	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	q := database.NewQuery()

	res := ItemGetListRes{
		Resourses: []queries.Item{},
		Pager: Pager{
			Current: params.Current,
			Size:    params.Size,
			Total:   0,
		},
		Income:   0,
		Expenses: 0,
	}

	count, err := q.CountItem(c, user.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}
	res.Pager.Total = count

	income, err := q.ItemsBalance(database.DBCtx, queries.ItemsBalanceParams{
		UserID:          user.ID,
		HappenedAtBegin: params.HappenedAtBegin,
		HappenedAtEnd:   params.HappenedAtEnd,
		Kind:            queries.KindInCome,
	})

	for _, i := range income {
		res.Income += i
	}

	if err != nil {
		log.Print("查询income ", err)
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}

	expenses, err := q.ItemsBalance(database.DBCtx, queries.ItemsBalanceParams{
		UserID:          user.ID,
		HappenedAtBegin: params.HappenedAtBegin,
		HappenedAtEnd:   params.HappenedAtEnd,
		Kind:            queries.KindExpenses,
	})
	if err != nil {
		// expenses = 0
		log.Print("查询expenses", err)
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}

	for _, i := range expenses {
		res.Expenses += i
	}

	list, err := q.ListItem(c, queries.ListItemParams{
		UserID:          user.ID,
		Offset:          (params.Current - 1) * params.Size,
		Limit:           params.Size,
		HappenedAtBegin: params.HappenedAtBegin,
		HappenedAtEnd:   params.HappenedAtEnd,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}
	res.Resourses = list

	c.JSON(http.StatusOK, res)
}
