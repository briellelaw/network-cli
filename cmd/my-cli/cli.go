package main

import (
	"fmt"
	"os"
	"log"
	"net"
	"github.com/urfave/cli"
)

func main() {
	app :=	cli.NewApp()
	app.Name = "Website Lookup CLI"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and Name Servers" 

	flag1 := cli.StringFlag {
		Name: "host", 
		Value: "tutorialedge.net",
	}

	myFlags := []cli.Flag {
		&flag1,
	}

	nsCommand := cli.Command {
		Name: "ns", 
		Usage: "Looks up the Name Servers for a Particular Host", 
		Flags: myFlags,
		Action: func(c *cli.Context) error {
			ns, err := net.LookupNS(c.String("host"))
			if err != nil {
				return err
			}
			for i := 0; i < len(ns); i++ {
				fmt.Println(ns[i].Host)
			}
			return nil
		},
	}

	ipCommand := cli.Command {
		Name: "ip", 
		Usage: "Looks up the IP for a Particular Host", 
		Flags: myFlags,
		Action: func(c *cli.Context) error {
			ip, err := net.LookupIP(c.String("host"))
			if err != nil {
				return err
			}
			for i := 0; i < len(ip); i++ {
				fmt.Println(ip[i])
			}
			return nil
		},
	}

	cnameCommand := cli.Command {
		Name: "cname", 
		Usage: "Looks up the CNAME for a Particular Host", 
		Flags: myFlags,
		Action: func(c *cli.Context) error {
			cname, err := net.LookupCNAME(c.String("host"))
			if err != nil {
				return err
			}
			fmt.Println(cname)
			return nil
		},
	}

	mxCommand := cli.Command {
		Name: "mx", 
		Usage: "Looks up the MX records for a Particular Host", 
		Flags: myFlags,
		Action: func(c *cli.Context) error {
			mx, err := net.LookupMX(c.String("host"))
			if err != nil {
				return err
			}
			for i :=0; i < len(mx); i++ {
				fmt.Println(mx[i].Host, mx[i].Pref)
			}
			return nil
		},
	}


	app.Commands = []*cli.Command {
		&nsCommand,
		&ipCommand,
		&cnameCommand,
		&mxCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}