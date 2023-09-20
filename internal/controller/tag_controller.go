package controller

import (
	"log"
	"mangosteen/config/queries"
	"mangosteen/internal/api"
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
	item.DELETE("/:id", ctrl.Delete)
	item.GET("/:id", ctrl.Get)
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
// @Param			body	body		api.CreateTagReq	true	"body参数"
// @Success		200		{object}	api.CreateTagRes
// @Failure		500
// @Router			/tag [post]
func (ctrl *TagController) Create(c *gin.Context) {
	req := api.CreateTagReq{}

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

	c.JSON(http.StatusOK, api.CreateTagRes{
		Resource: item,
	})
}

// TagDelete godoc
//
//	@Summary		tag
//	@Description	获取tag
//	@Tags			tag
//	@Accept			json
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	api.DeleteTagRes
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

	tag, err := q.DeleteTag(database.DBCtx, int32(id))
	// int32(id)
	if err != nil {
		c.String(500, err.Error())
		return
	}
	c.JSON(http.StatusOK, api.DeleteTagRes{
		Resource: tag,
	})
}

// TagUpdate godoc
//
//		@Summary		tag
//		@Description	获取tag
//		@Tags			tag
//		@Accept			json
//		@Security		Bearer
//		@Produce		json
//	  @Param			body	body		api.TagUpdateReq	true	"body参数"
//		@Success		200	{object}	api.TagUpdateRes
//		@Failure		500
//		@Router			/tag [post]
func (ctrl *TagController) Update(c *gin.Context) {
	req := api.TagUpdateReq{}

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

	c.JSON(http.StatusOK, api.CreateTagRes{
		Resource: tag,
	})
}

// FindTag godoc
//
//	@Summary		tag
//	@Description	获取tag
//	@Tags			tag
//	@Accept			json
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	api.FindTagRes
//	@Failure		500
//	@Router			/tag/:id [get]
func (ctrl *TagController) Get(c *gin.Context) {
	idString, _ := c.Params.Get("id")
	user, _ := middleware.GetMe(c)

	id, err := strconv.Atoi(idString)
	if err != nil {
		c.String(422, "参数错误")
		return
	}
	q := database.NewQuery()

	tag, err := q.FindTag(database.DBCtx,
		queries.FindTagParams{
			ID:     int32(id),
			UserID: user.ID,
		})
	if err != nil {
		c.String(500, err.Error())
		return
	}

	c.JSON(http.StatusOK, api.DeleteTagRes{
		Resource: tag,
	})
}

// TagList godoc
//
//	@Summary		tag
//	@Description	获取tag
//	@Tags			tag
//	@Accept			json
//	@Security		Bearer
//	@Produce		json
//	@Success		200	{object}	api.TagListRes
//	@Failure		500
//	@Router			/tag [get]
func (ctrl *TagController) getList(c *gin.Context) {
	res := api.TagListRes{}
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
