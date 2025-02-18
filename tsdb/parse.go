package tsdb

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

func parse(path string) (*string, error) {
	res, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, errors.New("file not existed. create a new chain")
	}
	str := string(res)
	return &str, nil
}

func loadFromStorage(raw *string) *[]BlockJSON {
	inst := []BlockJSON{}
	b := []byte(*raw)
	e := json.Unmarshal(b, &inst)
	fmt.Println(inst)
	if e != nil {
		panic(e)
	}
	return &inst
}

func saveToHDD(path string, data []byte) error {
	e := ioutil.WriteFile(path, data, 0644)
	if e != nil {
		return e
	}
	return nil
}

// Parser type for accessing parsing functions
type Parser struct{}

// Parse acts as an global interface for converting different types into the required structure
type Parse interface {
	ParseToJSON(a []Block) []byte
}

// ParseToJSON converts the chain into Marshallable JSON
func (p Parser) ParseToJSON(a []Block) (j []byte) {
	b := []BlockJSON{}

	for _, inst := range a {
		t := BlockJSON{
			Timestamp:      inst.Timestamp,
			NormalizedTime: inst.NormalizedTime,
			Datapoint:      inst.Datapoint,
		}
		b = append(b, t)
	}

	j, e := json.Marshal(b)
	if e != nil {
		panic(e)
	}
	return
}
