package json

import (
	"reflect"
	"testing"
)

func TestUnmarshalPerson(t *testing.T) {
	data := []byte(`
		{
			"name": "Jake",
			"age": 28
		}
	`)

	expected := &person{
		Name: "Jake",
		Age:  28,
	}

	result, err := unmarshalPerson(data)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestUnmarshalPersonBadJSON(t *testing.T) {
	_, err := unmarshalPerson([]byte(`"not a person"`))
	if err == nil {
		t.Error("Expected error")
	}
}

func TestUnmarshalPeople(t *testing.T) {
	data := []byte(`
		[
			{
				"name": "JB",
				"age": 27
			},
			{
				"name": "Brent",
				"age": 28
			},
			{
				"name": "Dustin",
				"age": 25
			}
		]
	`)

	expected := []person{
		person{
			Name: "JB",
			Age:  27,
		},
		person{
			Name: "Brent",
			Age:  28,
		},
		person{
			Name: "Dustin",
			Age:  25,
		},
	}

	result, err := unmarshalPeople(data)
	if err != nil {
		t.Errorf("Expected nil error, got %v", err)
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %+v, got %+v", expected, result)
	}
}

func TestUnmarshalPeopleBadJSON(t *testing.T) {
	_, err := unmarshalPeople([]byte(`"not a slice of people"`))
	if err == nil {
		t.Error("Expected error")
	}
}
