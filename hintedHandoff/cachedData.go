package hintedHandoff

type KeyValuePair struct {
	Type string `json: type`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// number of nodes, including temprarily down nodes
const (
	// todo - this number is hardcoded. it is the number of nodes created in the docker-compose file
	HARDCODED_NUMBER_OF_NODES int = 5
)

var (
	CachedData = make(map[int][]KeyValuePair)
)
