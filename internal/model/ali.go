package model

type AliEcsPageInput struct {
	RegionId string
	// PublicIpAddresses  string `json:"publicIpAddresses,omitempty"`
	// PrivateIpAddresses string `json:"privateIpAddresses,omitempty"`
	// InstanceName       string `json:"instanceName,omitempty"`
	DynamicSelectKey   string `json:"DynamicSelectKey,omitempty"`
	DynamicSelectValue string `json:"DynamicSelectValue,omitempty"`
	CurrentPage        int
	PageSize           int
}

// type AliEcsPageOutput struct {
// 	HostName        string `json:"hostName"`
// 	InstanceId      string `json:"instanceId"`
// 	OSName          string `json:"osName"`
// 	Cpu             int    `json:"cpu"`
// 	Memory          int    `json:"memory"`
// 	PublicIpAddress string `json:"publicIpAddress"`
// 	Status          string `json:"status"`
// 	ZoneId          string `json:"zoneId"`
// 	CreationTime    string `json:"creationTime"`
// 	ExpiredTime     string `json:"expiredTime"`
// }
