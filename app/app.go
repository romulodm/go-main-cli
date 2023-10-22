package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Generate will return the command line application
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Control Panel Application"
	app.Usage = "Search IP's and Server Names"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "github.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Search for IP addresses on the web",
			Flags:  flags,
			Action: getIps,
		},
		{
			Name:   "servers",
			Usage:  "Search for Servers addresses on the web",
			Flags:  flags,
			Action: getServers,
		},
	}

	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")

	ips, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func getServers(c *cli.Context) {
	host := c.String("host")

	servers, err := net.LookupNS(host)
	if err != nil {
		log.Fatal(err)
	}

	for _, server := range servers {
		fmt.Println(server)
	}
}
