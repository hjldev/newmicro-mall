package http_proxy_middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hjldev/newmicro-mall/gateway/dao"
	"github.com/hjldev/newmicro-mall/gateway/public"
	"strings"
)

//匹配接入方式 基于请求信息
func HTTPStripUriMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			c.JSON(2001, errors.New("service not found"))
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		if serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HTTPRule.NeedStripUri == 1 {
			//fmt.Println("c.Request.URL.Path",c.Request.URL.Path)
			c.Request.URL.Path = strings.Replace(c.Request.URL.Path, serviceDetail.HTTPRule.Rule, "", 1)
			//fmt.Println("c.Request.URL.Path",c.Request.URL.Path)
		}

		c.Next()
	}
}
