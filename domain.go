package main

import libvirtxml "github.com/libvirt/libvirt-go-xml"

func create_domain(name string) *libvirtxml.Domain {
	return &libvirtxml.Domain{
		Type:    "kvm",
		Name:    name,
		Memory:  &libvirtxml.DomainMemory{Value: 4096, Unit: "MB", DumpCore: "on"},
		VCPU:    &libvirtxml.DomainVCPU{Value: 1},
		CPU:     &libvirtxml.DomainCPU{Mode: "host-model"},
		OS:      &libvirtxml.DomainOS{Type: &libvirtxml.DomainOSType{Arch: "x86_64", Machine: "pc-i440fx-mantic", Type: "hvm"}},
		Devices: &libvirtxml.DomainDeviceList{},
	}

}
