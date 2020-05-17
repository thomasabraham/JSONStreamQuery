package jsq


type JSONStreamQuerier interface {
	Query(input []interface{}) []interface{}
}

type JSONStreamQuery struct {
}

func BuildJSONStreamQuerier(query string) JSONStreamQuerier {
	return JSONStreamQuery{}
}

func (JSONStreamQuery) Query(input []interface{}) []interface{} {
	return input
}
