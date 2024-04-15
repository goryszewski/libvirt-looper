package main

import "fmt"

func fadd_sshauth_key(domain string) bool {
	args := SSH_ADD{
		Username: "root",
		Keys:     []string{key},
		Reset:    true,
	}
	out, err := cmd_param(domain, "guest-ssh-add-authorized-keys", args)
	if err != nil {
		fmt.Printf("error : fget_sshauth_key :%+v [%v]", err, out)
		return false
	}
	fmt.Printf("[fadd_sshauth_key]%+v \n", out)
	return true

}

func fget_sshauth_key(domain string) bool {
	// virsh qemu-agent-command master01.autok8s.xyz '{"execute":"guest-ssh-get-authorized-keys","arguments":{"username":"root"}}'
	args := map[string]string{"username": "root"}
	out, err := cmd_param(domain, "guest-ssh-get-authorized-keys", args)
	if err != nil {
		fmt.Printf("error : fget_sshauth_key :%+v [%v]", err, out)
		return false
	}
	fmt.Printf("%+v\n", out)
	return true
}
