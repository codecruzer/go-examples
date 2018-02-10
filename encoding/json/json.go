package json

import (
	"encoding/json"
)

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func unmarshalPerson(data []byte) (*person, error) {
	var p person
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return &p, nil
}

func unmarshalPeople(data []byte) ([]person, error) {
	var p []person
	if err := json.Unmarshal(data, &p); err != nil {
		return nil, err
	}

	return p, nil
}
