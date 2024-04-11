package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"time"

	"libvirt.org/libvirt-go"
)

type inter struct {
	IPType   string `json:"ip-address-type"`
	IpAddres string `json:"ip-address"`
	Prefix   int    `json:"prefix"`
}

type networks struct {
	Name     string  `json:"name"`
	Ip       []inter `json:"ip-addresses"`
	Hardware string  `json:"hardware-address"`
}

type test_net struct {
	Output networks `json:"return"`
}

func fget_ip(domain string) {

	cmd := exec.Command("virsh", "qemu-agent-command", domain, "{\"execute\":\"guest-network-get-interfaces\"}")

	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	var test test_net
	fmt.Printf("%v\n", out)

	json.Unmarshal(out.Bytes(), &test)

	fmt.Printf("translated phrase: %+v\n", test)

}

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
			name, err := item.GetName()
			fget_ip(name)
			fmt.Printf("[%v] [%v]\n", name, err)

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
