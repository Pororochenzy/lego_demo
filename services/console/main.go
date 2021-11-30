package main

import (
	"flag"
	"lego_demo/modules/console"
	"lego_demo/services"

	"github.com/liwei1dao/lego"
	"github.com/liwei1dao/lego/base/cluster"
	"github.com/liwei1dao/lego/core"
)

var (
	conf = flag.String("conf", "./conf/demo.toml", "获取需要启动的服务配置文件")
)

func main() {
	flag.Parse()
	s := NewService(
		cluster.SetConfPath(*conf),
	)
	s.OnInstallComp( //装备组件
	)
	lego.Run(s, //运行模块
		console.NewModule(),
	)
}

func NewService(ops ...cluster.Option) core.IService {
	s := new(Service)
	s.Configure(ops...)
	return s
}

type Service struct {
	services.ServiceBase
}

func (this *Service) InitSys() {
	this.ServiceBase.InitSys()
}
