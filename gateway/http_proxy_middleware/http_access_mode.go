package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hjldev/newmicro-mall/gateway/dao"
)

//匹配接入方式 基于请求信息
func HTTPAccessModeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, err := dao.ServiceManagerHandler.HTTPAccessMode(c)
		if err != nil {
			c.JSON(200, err)
			c.Set("response", err)
			c.Abort()
			return
		}
		//fmt.Println("matched service",public.Obj2Json(service))
		c.Set("service", service)
		c.Next()
	}
}
