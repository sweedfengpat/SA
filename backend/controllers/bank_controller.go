package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/patza/app/ent"
	"github.com/patza/app/ent/bank"
)

type BankController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateBank handles POST requests for adding bank entities
// @Summary Create bank
// @Description Create bank
// @ID create-bank
// @Accept   json
// @Produce  json
// @Param bank body ent.Bank true "Bank entity"
// @Success 200 {object} ent.Bank
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks [post]
func (ctl *BankController) CreateBank(c *gin.Context) {
	obj := ent.Bank{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "bank binding failed",
		})
		return
	}

	r, err := ctl.client.Bank.
		Create().
		SetBank(obj.Bank).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, r)
}

// GetBank handles GET requests to retrieve a bank entity
// @Summary Get a bank entity by ID
// @Description get bank by ID
// @ID get-bank
// @Produce  json
// @Param id path int true "Bank ID"
// @Success 200 {object} ent.Bank
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks/{id} [get]
func (ctl *BankController) GetBank(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	r, err := ctl.client.Bank.
		Query().
		Where(bank.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, r)
}

// ListBank handles request to get a list of bank entities
// @Summary List bank entities
// @Description list bank entities
// @ID list-bank
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Bank
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /banks [get]
func (ctl *BankController) ListBank(c *gin.Context) {
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

	banks, err := ctl.client.Bank.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, banks)
}

// NewBankController creates and registers handles for the bank controller
func NewBankController(router gin.IRouter, client *ent.Client) *BankController {
	bc := &BankController{
		client: client,
		router: router,
	}

	bc.register()

	return bc

}

func (ctl *BankController) register() {
	bank := ctl.router.Group("/banks")

	bank.POST("", ctl.CreateBank)
	bank.GET(":id", ctl.GetBank)
	bank.GET("", ctl.ListBank)

}
