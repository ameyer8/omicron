package data

//DirectionReturn is the JSON structure for return port data
type DirectionReturn struct {
	Direction string `json:"direction"`
}

//ValueReturn is the JSON structure for return port data
type ValueReturn struct {
	Value int `json:"value"`
}

type StatusReturn struct {
	Direction string `json:"direction"`
	Value     int    `json:"value"`
	Name      string `json:"name"`
}

type PortName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
