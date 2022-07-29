package demo

import (
	"lego_demo/comm"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
)

/*
模块名:演示
描述:演示集群通信
开发:李伟
*/
func NewModule() core.IModule {
	m := new(Demo)
	return m
}

type Demo struct {
	cbase.ModuleBase
}

//模块名
func (this *Demo) GetType() core.M_Modules {
	return comm.ModuleDemo
}

//模块初始化接口 注册用户创建角色事件
func (this *Demo) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	err = this.ModuleBase.Init(service, module, options)
	return
}
