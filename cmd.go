package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
)

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

func cmd_param(domain string, execute string, arg any) ([]byte, error) {
	test1 := command{Execute: execute, Arguments: arg}
	execommand, _ := json.Marshal(test1)
	fmt.Printf("[%v]\n", string(execommand))
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
