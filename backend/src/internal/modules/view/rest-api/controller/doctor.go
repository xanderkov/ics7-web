package controller

import (
	"github.com/gin-gonic/gin"
	"hospital/internal/modules/view/rest-api/httputil"
	"net/http"
	"strconv"
)

// ShowDoctor godoc
// @Summary      Show an Doctor
// @Description  get string by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Doctor ID"
// @Success      200  {object}  dto.Doctor
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	doctor, err := c.doctorService.GetById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, doctor)
}
