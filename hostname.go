package main

import "encoding/json"

func get_hostname(domain string) (string, error) {
	// virsh qemu-agent-command master01.autok8s.xyz '{"execute":"guest-info"}'
	out, err := cmd(domain, "guest-get-host-name")
	if err != nil {
		return "", err
	}
	var hostname struct {
		Return struct {
			Hostname string `json:"host-name"`
		} `json:"return"`
	}

	json.Unmarshal(out, &hostname)

	return hostname.Return.Hostname, nil
}

func set_hostname(domain string) (bool, error) {
	// virsh qemu-agent-command master01.autok8s.xyz '{"execute":"guest-exec","arguments":{"path":"hostnamectl","arg":["hostname","master01"]}}'
	args := QAX_args{
		Path: "hostnamectl",
		Arg:  []string{"hostname", domain},
	}

	_, err := cmd_param(domain, "guest-exec", args)
	if err != nil {

		return false, err
	}

	return true, nil
}
