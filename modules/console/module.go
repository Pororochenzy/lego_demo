package console

import (
	"lego_demo/pb"
	"lego_demo/sys/db"

	"github.com/liwei1dao/lego/core"
	"github.com/liwei1dao/lego/lib"
	"github.com/liwei1dao/lego/lib/modules/console"
	"github.com/liwei1dao/lego/sys/event"
)

func NewModule() core.IModule {
	m := new(Console)
	return m
}

type Console struct {
	console.Console
	options        *Options
	articleApiComp *ArticleApiComp
}

func (this *Console) GetType() core.M_Modules {
	return lib.SM_ConsoleModule
}

func (this *Console) NewOptions() (options core.IModuleOptions) {
	return new(Options)
}

func (this *Console) Init(service core.IService, module core.IModule, options core.IModuleOptions) (err error) {
	this.options = options.(*Options)
	err = this.Console.Init(service, module, options)
	return
}

func (this *Console) Start() (err error) {
	err = this.Console.Start()
	event.RegisterGO(console.Event_Registered, this.NewUserRegistered)
	return
}

func (this *Console) OnInstallComp() {
	this.Console.OnInstallComp()
	this.articleApiComp = this.RegisterComp(new(ArticleApiComp)).(*ArticleApiComp)
}

// 新用户注册
func (this *Console) NewUserRegistered(user *console.DB_UserData) {
	db.CreateAuthor(&pb.DB_AuthorData{
		AuthoIrd:    user.Id,
		PhonOrEmail: user.PhonOrEmail,
		Password:    user.Password,
		NickName:    user.NickName,
		HeadUrl:     user.HeadUrl,
		ArticleIds:  make([]uint32, 0),
	})
}
