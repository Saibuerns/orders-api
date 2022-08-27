package server

import (
	"github.com/gin-gonic/gin"
	"orders.api/pkg/controller/order"
)

type Mapping struct {
	orderCtrl order.Controller
}

func NewMapping() *Mapping {
	return &Mapping{
		orderCtrl: resolveOrderController(),
	}
}

func (m Mapping) mapURLsToControllers(router *gin.Engine) {
	webGroup := router.Group("/orders")
	{
		webGroup.POST("", m.orderCtrl.Post)
	}
}
