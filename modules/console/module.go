package console

import (
	"github.com/liwei1dao/lego/base"
	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
	"github.com/liwei1dao/lego/lib/modules/console"
)

func NewModule() core.IModule {
	m := new(DataCollector)
	return m
}

type DataCollector struct {
	console.Console
	service base.IClusterService
	options IOptions
}

func (this *DataCollector) GetType() core.M_Modules {
	return lib.SM_ConsoleModule
}
