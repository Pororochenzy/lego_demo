package helloworld

import (
	"context"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/event"
	"lego_demo/pb"
)

func NewModule() core.IModule {
	m := new(HellWorld)
	return m
}

type HellWorld struct {
	cbase.ModuleBase
	//options        *Options
	//articleApiComp *ArticleApiComp
}

func (this *HellWorld) GetType() core.M_Modules {
	return "HelloWorld"
}

func (this *HellWorld) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	//this.options = options.(*Options)
	err = this.ModuleBase.Init(service, module, options)
	return
}

func (this *HellWorld) Start() (err error) {
	err = this.ModuleBase.Start()
	event.RegisterGO("HD_Hello", this.Hello)
	return
}

func (this *HellWorld) Hello(ctx context.Context, args *pb.DemoTestReq, reply *pb.DemoTestResp) (err error) {
	reply = &pb.DemoTestResp{
		Name: "Hello success",
	}
	return
}
