package console

import (
	"github.com/liwei1dao/lego/lib/modules/console"
	"github.com/liwei1dao/lego/utils/mapstructure"
)

type Options struct {
	console.Options
}

func (this *Options) LoadConfig(settings map[string]interface{}) (err error) {
	if err = this.Options.LoadConfig(settings); err == nil {
		if settings != nil {
			err = mapstructure.Decode(settings, this)
		}
	}
	return
}
