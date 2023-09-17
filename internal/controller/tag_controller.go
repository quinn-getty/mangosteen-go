package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/database"
	"mangosteen/internal/middleware"
	"net/http"
	"strconv"
	"time"

	_ "database/sql"

	"github.com/gin-gonic/gin"
)

type TagController struct {
}

func (ctrl *TagController) RegisterRouter(rg *gin.RouterGroup) {
	item := rg.Group("/tag")
	item.POST("", ctrl.Create)
	item.GET("", ctrl.getList)
	item.PATCH("", ctrl.Update)
}

type CreateTagReq struct {
	Sign string `json:"sign" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type CreateTagRes struct {
	Resource queries.Tag `json:"resource"`
}

// CreateTag godoc
//
// @Summary		tag
// @Description	获取tag
// @Tags			tag
// @Accept			json
// @Produce		json
//
// @Security		Bearer
//
// @Param			body	body		CreateTagReq	true	"body参数"
// @Success		200		{object}	CreateTagRes
// @Failure		500
// @Router			/tag [post]
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

type DeleteTagRes struct {
	Resource int32 `json: "resource"`
}

// TagUpdate godoc
//
//	@Summary		tag
//	@Description	获取tag
//	@Tags			tag
//	@Accept			json
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	DeleteTagRes
//	@Failure		500
//	@Router			/tag/:id [post]
func (ctrl *TagController) Delete(c *gin.Context) {
	idString, _ := c.Params.Get("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(422, "参数错误")
		return
	}
	q := database.NewQuery()

	err = q.DeleteTag(database.DBCtx, int32(id))

	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, DeleteTagRes{
		Resource: int32(id),
	})
}

type TagUpdateReq struct {
	Id   int32  `json:"id" binding:"required"`
	Sign string `json:"sign" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type TagUpdateRes struct {
	Resource queries.Tag `json:"resource"`
}

// TagUpdate godoc
//
//		@Summary		tag
//		@Description	获取tag
//		@Tags			tag
//		@Accept			json
//		@Security		Bearer
//		@Produce		json
//	  @Param			body	body		TagUpdateReq	true	"body参数"
//		@Success		200	{object}	TagUpdateRes
//		@Failure		500
//		@Router			/tag [post]
func (ctrl *TagController) Update(c *gin.Context) {
	req := TagUpdateReq{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Print("入参错误", err)
		c.String(http.StatusUnprocessableEntity, "参数错误")
		return
	}

	q := database.NewQuery()
	tag, err := q.UpdateTag(c, queries.UpdateTagParams{
		ID:        req.Id,
		Name:      req.Name,
		Sign:      req.Sign,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, CreateTagRes{
		Resource: tag,
	})
}

func (ctrl *TagController) Get(c *gin.Context) {}

type TagListReq struct{}

type TagListRes struct {
	Resource []queries.Tag `json:"resource"`
}

// TagList godoc
//
//	@Summary		tag
//	@Description	获取tag
//	@Tags			tag
//	@Accept			json
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	TagListRes
//	@Failure		500
//	@Router			/tag [get]
func (ctrl *TagController) getList(c *gin.Context) {
	res := TagListRes{}
	user, _ := middleware.GetMe(c)
	q := database.NewQuery()
	item, err := q.ListTag(c, user.ID)
	if err != nil {
		c.String(http.StatusInternalServerError, "服务器繁忙")
		return
	}
	res.Resource = item

	c.JSON(http.StatusOK, res)
}
