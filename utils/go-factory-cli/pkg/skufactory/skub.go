package skufactory

import (
	"github.com/valmikroy/go-factory-cli/pkg/actions"
)

type SkuB struct{}

func (f *SkuB) NewSku(param Params) actions.Actions {
	return &actions.SkuB{
		Ipmi:      param.Ipmi,
		Hibernate: param.Hibernate,
		Debug:     param.Debug,
	}
}
