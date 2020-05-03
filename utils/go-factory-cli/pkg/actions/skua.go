package actions

import (
	"fmt"
)

type SkuA struct {
	Ipmi      string
	Payload   int
	Hibernate bool
	Debug     bool
}

func (sku *SkuA) Reboot() error {
	fmt.Printf("Host rebooted based on IPMI IP %s and payloads %d\n", sku.Ipmi, sku.Payload)
	return nil
}

func (sku *SkuA) Shutdown() error {
	if sku.Hibernate {
		fmt.Printf("Host Hibernated based on IPMI IP %s and payloads %d\n", sku.Ipmi, sku.Payload)
	} else {
		fmt.Printf("Host Shutdown based on IPMI IP %s and payloads %d\n", sku.Ipmi, sku.Payload)
	}
	return nil
}
