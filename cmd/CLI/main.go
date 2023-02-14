package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Srgkharkov/anti-bruteforce/internal/pkg/cli"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("expected 'request', 'clear', 'addwl', 'delwl', 'addbl' or 'delbl'")
		os.Exit(1)
	}
	command := strings.ToLower(os.Args[1])
	flagset := flag.NewFlagSet("AccessRequest", flag.ExitOnError)
	addr := flagset.String("srvaddr", "localhost", "Server IP address")
	port := flagset.Int("port", 50051, "The anti-bruteforce port")
	login := flagset.String("login", "", "The login")
	pass := flagset.String("pass", "", "The password")
	ip := flagset.String("ip", "", "IP")
	subnet := flagset.String("subnet", "", "SubNet")
	flagset.Parse(os.Args[2:])
	if command != "request" && command != "clear" && command != "addwl" && command != "delwl" && command != "addbl" && command != "delbl" {
		fmt.Println("expected command: 'request', 'clear', 'addwl', 'delwl', 'addbl' or 'delbl'")
		os.Exit(1)
	}
	appCLI := cli.NewAppCLI(&command, addr, port, login, pass, ip, subnet)

	appCLI.Run()

}
