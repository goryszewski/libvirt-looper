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
	Return []networks `json:"return"`
}

type command struct {
	Execute string `json:"execute"`
}

func cmd(domain string, execute string) []byte {
	test1 := command{Execute: execute}
	execommand, _ := json.Marshal(test1)
	cmd := exec.Command("virsh", "qemu-agent-command", domain, fmt.Sprintf("%+v", string(execommand)))
	var out bytes.Buffer
	cmd.Stdout = &out
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
	return out.Bytes()
}

func fget_ip(domain string) {
	out := cmd(domain, "guest-network-get-interfaces") // https://qemu-project.gitlab.io/qemu/interop/qemu-ga-ref.html
	var test test_net
	err := json.Unmarshal(out, &test)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range test.Return {
		fmt.Printf("[%v]\n", item)
	}

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
			fmt.Printf("---------------[%v] [%v]\n", name, err)
			fget_ip(name)

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
