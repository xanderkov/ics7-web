package controller

import (
	"github.com/gin-gonic/gin"
	room_dto "hospital/internal/modules/domain/room/dto"
	"hospital/internal/modules/view/rest-api/httputil"
	"hospital/internal/modules/view/rest-api/model"
	"net/http"
	"strconv"
)

// ShowRoom godoc
// @Summary      Show a Room
// @Description  get string by ID
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Room ID"
// @Success      200  {object}  dto.Room
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /rooms/{id} [get]
func (c *Controller) ShowRoom(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)

	room, err := c.roomService.GetByNum(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, room)
}

// ListRooms godoc
//
//	@Summary		List Rooms
//	@Description	get Rooms
//	@Tags			Rooms
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		dto.Room
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/rooms [get]
func (c *Controller) ListRooms(ctx *gin.Context) {
	//q := ctx.Request.URL.Query().Get("q")
	rooms, err := c.roomService.List(ctx)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, rooms)
}

// UpdateRoom godoc
//
//	@Summary		Update a room
//	@Description	Update by json room
//	@Tags			Rooms
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Room ID"
//	@Param			room	body		model.UpdateRoom	true	"Update room"
//	@Success		200		{object}	dto.Room
//	@Failure		400		{object}	httputil.HTTPError
//	@Failure		404		{object}	httputil.HTTPError
//	@Failure		500		{object}	httputil.HTTPError
//	@Router			/rooms/{id} [patch]
func (c *Controller) UpdateRoom(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	var updateRoom model.UpdateRoom
	if err := ctx.ShouldBindJSON(&updateRoom); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}

	room := &room_dto.UpdateRoom{
		Num:            updateRoom.Num,
		Floor:          updateRoom.Floor,
		NumberBeds:     updateRoom.NumberBeds,
		TypeRoom:       updateRoom.TypeRoom,
		NumberPatients: updateRoom.NumberPatients,
	}
	roomUpdated, err := c.roomService.Update(ctx, aid, room)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusOK, roomUpdated)
}

// DeleteRoom godoc
//
//	@Summary		Delete a room
//	@Description	Delete by room ID
//	@Tags			Rooms
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Room ID"	Format(int64)
//	@Success		204	{object}	model.Room
//	@Failure		400	{object}	httputil.HTTPError
//	@Failure		404	{object}	httputil.HTTPError
//	@Failure		500	{object}	httputil.HTTPError
//	@Router			/rooms/{id} [delete]
func (c *Controller) DeleteRoom(ctx *gin.Context) {
	id := ctx.Param("id")
	aid, err := strconv.Atoi(id)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	err = c.roomService.Delete(ctx, aid)
	if err != nil {
		httputil.NewError(ctx, http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// AddRoom godoc
// @Summary      Add a Room
// @Description  add Room by json
// @Tags         Rooms
// @Accept       json
// @Produce      json
// @Param		 Room	body	model.AddRoom	true	"Add room"
// @Success      200  {object}  dto.Room
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /rooms [post]
func (c *Controller) AddRoom(ctx *gin.Context) {
	var addRoom model.AddRoom
	if err := ctx.ShouldBindJSON(&addRoom); err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	newRoom := &room_dto.CreateRoom{
		Num:            addRoom.Num,
		Floor:          addRoom.Floor,
		NumberBeds:     addRoom.NumberBeds,
		TypeRoom:       addRoom.TypeRoom,
		NumberPatients: addRoom.NumberPatients,
	}
	room, err := c.roomService.Create(ctx, newRoom)
	if err != nil {
		httputil.NewError(ctx, http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, room)
}
