package main

import (
	"flag"
	"lego_demo/modules/live"
	"lego_demo/services"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
)

var (
	conf = flag.String("conf", "./conf/live.toml", "获取需要启动的服务配置文件") //启动服务的Id
)

func main() {
	flag.Parse()
	s := NewService(
		cluster.SetConfPath(*conf),
	)
	s.OnInstallComp( //装备组件
	)
	lego.Run(s,
		live.NewModule(),
	) //运行模块

}

func NewService(ops ...cluster.Option) core.IService {
	s := new(Demo2Service)
	s.Configure(ops...)
	return s
}

type Demo2Service struct {
	services.ServiceBase
}

func (this *Demo2Service) InitSys() {
	this.ServiceBase.InitSys()
}
