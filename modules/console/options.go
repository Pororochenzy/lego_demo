package console

import (
	"github.com/liwei1dao/lego/lib/modules/console"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type (
	IOptions interface {
		console.IOptions
		GetSignKey() string
		GetScanNoStartRunnerInterval() string
		GetRinfoSyncInterval() string
		GetRStatisticInterval() string
	}
	Options struct {
		console.Options
		SignKey                   string //签名密钥
		ScanNoStartRunnerInterval string //扫描为启动采集任务间隔
		RinfoSyncInterval         string //RunerInfo 同步间隔
		RStatisticInterval        string //RunerInfo 统计间隔
	}
)

func (this *Options) GetSignKey() string {
	return this.SignKey
}

func (this *Options) GetScanNoStartRunnerInterval() string {
	return this.ScanNoStartRunnerInterval
}

func (this *Options) GetRinfoSyncInterval() string {
	return this.RinfoSyncInterval
}

func (this *Options) GetRStatisticInterval() string {
	return this.RStatisticInterval
}

func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	this.ScanNoStartRunnerInterval = "0 */10 * * * ?" //默认每十分钟执行一次
	this.RinfoSyncInterval = "0 */1 * * * ?"          //默认每分钟执行一次
	this.RStatisticInterval = "59 59 * * * ?"         //每小时的59分59秒执行
	if err = this.Options.LoadConfig(settings); err == nil {
		if settings != nil {
			err = mapstructure.Decode(settings, this)
		}
	}
	return
}
