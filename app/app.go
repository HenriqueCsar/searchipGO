package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//generate rows
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Command line application = SEARCHIPGO"
	app.Usage = "Ip search on the internet"

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Search Ips Address",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "www.google.com",
				},
			},
			Action: searchIps,
		},
	}

	return app
}

func searchIps(c *cli.Context) {
	host := c.String("host")
	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}
