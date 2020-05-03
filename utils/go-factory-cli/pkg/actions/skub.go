package actions

import (
	"fmt"
)

type SkuB struct {
	Ipmi      string
	Hibernate bool
	Debug     bool
}

func (sku *SkuB) Reboot() error {
	fmt.Printf("Host rebooted based only on IPMI IP %s\n", sku.Ipmi)
	return nil
}

func (sku *SkuB) Shutdown() error {
	if sku.Hibernate {
		fmt.Printf("Host hibernated based only on IPMI IP %s\n", sku.Ipmi)
	} else {

		fmt.Printf("Host rebooted based only on IPMI IP %s\n", sku.Ipmi)
	}
	return nil
}
