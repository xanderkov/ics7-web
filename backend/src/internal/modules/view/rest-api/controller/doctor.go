package controller

import (
	"github.com/gin-gonic/gin"
	"hospital/internal/modules/view/rest-api/httputil"
	"net/http"
	"strconv"
)

// ShowDoctor godoc
// @Summary      Show a Doctor
// @Description  get string by ID
// @Tags         doctors
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
