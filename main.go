package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strings"
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
	Execute   string            `json:"execute"`
	Arguments map[string]string `json:"arguments,omitempty"`
}

func cmd(domain string, execute string) ([]byte, error) {
	test1 := command{Execute: execute}
	execommand, _ := json.Marshal(test1)
	cmd := exec.Command("virsh", "qemu-agent-command", domain, fmt.Sprintf("%+v", string(execommand)))
	var out, err bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func cmd_param(domain string, execute string, arg map[string]string) ([]byte, error) {
	test1 := command{Execute: execute, Arguments: arg}
	execommand, _ := json.Marshal(test1)
	cmd := exec.Command("virsh", "qemu-agent-command", domain, fmt.Sprintf("%+v", string(execommand)))
	var out, err1 bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &err1
	if err := cmd.Run(); err != nil {
		fmt.Printf("log_tmp: %v\n", err1.String())
		return err1.Bytes(), err
	}

	return out.Bytes(), nil
}

func fget_ip(domain string) {
	out, err := cmd(domain, "guest-network-get-interfaces") // https://qemu-project.gitlab.io/qemu/interop/qemu-ga-ref.html
	if err != nil {
		return
	}
	var test test_net
	err = json.Unmarshal(out, &test)
	if err != nil {
		log.Fatal(err)
	}

	for _, item := range test.Return {
		fmt.Printf("Name:[%+v] MAC:[%+v]\n", item.Name, item.Hardware)
		for _, ip := range item.Ip {
			fmt.Printf(" __ IP: [%v] [%v] [%v]\n", ip.IPType, ip.IpAddres, ip.Prefix)
		}
	}

}
func fget_hostname(domain string) {
	// virsh qemu-agent-command master01.autok8s.xyz '{"execute":"guest-ssh-add-authorized-keys","arguments":{"username":"root","keys":["key"]}}'
}

func fget_sshauth_key(domain string) {
	// virsh qemu-agent-command master01.autok8s.xyz '{"execute":"guest-ssh-get-authorized-keys","arguments":{"username":"root"}}'
	args := map[string]string{"username": "roo1t"}
	out, err := cmd_param(domain, "guest-ssh-get-authorized-keys", args)
	if err != nil {
		fmt.Printf("error : fget_sshauth_key :%+v [%v]", err, out)
		return
	}
	fmt.Printf("%+v", out)
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
			// bo, err := item.IsActive()
			// fmt.Printf("[%v] [%v]\n", bo, err)
			name, err := item.GetName()
			if !strings.Contains(name, "autok8s") {
				continue
			}
			fmt.Printf("---------------[%v] [%v]\n", name, err)
			fget_ip(name)
			fget_sshauth_key(name)

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
