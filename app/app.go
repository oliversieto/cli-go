package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de servidores na internet"
	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Busca IPS de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: getIps,
		},
	}
	return app
}

func getIps(c *cli.Context) {
	host := c.String("host")
	ips, error := net.LookupIP(host)
	if error != nil {
		log.Fatal(error)
	}
	for _, ip := range ips {
		fmt.Println(ip)
	}
}
