package controller

import (
	"log"
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
	Amount     int       `json:"amount" binding:"required"`
	Kind       string    `json:"kind" binding:"required"`
	HappenedAt time.Time `json:"happenedAt" binding:"required"`
	TagIds     []int32   `json:"tagIds" binding:"required"`
}

// CreateItem godoc
// @Summary      创建item
// @Description  创建item
// @Tags         item
// @Accept       json
// @Produce      json
// @Success      200
// @Failure      500
// @Router       /item [post]
func (ctrl *ItemController) Create(c *gin.Context) {
	req := CreateItemReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Println("入参错误")
		c.String(http.StatusUnprocessableEntity, "参数错误")
		return
	}

}

func (ctrl *ItemController) Delete(c *gin.Context)  {}
func (ctrl *ItemController) Update(c *gin.Context)  {}
func (ctrl *ItemController) Get(c *gin.Context)     {}
func (ctrl *ItemController) getList(c *gin.Context) {}
