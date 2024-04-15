package main

import (
	"encoding/json"
	"fmt"
	"log"
)

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
