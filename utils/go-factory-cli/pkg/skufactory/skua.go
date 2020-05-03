package skufactory

import (
	"github.com/valmikroy/go-factory-cli/pkg/actions"
)

type SkuA struct{}

func (f *SkuA) NewSku(param Params) actions.Actions {
	return &actions.SkuA{
		Ipmi:      param.Ipmi,
		Payload:   param.Payload,
		Hibernate: param.Hibernate,
		Debug:     param.Debug,
	}
}
