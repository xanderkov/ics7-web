package controller

import (
	"github.com/gin-gonic/gin"
	treatment_dto "hospital/internal/modules/domain/treatment/dto"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
	"strconv"
)

// ShowTreatment godoc
// @Summary      Show a Treatment
// @Description  get string by ID
// @Tags         Treatments
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Treatment ID"
// @Success      200  {object}  dto.Treatment
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /treatments/{id} [get]
func (c *Controller) ShowTreatment(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	treatment, err := c.treatmentService.GetById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, treatment)
}

// ListTreatments godoc
//
//	@Summary		List Treatments
//	@Description	get Treatments
//	@Tags			Treatments
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.Treatment
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/treatments [get]
func (c *Controller) ListTreatments(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	treatments, err := c.treatmentService.List(ctx)

	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, treatments)
}

// UpdateTreatment godoc
//
//	@Summary		Update a treatment
//	@Description	Update by json treatment
//	@Tags			Treatments
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Treatment ID"
//	@Param			treatment	body		model.UpdateTreatment	true	"Update treatment"
//	@Success		200		{object}	dto.Treatment
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/treatments/{id} [patch]
func (c *Controller) UpdateTreatment(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateTreatment model.UpdateTreatment
	if err := ctx.ShouldBindJSON(&updateTreatment); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	treatment := &treatment_dto.UpdateTreatment{
		PsychologicalTreatment: updateTreatment.PsychologicalTreatment,
		Tablets:                updateTreatment.Tablets,
		Survey:                 updateTreatment.Survey,
	}
	treatmentUpdated, err := c.treatmentService.Update(ctx, aid, treatment)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, treatmentUpdated)
}

// DeleteTreatment godoc
//
//	@Summary		Delete a treatment
//	@Description	Delete by treatment ID
//	@Tags			Treatments
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Treatment ID"	Format(int64)
//	@Success		204	{object}	model.Treatment
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/treatments/{id} [delete]
func (c *Controller) DeleteTreatment(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.treatmentService.Delete(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// AddTreatment godoc
// @Summary      Add a Treatment
// @Description  add Treatment by json
// @Tags         Treatments
// @Accept       json
// @Produce      json
// @Param		 Treatment	body	model.AddTreatment	true	"Add treatment"
// @Success      200  {object}  dto.Treatment
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /treatments [post]
func (c *Controller) AddTreatment(ctx *gin.Context) {
	var addTreatment model.AddTreatment
	if err := ctx.ShouldBindJSON(&addTreatment); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	newTreatment := &treatment_dto.CreateTreatment{
		PsychologicalTreatment: addTreatment.PsychologicalTreatment,
		Tablets:                addTreatment.Tablets,
		Survey:                 addTreatment.Survey,
	}
	treatment, err := c.treatmentService.Create(ctx, newTreatment)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, treatment)
}
