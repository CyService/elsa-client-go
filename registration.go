package reg


type Instance struct {

	Location string `json:"location"`
	Capacity interface{}    `json:"capacity"`

}


type Registration struct {

	Service   string `json:"service"`
	Version   string `json:"version"`
	Instances []Instance `json:"instances"`

}
