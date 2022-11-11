package demo

import (
	"context"
	"fmt"
	"lego_demo/comm"
	"lego_demo/pb"

	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/core/cbase"
	"github.com/liwei1dao/lego/sys/event"
	"github.com/liwei1dao/lego/sys/log"
	"github.com/liwei1dao/lego/sys/rpc"
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
	service base.IClusterService
}

//模块名
func (this *Demo) GetType() core.M_Modules {
	return comm.ModuleDemo
}

//模块初始化接口 注册用户创建角色事件
func (this *Demo) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	err = this.ModuleBase.Init(service, module, options)
	this.service = service.(base.IClusterService)
	return
}

func (this *Demo) Start() (err error) {
	err = this.ModuleBase.Start()
	this.service.RegisterFunctionName(comm.Rpc_ModuleDemoTest, this.Rpc_ModuleDemoTest)
	event.RegisterGO(rpc.Event_RpcDiscoverNewNodes, this.Event_RpcDiscoverNewNodes)
	return
}

func (this *Demo) Rpc_ModuleDemoTest(ctx context.Context, args *pb.DemoTestReq, reply *pb.DemoTestResp) (err error) {
	log.Debug("Rpc_ModulePayDelivery", log.Field{Key: "args", Value: args.String()})
	reply.Name = this.service.GetId()
	return
}

//发现新的服务节点时间
func (this *Demo) Event_RpcDiscoverNewNodes(nodes []*core.ServiceNode) {
	for _, v := range nodes {
		if v.Id != this.service.GetId() && v.Id == "demo2" {
			resp := &pb.DemoTestResp{}
			err := this.service.RpcCall(context.Background(), fmt.Sprintf("%s/%s", v.Type, v.Id), comm.Rpc_ModuleDemoTest, &pb.DemoTestReq{
				Name: this.service.GetId(),
			}, resp)
			log.Debug("Event_RpcDiscoverNewNodes", log.Field{Key: "resp", Value: resp.String()}, log.Field{Key: "err", Value: err})
		}
	}
}
