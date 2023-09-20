package controller

import (
	"github.com/gin-gonic/gin"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	doctor_dto "hospital/internal/modules/domain/doctor/dto"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
	"strconv"
)

// ShowDoctor godoc
// @Summary      Show a Doctor
// @Description  get string by ID
// @Tags         Doctors
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Doctor ID"
// @Success      200  {object}  dto.Doctor
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /doctors/{id} [get]
func (c *Controller) ShowDoctor(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	doctor, err := c.doctorService.GetById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, doctor)
}

// ListDoctors godoc
//
//	@Summary		List Doctors
//	@Description	get Doctors
//	@Tags			Doctors
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.Doctor
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/doctors [get]
func (c *Controller) ListDoctors(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	doctors, err := c.doctorService.List(ctx)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, doctors)
}

// SingUp AddDoctor godoc
// @Summary      Add a Doctor
// @Description  add Doctor by json
// @Tags         Doctors
// @Accept       json
// @Produce      json
// @Param		 Doctor	body	model.AddDoctor	true	"Add doctor"
// @Success      200  {object}  dto.Doctor
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /doctors [post]
func (c *Controller) SingUp(ctx *gin.Context) {

	var addAccount model.AddDoctor
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	newDoctor := &auth_dto.NewDoctor{
		TokenId:    addAccount.TokenId,
		Surname:    addAccount.Surname,
		Speciality: addAccount.Speciality,
		Role:       addAccount.Role,
	}

	doctor, err := c.authService.SignUp(ctx, newDoctor)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, doctor)
}

// UpdateDoctor godoc
//
//	@Summary		Update a doctor
//	@Description	Update by json doctor
//	@Tags			Doctors
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Doctor ID"
//	@Param			doctor	body		model.UpdateDoctor	true	"Update doctor"
//	@Success		200		{object}	dto.Doctor
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/doctors/{id} [patch]
func (c *Controller) UpdateDoctor(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateDoctor model.UpdateDoctor
	if err := ctx.ShouldBindJSON(&updateDoctor); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	doctor := &doctor_dto.UpdateDoctor{
		TokenId:    updateDoctor.TokenId,
		Surname:    updateDoctor.Surname,
		Speciality: updateDoctor.Speciality,
		Role:       updateDoctor.Role,
	}
	doctorUpdated, err := c.doctorService.Update(ctx, aid, doctor)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, doctorUpdated)
}

// DeleteDoctor godoc
//
//	@Summary		Delete a doctor
//	@Description	Delete by doctor ID
//	@Tags			Doctors
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Doctor ID"	Format(int64)
//	@Success		204	{object}	model.Doctor
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/doctors/{id} [delete]
func (c *Controller) DeleteDoctor(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.doctorService.Delete(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}
