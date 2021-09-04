package http_proxy_middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hjldev/newmicro-mall/gateway/dao"
	"github.com/hjldev/newmicro-mall/gateway/reverse_proxy"
)

//匹配接入方式 基于请求信息
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			c.JSON(2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		lb, err := dao.LoadBalancerHandler.GetLoadBalancer(serviceDetail)
		if err != nil {
			c.JSON(2002, err)
			c.Abort()
			return
		}
		trans, err := dao.TransportorHandler.GetTrans(serviceDetail)
		if err != nil {
			c.JSON(2003, err)
			c.Abort()
			return
		}
		proxy := reverse_proxy.NewLoadBalanceReverseProxy(c, lb, trans)
		proxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return
	}
}
