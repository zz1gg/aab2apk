package main

import (
	"aab2apk/internal"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var target string
var outputname string

func main() {

	cli.AppHelpTemplate = `{{if .Version}} {{.Name}} VERSION: {{.Version}} {{end}}
                    /$$        /$$$$$$                      /$$      
                    | $$       /$$__  $$                    | $$      
  /$$$$$$   /$$$$$$ | $$$$$$$ |__/  \ $$  /$$$$$$   /$$$$$$ | $$   /$$
 |____  $$ |____  $$| $$__  $$  /$$$$$$/ |____  $$ /$$__  $$| $$  /$$/
  /$$$$$$$  /$$$$$$$| $$  \ $$ /$$____/   /$$$$$$$| $$  \ $$| $$$$$$/ 
 /$$__  $$ /$$__  $$| $$  | $$| $$       /$$__  $$| $$  | $$| $$_  $$ 
|  $$$$$$$|  $$$$$$$| $$$$$$$/| $$$$$$$$|  $$$$$$$| $$$$$$$/| $$ \  $$
 \_______/ \_______/|_______/ |________/ \_______/| $$____/ |__/  \__/
Description: {{.Description}}                 |$$|               
                                                  |__/ {{if len .Authors}}AUTHOR:{{range .Authors}}{{ . }}{{end}}  

Usage: 
   aab2apk --extract test.aab --output test.apk
   aab2apk -e test.aab -o test.apk 
{{end}}{{if .Commands}}COMMANDS:
 {{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
OPTIONS:
  {{range .VisibleFlags}}{{.}}
  {{end}}{{end}}
`

	app := &cli.App{
		Description: "extract apk from aab",
		Version:     "v0.0.1",
		Authors: []*cli.Author{
			&cli.Author{
				Name: "zz1gg",
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "extract",
				Aliases:     []string{"e"},
				Usage:       "aab extract test.apk",
				Destination: &target,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "output",
				Aliases:     []string{"o"},
				Usage:       "output apk's name",
				Destination: &outputname,
			},
		},
		Action: func(c *cli.Context) error {
			if target == "" {
				log.Println("There is no target abb file")
			} else {
				log.Println("Target abb file：", target)
				internal.ExtractAPK(target, outputname)
			}
			//
			return nil
		},
	}

	err := app.Run(os.Args) //app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
