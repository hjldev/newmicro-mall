package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

type ServiceConfig struct {
	Name          string `json:"name"`
	RuleType      int    `json:"rule_type"`
	Rule          string `json:"rule"`
	CheckTimeout  int    `json:"check_timeout"`
	CheckInterval int    `json:"check_interval"`
	RoundType     int    `json:"round_type"`
	IpList        string `json:"ip_list"`
	WeightList    string `json:"weight_list"`
	UrlRewrite    string
	NeedStripUri  int
}

type Config struct {
	Service []ServiceConfig
}

var Conf Config

func init() {
	if _, err := toml.DecodeFile("conf/config.toml", &Conf); err != nil {
		log.Println(err)
	}
}
