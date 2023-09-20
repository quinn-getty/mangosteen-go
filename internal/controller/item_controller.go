package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/api"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"
	"net/http"
	"sort"
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
	item.GET("/summary", ctrl.GetSummary)
}

// CreateItem godoc
//
//	@Summary		创建item
//	@Description	创建item
//	@Security		Bearer
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			body	body		api.CreateItemReq	true	"body参数"
//	@Success		200		{object}	api.CreateItemRes
//	@Router			/item [post]
func (ctrl *ItemController) Create(c *gin.Context) {
	req := api.CreateItemReq{}

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

	c.JSON(http.StatusOK, api.CreateItemRes{
		Resource: item,
	})

}

func (ctrl *ItemController) Delete(c *gin.Context) {}
func (ctrl *ItemController) Update(c *gin.Context) {}
func (ctrl *ItemController) Get(c *gin.Context)    {}

// ItemList godoc
//
//	@Summary		item List
//	@Description	item List
//	@Security		Bearer
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			query	query		api.ItemGetListReq	true	"query参数"
//	@Success		200		{object}	api.ItemGetListRes
//	@Router			/item [get]
func (ctrl *ItemController) getList(c *gin.Context) {
	params := api.ItemGetListReq{
		HappenedAtBegin: time.Now().AddDate(-100, 0, 0),
		HappenedAtEnd:   time.Now().AddDate(0, 0, 1),
	}

	user, ok := middleware.GetMe(c)
	if !ok {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}

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
	if happenedAtBegin, err := time.Parse(time.DateTime, happenedAtBeginStr); err != nil {
		log.Print("解析出错happenedAtBegin ", happenedAtBeginStr, " ", err)
	} else {
		log.Println(happenedAtBegin)
		params.HappenedAtBegin = happenedAtBegin.Add(time.Minute * 59).Add(time.Second * 59)
	}

	happenedAtEndStr, _ := c.GetQuery("happenedAtEnd")
	if happenedAtEnd, err := time.Parse(time.DateTime, happenedAtEndStr); err != nil {
		log.Print("解析出错happenedAtEnd ", happenedAtEnd, " ", err)
	} else {
		log.Println(happenedAtEnd)
		params.HappenedAtEnd = happenedAtEnd.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
	}

	q := database.NewQuery()

	res := api.ItemGetListRes{
		Resourses: []queries.Item{},
		Pager: api.Pager{
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
	log.Print("params: ", params)
	res.Resourses = list

	c.JSON(http.StatusOK, res)
}

// ItemList godoc
//
//	@Summary		item List
//	@Description	item List
//	@Security		Bearer
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			query	query		api.ItemGetSummaryReq	true	"query参数"
//	@Success		200		{object}	api.ItemGetSummaryRes
//	@Router			/summary [get]
func (ctrl *ItemController) GetSummary(c *gin.Context) {
	req := api.ItemGetSummaryReq{}
	res := api.ItemGetSummaryRes{}
	res.Groups = []api.ItemGetSummaryResGroups{}
	res.Total = 0
	q := database.NewQuery()
	user, _ := middleware.GetMe(c)

	log.Print("url", c.Request.URL)

	if kind, _ := c.GetQuery("kind"); kind == "" {
		log.Print(kind)
		c.Status(http.StatusBadRequest)
		return
	} else {
		req.Kind = queries.Kind(kind)
	}

	happenedAtBeginStr, _ := c.GetQuery("happenedAtBegin")
	if happenedAtBegin, err := time.Parse(time.DateOnly, happenedAtBeginStr); err != nil {
		log.Print("解析出错happenedAtBegin ", happenedAtBeginStr, " ", err)
		c.Status(http.StatusBadRequest)
		return
	} else {
		req.HappenedAtBegin = happenedAtBegin
	}

	happenedAtEndStr, _ := c.GetQuery("happenedAtEnd")
	if happenedAtEnd, err := time.Parse(time.DateOnly, happenedAtEndStr); err != nil {
		log.Print("解析出错happenedAtEnd ", happenedAtEnd, " ", err)
		c.Status(http.StatusBadRequest)
		return
	} else {
		req.HappenedAtEnd = happenedAtEnd
	}

	log.Println(req.Kind)

	items, err := q.ListItemsByHappenedAtAndKind(database.DBCtx, queries.ListItemsByHappenedAtAndKindParams{
		Kind:            req.Kind,
		HappenedAtBegin: req.HappenedAtBegin,
		HappenedAtEnd:   req.HappenedAtEnd,
		UserID:          user.ID,
	})
	log.Println(items)
	if err != nil {
		c.Status(500)
	}

	log.Print("items:", items)

	for _, item := range items {
		k := item.HappenedAt.Format(time.DateOnly)
		res.Total += item.Amount

		found := false
		for index, group := range res.Groups {
			if group.HappenedAt == k {
				found = true
				res.Groups[index].Amount += item.Amount
			}
		}
		if !found {
			res.Groups = append(res.Groups, api.ItemGetSummaryResGroups{
				HappenedAt: k,
				Amount:     item.Amount,
			})
		}
	}
	sort.Slice(res.Groups, func(i, j int) bool {
		return res.Groups[i].HappenedAt < res.Groups[j].HappenedAt
	})

	log.Println(res)

	c.JSON(http.StatusOK, res)
}
