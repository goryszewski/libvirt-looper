package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"libvirt.org/libvirt-go"
)

func main() {
	fmt.Println("Init Looper")
	conn, err := libvirt.NewConnect("qemu:///system")
	if err != nil {
		fmt.Println(err)
	}
	for true {
		doms, err := conn.ListAllDomains(libvirt.CONNECT_LIST_DOMAINS_ACTIVE)
		if err != nil {
			fmt.Println(err)
		}

		for _, item := range doms {
			// bo, err := item.IsActive()
			// fmt.Printf("[%v] [%v]\n", bo, err)
			name, err := item.GetName()
			if !strings.Contains(name, "autok8s") {
				continue
			}
			fmt.Printf("---------------[%v] [%v]\n", name, err)
			fget_ip(name)
			fget_sshauth_key(name)

			fadd_sshauth_key(name)
			get_hostname(name)
			set_hostname(name)
			os.Exit(0)
			// hostname, err := item.GetHostname(libvirt.DOMAIN_GET_HOSTNAME_AGENT)
			// fmt.Printf("[%v] [%v]\n", hostname, err)

			// state, i, err := item.GetState()
			// fmt.Printf("[%v][%v] [%v]\n", state, i, err)

			// xml, err := item.GetXMLDesc(libvirt.DOMAIN_XML_SECURE)
			// fmt.Printf("[%v] [%v]\n", xml, err)

		}

		time.Sleep(time.Second * 5)
	}
	defer conn.Close()
}
