package mdserver

import (
	"flag"
	"fmt"
	"os"
)

//Default values
const (
	DefAddr       = "127.0.0.1"
	DefPort       = "5001"
	DefDebug      = false
	DefDir        = "."
	DefTheme      = "solarized"
	DefShowHidden = false
	DefQuiet      = false
	DefNoOpen     = false
	DefHelp       = false
	Version       = "1.0.0"
)

/*MdOpts Options for running MdServer*/
type MdOpts struct {
	BindAddr, BindPort, Dir, Theme         *string
	Debug, ShowHidden, Quiet, NoOpen, help *bool
}

//ParseMdOpts parses provided commandline options returning MdOpts
func ParseMdOpts() MdOpts {
	defer flag.Parse()
	return MdOpts{
		BindAddr:   flag.String("bind-addr", DefAddr, "Binding address"),
		BindPort:   flag.String("bind-port", DefPort, "Binding port"),
		Debug:      flag.Bool("debug", DefDebug, "Display debug output. Overrides all other flags"),
		Dir:        flag.String("dir", DefDir, "The directory to serve"),
		Theme:      flag.String("theme", DefTheme, "Available dracula/paraiso/monokai/solarized/github/vs/xcode"),
		ShowHidden: flag.Bool("hidden", DefShowHidden, "Display hidden files"),
		Quiet:      flag.Bool("quiet", DefQuiet, "Hide info output. Only errors are displayed"),
		NoOpen:     flag.Bool("no-open", DefNoOpen, "Don't open new browser tab."),
		help:       flag.Bool("help", DefHelp, "Print help"),
	}
}

//CheckHelp exits with status code 1 if help flag is provided printing defaults
func (opts *MdOpts) CheckHelp() {
	if *opts.help {
		fmt.Println("gomd -", Version, "\n\nUSAGE:")
		flag.PrintDefaults()
		os.Exit(1)
	}
}
