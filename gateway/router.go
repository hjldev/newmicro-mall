package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hjldev/newmicro-mall/gateway/http_proxy_middleware"
)

func initRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.Default()
	router.Use(middlewares...)

	router.Use(
		http_proxy_middleware.HTTPAccessModeMiddleware(),
		http_proxy_middleware.HTTPStripUriMiddleware(),
		http_proxy_middleware.HTTPUrlRewriteMiddleware(),
		http_proxy_middleware.HTTPReverseProxyMiddleware())

	return router
}
