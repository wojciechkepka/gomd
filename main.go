package main

import (
	"flag"
	"fmt"
	"gomd/mdserver"
	"strconv"
)

func main() {
	bindAddr := flag.String("bind-addr", "127.0.0.1", "Binding address")
	bindPort := flag.String("bind-port", "5001", "Binding port")
	dir := flag.String("dir", ".", "The directory to serve")
	theme := flag.String("theme", "gruvbox", "Available gruvbox/solarized")
	showHidden := flag.Bool("hidden", false, "Display hidden files")
	quiet := flag.Bool("quiet", false, "Hide info output. Only errors are displayed")
	help := flag.Bool("help", false, "Print help")
	flag.Parse()
	if *help {
		fmt.Println("gomd\n\nUSAGE:")
		flag.PrintDefaults()
		return
	}
	port, err := strconv.Atoi(*bindPort)
	if err != nil {
		panic(err)
	}
	md := mdserver.NewMdServer(*bindAddr, port, *dir, *theme, *showHidden, *quiet)

	md.Serve()

}
