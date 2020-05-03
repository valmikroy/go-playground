package skufactory

import (
	"github.com/valmikroy/go-factory-cli/pkg/actions"
)

type Params struct {
	Ipmi      string
	Payload   int
	Hibernate bool
	Debug     bool
}

type SkuFactory interface {
	NewSku(param Params) actions.Actions
}
