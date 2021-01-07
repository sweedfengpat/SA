package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/patza/app/ent/confirmation"
	"github.com/patza/app/ent/user"

	"github.com/gin-gonic/gin"
	"github.com/patza/app/ent"
)

type ConfirmationController struct {
	client *ent.Client
	router gin.IRouter
}

type Confirmation struct {
	AddDate string
	BookingStart string
	BookingEnd   string
	HoursTime int
	Owner     int
}

// CreateConfirmation handles POST requests for adding confirmation entities
// @Summary Create vconfirmationideo
// @Description Create confirmation
// @ID create-confirmation
// @Accept   json
// @Produce  json
// @Param video body ent.Confirmation true "Confirmation entity"
// @Success 200 {object} ent.Confirmation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations [post]

func (ctl *ConfirmationController) CreateConfirmation(c *gin.Context) {
	obj := Confirmation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "confirmation binding failed",
		})
		return
	}
	time1, err := time.Parse(time.RFC3339, obj.AddDate)
	time2, err := time.Parse(time.RFC3339, obj.BookingStart)
	time3, err := time.Parse(time.RFC3339, obj.BookingEnd)
	con, err := ctl.client.Confirmation.
		Create().
		SetAdddate(time1).
		SetBookingstart(time2).
		SetBookingend(time3).
		SetHourstime(obj.HoursTime).
		SetOwnerID(obj.Owner).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving confirmation failed",
		})
		return
	}

	u, err := ctl.client.User.
		UpdateOneID(int(obj.Owner)).
		AddConfirmation(con).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving edge failed",
		})
		return
	}

	c.JSON(200, u)
}

// GetConfirmation handles GET requests to retrieve a confirmation entity
// @Summary Get a confirmation entity by ID
// @Description get confirmation by ID
// @ID get-confirmation
// @Produce  json
// @Param id path int true "Confirmation ID"
// @Success 200 {object} ent.Confirmation
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations/{id} [get]

func (ctl *ConfirmationController) GetConfirmations(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	v, err := ctl.client.Confirmation.
		Query().
		WithOwner().
		Where(confirmation.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, v)
}

// ListConfirmation handles request to get a list of confirmation entities
// @Summary List confirmation entities
// @Description list confirmation entities
// @ID list-confirmation
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Confirmation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations [get]
func (ctl *ConfirmationController) ListConfirmation(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	confirmations, err := ctl.client.Confirmation.
		Query().
		WithOwner().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, confirmations)
}

// DeleteConfirmation handles DELETE requests to delete a confirmation entity
// @Summary Delete a confirmation entity by ID
// @Description get confirmation by ID
// @ID delete-confirmation
// @Produce  json
// @Param id path int true "Confirmation ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations/{id} [delete]

func (ctl *ConfirmationController) DeleteConfirmation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Confirmation.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateConfirmation handles PUT requests to update a confirmation entity
// @Summary Update a confirmation entity by ID
// @Description update confirmation by ID
// @ID update-confirmation
// @Accept   json
// @Produce  json
// @Param id path int true "Confirmation ID"
// @Param confirmation body ent.Confirmation true "Confirmation entity"
// @Success 200 {object} ent.Confirmation
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /confirmations/{id} [put]

func (ctl *ConfirmationController) UpdateConfirmation(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := Confirmation{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "video binding failed",
		})
		return
	}

	u, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	time1, err := time.Parse(time.RFC3339, obj.AddDate)
	time2, err := time.Parse(time.RFC3339, obj.BookingStart)
	time3, err := time.Parse(time.RFC3339, obj.BookingEnd)
	con, err := ctl.client.Confirmation.
		Create().
		SetAdddate(time1).
		SetBookingstart(time2).
		SetBookingend(time3).
		SetHourstime(obj.HoursTime).
		SetOwner(u).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "update confirmation failed",
		})
		return
	}

	c.JSON(200, con)
}

func NewConfirmationController(router gin.IRouter, client *ent.Client) *ConfirmationController {
	conc := &ConfirmationController{
		client: client,
		router: router,
	}

	conc.register()

	return conc

}

func (ctl *ConfirmationController) register() {
	confirmation := ctl.router.Group("/confirmations")

	confirmation.POST("", ctl.CreateConfirmation)
	confirmation.GET(":id", ctl.GetConfirmations)
	confirmation.GET("", ctl.ListConfirmation)
	confirmation.DELETE(":id", ctl.DeleteConfirmation)
	confirmation.PUT(":id", ctl.UpdateConfirmation)

}
