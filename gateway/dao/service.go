package dao

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hjldev/newmicro-mall/gateway/conf"
	"github.com/hjldev/newmicro-mall/gateway/public"
	"strings"
	"sync"
)

type ServiceDetail struct {
	Info        *ServiceInfo `json:"info" description:"基本信息"`
	HTTPRule    *HttpRule    `json:"http_rule" description:"http_rule"`
	LoadBalance *LoadBalance `json:"load_balance" description:"load_balance"`
}

var ServiceManagerHandler *ServiceManager

func init() {
	ServiceManagerHandler = NewServiceManager()
}

type ServiceManager struct {
	ServiceMap   map[string]*ServiceDetail
	ServiceSlice []*ServiceDetail
	Locker       sync.RWMutex
	init         sync.Once
	err          error
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{
		ServiceMap:   map[string]*ServiceDetail{},
		ServiceSlice: []*ServiceDetail{},
		Locker:       sync.RWMutex{},
		init:         sync.Once{},
	}
}

func (s *ServiceManager) HTTPAccessMode(c *gin.Context) (*ServiceDetail, error) {
	//1、前缀匹配 /abc ==> serviceSlice.rule
	//2、域名匹配 www.test.com ==> serviceSlice.rule
	//host c.Request.Host
	//path c.Request.URL.Path
	host := c.Request.Host
	host = host[0:strings.Index(host, ":")]
	path := c.Request.URL.Path
	for _, serviceItem := range s.ServiceSlice {
		if serviceItem.Info.LoadType != public.LoadTypeHTTP {
			continue
		}
		if serviceItem.HTTPRule.RuleType == public.HTTPRuleTypeDomain {
			if serviceItem.HTTPRule.Rule == host {
				return serviceItem, nil
			}
		}
		if serviceItem.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL {
			if strings.HasPrefix(path, serviceItem.HTTPRule.Rule) {
				return serviceItem, nil
			}
		}
	}
	return nil, errors.New("not matched service")
}

func (s *ServiceManager) LoadOnce() error {
	s.init.Do(func() {
		for _, item := range conf.Conf.Service {
			serviceInfo := &ServiceInfo{
				LoadType:    0,
				ServiceName: item.Name,
			}
			httpRule := &HttpRule{
				RuleType:     item.RuleType,
				Rule:         item.Rule,
				NeedHttps:    0,
				UrlRewrite:   "",
				NeedStripUri: item.NeedStripUri,
			}
			loadBalance := &LoadBalance{
				CheckTimeout:  item.CheckTimeout,
				CheckInterval: item.CheckInterval,
				RoundType:     item.RoundType,
				IpList:        item.IpList,
				WeightList:    item.WeightList,
			}
			serviceDetail := &ServiceDetail{
				Info:        serviceInfo,
				HTTPRule:    httpRule,
				LoadBalance: loadBalance,
			}
			s.ServiceMap[item.Name] = serviceDetail
			s.ServiceSlice = append(s.ServiceSlice, serviceDetail)
		}
	})
	return s.err
}
