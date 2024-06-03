package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// retorna a aplicação para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "App Busca IP"
	app.Usage = "Busca IPs e Nomes de Servidor na internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Busca IPs de endereços",
			Flags:  flags,
			Action: buscarIPs,
		},
		{
			Name:   "servidores",
			Usage:  "Busca nome dos servidores",
			Flags:  flags,
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIPs(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal(erro)
	}

	for _, servidor := range servidores {
		fmt.Println(servidor.Host)
	}
}
