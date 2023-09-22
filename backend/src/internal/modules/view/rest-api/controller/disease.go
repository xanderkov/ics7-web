package controller

import (
	"github.com/gin-gonic/gin"
	disease_dto "hospital/internal/modules/domain/disease/dto"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
	"strconv"
)

// ShowDisease godoc
// @Summary      Show a Disease
// @Description  get string by ID
// @Tags         Disease
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Disease ID"
// @Success      200  {object}  dto.Disease
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /diseases/{id} [get]
func (c *Controller) ShowDisease(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	disease, err := c.diseaseService.GetById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, disease)
}

// ListDisease godoc
//
//	@Summary		List Disease
//	@Description	get Disease
//	@Tags			Disease
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.Disease
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/diseases [get]
func (c *Controller) ListDisease(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	disease, err := c.diseaseService.List(ctx)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, disease)
}

// UpdateDisease godoc
//
//	@Summary		Update a disease
//	@Description	Update by json disease
//	@Tags			Disease
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Disease ID"
//	@Param			disease	body		model.UpdateDisease	true	"Update disease"
//	@Success		200		{object}	dto.Disease
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/diseases/{id} [patch]
func (c *Controller) UpdateDisease(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateDisease model.UpdateDisease
	if err := ctx.ShouldBindJSON(&updateDisease); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	disease := &disease_dto.UpdateDisease{
		Name:           updateDisease.Name,
		Threat:         updateDisease.Threat,
		DegreeOfDanger: updateDisease.DegreeOfDanger,
	}
	diseaseUpdated, err := c.diseaseService.Update(ctx, aid, disease)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, diseaseUpdated)
}

// DeleteDisease godoc
//
//	@Summary		Delete a disease
//	@Description	Delete by disease ID
//	@Tags			Disease
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Disease ID"	Format(int64)
//	@Success		204	{object}	model.Disease
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/diseases/{id} [delete]
func (c *Controller) DeleteDisease(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.diseaseService.Delete(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// AddDisease godoc
// @Summary      Add a Disease
// @Description  add Disease by json
// @Tags         Disease
// @Accept       json
// @Produce      json
// @Param		 Disease	body	model.AddDisease	true	"Add disease"
// @Success      200  {object}  dto.Disease
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /diseases [post]
func (c *Controller) AddDisease(ctx *gin.Context) {
	var addDisease model.AddDisease
	if err := ctx.ShouldBindJSON(&addDisease); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	newDisease := &disease_dto.CreateDisease{
		Name:           addDisease.Name,
		Threat:         addDisease.Threat,
		DegreeOfDanger: addDisease.DegreeOfDanger,
	}
	disease, err := c.diseaseService.Create(ctx, newDisease)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, disease)
}
