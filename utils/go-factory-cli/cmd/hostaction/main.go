package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/valmikroy/go-factory-cli/pkg/skufactory"
	"log"
	"os"
)

func globalFlags() []cli.Flag {
	return []cli.Flag{

		&cli.BoolFlag{
			Name:  "debug",
			Usage: "Debug flag",
		},
		&cli.StringFlag{
			Name:     "sku",
			Usage:    "SKU type",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "ipmi",
			Usage:    "ipmi IP ",
			Required: true,
		},
		&cli.IntFlag{
			Name:  "payload",
			Usage: "Payload is SKU specific param",
		},
	}
}

func actionRebootFlags() []cli.Flag {
	return []cli.Flag{}
}

func actionShutdownFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  "hibernate",
			Usage: "Hibernate host",
		},
	}
}

func actionReboot() *cli.Command {

	var f []cli.Flag
	f = append(f, globalFlags()...)
	f = append(f, actionRebootFlags()...)

	return &cli.Command{
		Name:  "reboot",
		Usage: "Reboot host",
		Flags: f,
		Action: func(c *cli.Context) error {
			p := skufactory.Params{
				Ipmi:  c.String("ipmi"),
				Debug: c.Bool("debug"),
			}

			sku := c.String("sku")
			switch sku {
			case "skua":
				p.Payload = c.Int("payload")
				s := &skufactory.SkuA{}
				s.NewSku(p).Reboot()
			case "skub":
				s := &skufactory.SkuB{}
				s.NewSku(p).Reboot()
			default:
				return fmt.Errorf("Invalid SKU type %s", sku)
			}
			return nil
		},
	}
}

func actionShutdown() *cli.Command {
	f := []cli.Flag{}
	f = append(f, globalFlags()...)
	f = append(f, actionShutdownFlags()...)

	return &cli.Command{
		Name:  "shutdown",
		Usage: "Shutdown host",
		Flags: f,
		Action: func(c *cli.Context) error {
			p := skufactory.Params{
				Ipmi:      c.String("ipmi"),
				Debug:     c.Bool("debug"),
				Hibernate: c.Bool("hibernate"),
			}

			sku := c.String("sku")
			switch sku {
			case "skua":
				p.Payload = c.Int("payload")
				s := &skufactory.SkuA{}
				s.NewSku(p).Shutdown()
			case "skub":
				s := &skufactory.SkuB{}
				s.NewSku(p).Shutdown()
			default:
				return fmt.Errorf("Invalid SKU type %s", sku)
			}
			return nil
		},
	}
}

func main() {

	app := cli.NewApp()
	app.Name = "hostaction"
	app.Usage = "Host Action"

	app.Commands = []*cli.Command{
		actionReboot(),
		actionShutdown(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
