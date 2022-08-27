package server

import (
	"github.com/gin-gonic/gin"
	"orders.api/pkg/controller/order"
	"orders.api/pkg/controller/product"
)

type Mapping struct {
	orderCtrl   order.Controller
	productCtrl product.Controller
}

func NewMapping() *Mapping {
	return &Mapping{
		orderCtrl:   resolveOrderController(),
		productCtrl: resolveProductController(),
	}
}

func (m Mapping) mapURLsToControllers(router *gin.Engine) {
	webGroup := router.Group("/orders")
	{
		webGroup.POST("", m.orderCtrl.Post)
	}
	webGroup = webGroup.Group("/products")
	{
		webGroup.POST("", m.productCtrl.Post)
		webGroup.GET("", m.productCtrl.GetAll)
	}
}
