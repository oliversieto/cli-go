package app

import "github.com/urfave/cli"

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada.
func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca Ips e Nomes de servidores na internet"
	return app
}
