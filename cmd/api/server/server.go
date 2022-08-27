package server

import "github.com/gin-gonic/gin"

func NewGinServer(urlMapping *Mapping) *gin.Engine {
	router := gin.New()

	urlMapping.mapURLsToControllers(router)

	return router
}
