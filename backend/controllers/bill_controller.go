package controllers

import (
	"context"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patza/app/ent"
	"github.com/patza/app/ent/bank"
	"github.com/patza/app/ent/confirmation"
	"github.com/patza/app/ent/user"
)

type BillController struct {
	client *ent.Client
	router gin.IRouter
}

type Bill struct {
	Paynumber    int
	Paytotal     int
	Bank         int
	Confirmation int
	User         int
	Paytime      string
}

// CreateBill handles POST requests for adding Bill entities
// @Summary Create bill
// @Description Create bill
// @ID create-bill
// @Accept   json
// @Produce  json
// @Param bill body ent.Bill true "Bill entity"
// @Success 200 {object} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [post]
func (ctl *BillController) CreateBill(c *gin.Context) {
	obj := Bill{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bill binding failed",
		})
		return
	}

	b, err := ctl.client.Bank.
		Query().
		Where(bank.IDEQ(int(obj.Bank))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bill not found",
		})
		return
	}

	con, err := ctl.client.Confirmation.
		Query().
		Where(confirmation.IDEQ(int(obj.Confirmation))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bill not found",
		})
		return
	}

	id, err := ctl.client.User.
		Query().
		Where(user.IDEQ(int(obj.User))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "bill not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Paytime)
	bi, err := ctl.client.Bill.
		Create().
		SetPayTime(time).
		SetPayTotal(obj.Paytotal).
		SetPayNumber(obj.Paynumber).
		SetBank(b).
		SetConfirmation(con).
		SetOwner(id).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, bi)
}

// ListBill handles request to get a list of bill entities
// @Summary List bill entities
// @Description list bill entities
// @ID list-bill
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bill
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /bills [get]
func (ctl *BillController) ListBill(c *gin.Context) {
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

	bills, err := ctl.client.Bill.
		Query().
		WithOwner().
		WithConfirmation().
		WithBank().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, bills)
}

// NewBillController creates and registers handles for the bill controller
func NewBillController(router gin.IRouter, client *ent.Client) *BillController {
	bc := &BillController{
		client: client,
		router: router,
	}

	bc.register()

	return bc

}

func (ctl *BillController) register() {
	bill := ctl.router.Group("/bills")

	bill.POST("", ctl.CreateBill)
	bill.GET("", ctl.ListBill)

}
