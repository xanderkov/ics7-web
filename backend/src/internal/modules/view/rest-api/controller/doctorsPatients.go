package controller

import (
	"github.com/gin-gonic/gin"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
	"strconv"
)

// ShowDoctorsPatients godoc
// @Summary      Show Doctor's patients
// @Description  get string by ID
// @Tags         doctorsPatients
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Doctor ID"
// @Success      200  {object}  dto.Patient
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /doctorsPatients/{id} [get]
func (c *Controller) ShowDoctorsPatients(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	patients, err := c.doctorService.GetPatientsById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, patients)
}

// AddDoctorsPatients  godoc
// @Summary      Add a DoctorsPatients
// @Description  add DoctorsPatients by json
// @Tags         doctorsPatients
// @Accept       json
// @Produce      json
// @Param		 Doctor	body	model.AddDoctorsPatients	true	"Add doctorsPatients"
// @Success      200  {object}  dto.Patient
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /doctorsPatients [post]
func (c *Controller) AddDoctorsPatients(ctx *gin.Context) {

	var addAccount model.AddDoctorsPatients
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	doctor, err := c.doctorService.AddPatientsById(ctx, addAccount.DoctorId, addAccount.PatientId)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, doctor)
}
