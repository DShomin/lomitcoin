package cli

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	explorer "github.com/DShomin/lomitcoin/explorer/templates"
	"github.com/DShomin/lomitcoin/rest"
)

func usage() {
	fmt.Println("Welcome to 로밋코인")
	fmt.Println("")
	fmt.Println("Please use the following commands:")
	fmt.Println("")
	fmt.Println("-port : 	Set port of the serve")
	fmt.Println("-mode : 	Choose mode 'html', 'rest', 'dual'")
	runtime.Goexit()
}

func Start() {
	if len(os.Args) == 1 {
		usage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose mode 'html', 'rest', 'dual'")

	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "dual":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default:
		usage()
	}
}
