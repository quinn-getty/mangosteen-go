package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TagController struct {
}

type CreateTagReq struct {
	Sign string `json:"sign" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateTagRes struct {
	Resource queries.Tag `json:"resource"`
}

// CreateItem godoc
//
//	@Summary		创建tag
//	@Description	创建tag
//	@Security		Beare
//	@Tags			tag
//	@Accept			json
//	@Produce		json
//	@Param			body	body		CreateTagReq	true	"body参数"
//	@Success		200		{object}	CreateTagRes
//	@Router			/tag [post]
func (ctrl *TagController) Create(c *gin.Context) {
	req := CreateTagReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Print("入参错误", err)
		c.String(http.StatusUnprocessableEntity, "参数错误")
		return
	}

	user, _ := middleware.GetMe(c)
	q := database.NewQuery()
	item, err := q.CreateTag(c, queries.CreateTagParams{
		UserID: user.ID,
		Name:   req.Name,
		Sign:   req.Sign,
	})

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateTagRes{
		Resource: item,
	})
}

func (ctrl *TagController) Delete(c *gin.Context)  {}
func (ctrl *TagController) Update(c *gin.Context)  {}
func (ctrl *TagController) Get(c *gin.Context)     {}
func (ctrl *TagController) getList(c *gin.Context) {}
