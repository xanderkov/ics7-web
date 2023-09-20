package controller

import (
	"github.com/gin-gonic/gin"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
)

// AddDoctor godoc
// @Summary      Add a Doctor
// @Description  add Doctor by json
// @Tags         doctors
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
