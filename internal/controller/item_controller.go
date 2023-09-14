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
}

type ItemGetListReq struct {
	Current         int32     `json:"current" binding:"required"`
	Size            int32     `json:"size" binding:"required"`
	HappenenAtBegin time.Time `json:"happenenAtBegin" `
	HappenenAtEnd   time.Time `json:"happenenAtEnd" `
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

	var offset int32

	currentStr, _ := c.GetQuery("current")
	if current, err := strconv.Atoi(currentStr); err == nil {
		params.Current = int32(current)
		if current == 1 {
			offset = 0
		} else {
			offset = int32(current)
		}
	}

	sizeStr, _ := c.GetQuery("size")
	if size, err := strconv.Atoi(sizeStr); err == nil {
		params.Size = int32(size)
	}

	// happenenAtBegin, err := time.Parse(time.RFC3339, dateString)
	// if err != nil {
	// 	fmt.Println("解析日期时间出错:", err)
	// 	return
	// }
	// happenenAtBeginStr, _ := c.GetQuery("happenenAtBegin")
	// happenenAtEndStr, _ := c.GetQuery("happenenAtEnd")

	if !ok {
		c.Status(http.StatusUnauthorized)
		return
	}

	q := database.NewQuery()
	count, err := q.CountItem(c, user.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}

	list, err := q.ListItem(c, queries.ListItemParams{
		UserID: user.ID,
		Offset: offset * params.Size,
		Limit:  params.Size,
		// HappenedAt: ,
		// HappenedAt_2: ,
	})
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}

	res := ItemGetListRes{
		Resourses: list,
		Pager: Pager{
			Current: params.Current,
			Size:    params.Size,
			Total:   count,
		},
	}

	c.JSON(http.StatusOK, res)
}
