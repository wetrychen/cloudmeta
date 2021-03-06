package main

import (
	"encoding/json"
	"github.com/spotmaxtech/gokit"
)

// factory for now manually
const (
	ConsulAddr = "consul.spotmaxtech.com"
	RegionKey  = "cloudmeta/aws/region.json"
)

func main() {
	// consul
	consul := gokit.NewConsul(ConsulAddr)

	type MsData struct {
		Text string `json:"text"`
	}
	data := make(map[string]*MsData)
	data["us-east-1"] = &MsData{
		Text: "US East (N. Virginia)",
	}
	data["us-east-2"] = &MsData{
		Text: "US East (Ohio)",
	}
	data["us-west-2"] = &MsData{
		Text: "US West (Oregon)",
	}
	data["ap-northeast-2"] = &MsData{
		Text: "Asia Pacific (Seoul)",
	}
	data["ap-southeast-1"] = &MsData{
		Text: "Asia Pacific (Singapore)",
	}
	data["eu-central-1"] = &MsData{
		Text: "EU (Frankfurt)",
	}
	data["eu-west-3"] = &MsData{
		Text: "EU (Paris)",
	}

	bytes, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		panic(err)
	}
	if err := consul.PutKey(RegionKey, bytes); err != nil {
		panic(err)
	}
}
