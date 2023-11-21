package controller

import (
	"github.com/gin-gonic/gin"
	Account_dto "hospital/internal/modules/domain/auth/dto"
	auth_dto "hospital/internal/modules/domain/auth/dto"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
	"strconv"
)

// SingUp AddAccount godoc
// @Summary      Add a Doctor
// @Description  add Doctor by json
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param		 Doctor	body	model.AddDoctor	true	"Add Doctor"
// @Success      200  {object}  dto.Doctor
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/old [post]
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

// ShowAccount godoc
// @Summary      Show an Account
// @Description  get string by ID
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  dto.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]
func (c *Controller) ShowAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	Account, err := c.accountService.GetById(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, Account)
}

// ListAccount godoc
//
//	@Summary		List Account
//	@Description	get Account
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts [get]
func (c *Controller) ListAccount(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	Account, err := c.accountService.List(ctx)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, Account)
}

// UpdateAccount godoc
//
//	@Summary		Update a Account
//	@Description	Update by json Account
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Account ID"
//	@Param			Account	body		model.UpdateAccount	true	"Update Account"
//	@Success		200		{object}	dto.Account
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/accounts/{id} [patch]
func (c *Controller) UpdateAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateAccount model.UpdateAccount
	if err := ctx.ShouldBindJSON(&updateAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	Account := &Account_dto.UpdateAccount{
		Login:    updateAccount.Login,
		Password: updateAccount.PasswordHash,
	}
	AccountUpdated, err := c.accountService.Update(ctx, aid, Account)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, AccountUpdated)
}

// DeleteAccount godoc
//
//	@Summary		Delete a Account
//	@Description	Delete by Account ID
//	@Tags			Account
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"	Format(int64)
//	@Success		204	{object}	model.Account
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/accounts/{id} [delete]
func (c *Controller) DeleteAccount(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.accountService.Delete(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// AddAccount godoc
// @Summary      Add an Account
// @Description  add Account by json
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param		 Account	body	model.AddAccount	true	"Add Account"
// @Success      200  {object}  dto.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts [post]
func (c *Controller) AddAccount(ctx *gin.Context) {
	var addAccount model.AddAccount
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	newAccount := &Account_dto.CreateAccount{
		Login:    addAccount.Login,
		Password: addAccount.PasswordHash,
	}
	Account, err := c.accountService.Create(ctx, newAccount)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, Account)
}

// LoginAccount godoc
// @Summary      Login an Account
// @Description  login Account by json
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param		 Account	body	model.AddAccount	true	"Login Account"
// @Success      200  {object}  dto.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/login [post]
func (c *Controller) LoginAccount(ctx *gin.Context) {
	var addAccount model.AddAccount
	if err := ctx.ShouldBindJSON(&addAccount); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	loginAccount := &Account_dto.CreateAccount{
		Login:    addAccount.Login,
		Password: addAccount.PasswordHash,
	}
	_, err := c.accountService.Login(ctx, loginAccount)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
