package cloudmeta

import "github.com/spotmaxtech/gokit"

type RegionInfo struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type Region interface {
	Fetch(consul *gokit.Consul) error
	List() []*RegionInfo
	GetRegionInfo(name string) *RegionInfo
}

// TODO: more info item?
// TODO: make category const?
type InstInfo struct {
	Name    string  `json:"name"`
	Core    int8    `json:"core"`
	Mem     float64 `json:"mem"`
	Storage string  `json:"storage"`
	Family  string  `json:"family"`
}

type Instance interface {
	Fetch(consul *gokit.Consul) error
	List(region string) []*InstInfo
	GetInstInfo(region string, name string) *InstInfo
}

type ODPrice interface {
	Fetch(consul *gokit.Consul) error
	GetPrice(region string, instance string) float32
}

type SpotPrice interface {
	Fetch(consul *gokit.Consul) error
	GetPrice(region string, instance string) float32
}

type InterruptInfo struct {
	Name     string `json:"name"`
	Rate     int    `json:"rate"`
	RateDesc string `json:"rate_desc"`
}

type Interrupt interface {
	Fetch(consul *gokit.Consul) error
	GetInterruptInfo(region string, instance string) *InterruptInfo
}
