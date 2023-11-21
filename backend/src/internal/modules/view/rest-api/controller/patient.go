package controller

import (
	"github.com/gin-gonic/gin"
	patient_dto "hospital/internal/modules/domain/patient/dto"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
	"strconv"
)

// ShowPatient godoc
// @Summary      Show a Patient
// @Description  get string by ID
// @Tags         Patients
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Patient ID"
// @Success      200  {object}  dto.Patient
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /patients/{id} [get]
func (c *Controller) ShowPatient(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	patient, err := c.patientService.GetById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

// ShowPatientDisease godoc
// @Summary      Show a Patient Disease
// @Description  get string by ID
// @Tags         Patients
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Patient ID"
// @Success      200  {object}  dto.Disease
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /patientsDisease/{id} [get]
func (c *Controller) ShowPatientDisease(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	patient, err := c.patientService.GetDiseaseById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

// ShowPatientThreat godoc
// @Summary      Show a Patient Treatment
// @Description  get string by ID
// @Tags         Patients
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Patient ID"
// @Success      200  {object}  dto.Treatment
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /patientsThreat/{id} [get]
func (c *Controller) ShowPatientThreat(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	patient, err := c.patientService.GetThreatById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, patient)
}

// AddDiseasePatient  godoc
// @Summary      Add a Disease
// @Description  add Disease by json
// @Tags         Patients
// @Accept       json
// @Produce      json
// @Param		 Doctor	body	model.AddDiseaseToPatient	true	"Add doctorsPatients"
// @Success      200  {object}  dto.Disease
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /patientsDisease [post]
func (c *Controller) AddDiseasePatient(ctx *gin.Context) {

	var addAccount model.AddDiseaseToPatient
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	doctor, err := c.patientService.AddDiseaseById(ctx, addAccount.PatientId, addAccount.DiseaseId)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, doctor)
}

// ListPatients godoc
//
//	@Summary		List Patients
//	@Description	get Patients
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.Patient
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/patients [get]
func (c *Controller) ListPatients(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	patients, err := c.patientService.List(ctx)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, patients)
}

// UpdatePatient godoc
//
//	@Summary		Update a patient
//	@Description	Update by json patient
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Patient ID"
//	@Param			patient	body		model.UpdatePatient	true	"Update patient"
//	@Success		200		{object}	dto.Patient
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/patients/{id} [patch]
func (c *Controller) UpdatePatient(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updatePatient model.UpdatePatient
	if err := ctx.ShouldBindJSON(&updatePatient); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	patient := &patient_dto.UpdatePatient{
		Surname:        updatePatient.Surname,
		Name:           updatePatient.Name,
		Patronymic:     updatePatient.Patronymic,
		Weight:         updatePatient.Weight,
		Height:         updatePatient.Height,
		RoomNumber:     updatePatient.RoomNumber,
		DegreeOfDanger: updatePatient.DegreeOfDanger,
	}
	patientUpdated, err := c.patientService.Update(ctx, aid, patient)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, patientUpdated)
}

// DeletePatient godoc
//
//	@Summary		Delete a patient
//	@Description	Delete by patient ID
//	@Tags			Patients
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Patient ID"	Format(int64)
//	@Success		204	{object}	model.Patient
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/patients/{id} [delete]
func (c *Controller) DeletePatient(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.patientService.Delete(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// AddPatient godoc
// @Summary      Add a Patient
// @Description  add Patient by json
// @Tags         Patients
// @Accept       json
// @Produce      json
// @Param		 Patient	body	model.AddPatient	true	"Add patient"
// @Success      200  {object}  dto.Patient
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /patients [post]
func (c *Controller) AddPatient(ctx *gin.Context) {
	var addPatient model.AddPatient
	if err := ctx.ShouldBindJSON(&addPatient); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	newPatient := &patient_dto.CreatePatient{
		Surname:        addPatient.Surname,
		Name:           addPatient.Name,
		Patronymic:     addPatient.Patronymic,
		Weight:         addPatient.Weight,
		Height:         addPatient.Height,
		RoomNumber:     addPatient.RoomNumber,
		DegreeOfDanger: addPatient.DegreeOfDanger,
	}
	patient, err := c.patientService.Create(ctx, newPatient)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, patient)
}
