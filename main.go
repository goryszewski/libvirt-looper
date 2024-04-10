package main

import (
	"fmt"
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
			name, err := item.GetName()
			fmt.Printf("[%v] [%v]\n", name, err)
		}

		time.Sleep(time.Second * 5)
	}
	defer conn.Close()
}
