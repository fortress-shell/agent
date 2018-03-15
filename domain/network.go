package domain

type AgentNetworkCommand struct {
	Execute string `json:"execute"`
}

type NetworkInterface struct {
	Prefix        uint32 `json:"prefix"`
	IpAddress     string `json:"ip-address"`
	IpAddressType string `json:"ip-address-type"`
}

type NetworkInterfaces struct {
	Name            string             `json:"name"`
	HardwareAddress string             `json:"hardware-address"`
	IpAddresses     []NetworkInterface `json:"ip-addresses"`
}

type AgentNetworkCommandReply struct {
	Return []NetworkInterfaces `json:"return"`
}
