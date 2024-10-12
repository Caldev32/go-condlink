package main

import(
	"fmt"
	"log"
	"os"
	"github.com/urfave/cli/v2"
)

func main(){
	app := &cli.App{
		Name:"Healthchecker",
		Usage:"uma ferramenta que checa se um está no ar ou fora do ar",
		Flags:[]cli.Flag{
			&cli.StringFlag{
				Name: "dominio",
				Aliases: []string{"d"},
				Usage: "Dominio do sie a ser checado",
				Required: true,
			},
			&cli.StringFlag{
				Name: "port",
				Aliases: []string{"p"},
				Usage: "Numero da porta do site a ser checado",
				Required: false,
			},

		},
		Action: func(c *cli.Context) error{
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}
			status := Check(c.String("dominio"),port)
			fmt.Println(status)
			return nil
		},
		
	}
	err := app.Run(os.Args)
	if err != nil{
		log.Fatal(err)
	}
}