package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p ProductCatalog) MerchandiseRegister(router *gin.RouterGroup) {
	router.POST("/updateMerchandise", p.updateMerchandise)
	router.POST("/createMerchandise", p.createMerchandise)
	router.GET("/getMerchandises", p.getMerchandises)
}

func (p ProductCatalog) getMerchandises(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

func (p ProductCatalog) createMerchandise(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}

func (p ProductCatalog) updateMerchandise(c *gin.Context) {

	c.JSON(http.StatusOK, nil)
}
