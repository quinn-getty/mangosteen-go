package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
}

func (ctrl *ItemController) RegisterRouter(rg *gin.RouterGroup) {
	item := rg.Group("/item")
	item.POST("", ctrl.Create)
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
//	@Tags			item
//	@Accept			json
//	@Produce		json
//	@Param			body		body		CreateItemReq			true	"body参数"
//	@Success		200			{object}	CreateItemRes
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

func (ctrl *ItemController) Delete(c *gin.Context)  {}
func (ctrl *ItemController) Update(c *gin.Context)  {}
func (ctrl *ItemController) Get(c *gin.Context)     {}
func (ctrl *ItemController) getList(c *gin.Context) {}
