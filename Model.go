package main

var key string = "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQCznq3WHI8Xwa7d6T1JZRC0PDenaMH24tt5la94kVGW/09kDliNW1zukl/SJn77QTfli9gSWPlqRbdkF6+QfbOdeH/TsflsFLzNdfNcfXcmnwz84K6cpHHUygfyHVRYorLP0lQhx/3AzAzwzRgn7Uu+Rr09lF+d1tZuyziJnE7OBMLfUVo9RpVvw4O7wmYl07GrE1anXP299g3Ra3hXEwD7Rkldt5kSSiKG0gRC0DPBZCDJ6hqDmt7V8XL4D1hrooOEsokLoAX6alHB5j9J1Yr95prRH1A26Rankx7DsJnTL7lCh9PkL4KVPWVqIysRvN0k3Jq4WU7jknRetA+7sLVoDb6Z8m89mhhPejJfifAJiohddT6k86MEwyaWVbzSUGON2v+BRsLFDng/qubpR4tMy+YyseLuz94A+aAmdpxFenQRv5QCOeHEQEnsm9EzcXf81iXoydRgatztXrIAf5wpzv497pz+sEu0jFuUGDdFhLmiZOX7D6cg5SQeHe8IsMM= root@core"

type SSH_ADD struct {
	Username string   `json:"username"`
	Keys     []string `json:"keys"`
	Reset    bool     `json:"reset"`
}

type QAX_args struct {
	Path string   `json:"path"`
	Arg  []string `json:"arg"`
}

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
	Execute   string `json:"execute"`
	Arguments any    `json:"arguments,omitempty"`
}
