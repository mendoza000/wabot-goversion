package main

import (
	_ "embed"
	"flag"
	"fmt"
	"github.com/fatih/color"
	_ "github.com/mattn/go-sqlite3"
	"os"
	"wabot/pkgs/client"
	"wabot/pkgs/utils"
)

func main() {
	groups := flag.Bool("groups", false, "Lista de todos lo grupos de WhatsApp")
	c15 := flag.Bool("c15", false, "Envia mensaje para cobrarle a los clientes del dia 15")
	c15r := flag.Bool("c15r", false, "Envia mensaje para recordarle a los clientes del dia 15 que deben pagar")
	c25 := flag.Bool("c25", false, "Envia mensaje para cobrarle a los clientes del dia 25")
	c25r := flag.Bool("c25r", false, "Envia mensaje para recordarle a los clientes del dia 25 que deben pagar")

	flag.Usage = func() {
		_, err := fmt.Fprintln(os.Stderr, "Options:")
		if err != nil {
			return
		}
		flag.VisitAll(func(f *flag.Flag) {
			utils.PrintFlag(f.Name, f.Usage)
		})
	}

	flag.Parse()

	if flag.NFlag() == 0 {
		color.Red("\nError: No flag provided.\n\n")
		flag.Usage()
		os.Exit(1)
	}

	if *groups {
		color.Green("\nGetting groups...\n\n")
		client.ConnectClient(1)
	}

	if *c15 {
		color.Green("Sending messages of 15h clients...\n")
		client.ConnectClient(2)
	}

	if *c15r {
		color.Yellow("Sending remember messages of 15th clients...\n")
		client.ConnectClient(3)
	}

	if *c25 {
		color.Green("Sending messages of 25th clients...\n")
		client.ConnectClient(4)
	}

	if *c25r {
		color.Yellow("Sending remember messages of 25th clients...\n")
		client.ConnectClient(5)
	}
}
