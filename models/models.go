package models

// Distributor represents a distributor
type Distributor struct {
	ID             int      `json:"id" avro:"id"`
	Name           string   `json:"name" avro:"name"`
	Include        []string `json:"includes" avro:"includes"`
	Exclude        []string `json:"excludes" avro:"excludes"`
	SubDistributor bool     `json:"sub_distributor" avro:"sub_distributor"`
	Parent         string   `json:"parent" avro:"parent"`
}

// DistributorData represents the data for the distributor
type DistributorsModel struct {
	CountryStateMap    map[string][]string
	StateCityMap       map[string][]string
	CurrentDistributor Distributor
	Distributors       []Distributor
}
