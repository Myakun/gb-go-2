package hydrator_test

import (
	"lesson7/hydrator"
	"testing"
)

type customStruct struct {
	IntVal     int
	unexported interface{}
}

func TestHydrateInIsNil(t *testing.T) {
	values := getHydratorValues()
	err := hydrator.Hydrate(nil, &values)
	if nil == err {
		t.Error("Incorrect return value. Expected error, got nil")
	}
}

func TestHydrateInIsNotStruct(t *testing.T) {
	values := getHydratorValues()
	err := hydrator.Hydrate(1, &values)
	if nil == err {
		t.Error("Incorrect return value. Expected error, got nil")
	}
}

func TestHydrateSuccess(t *testing.T) {
	myCustomStruct := customStruct{}
	values := getHydratorValues()
	err := hydrator.Hydrate(&myCustomStruct, &values)
	if nil != err {
		t.Error(err)
	}
}

func TestHydrateWithIncorrectType(t *testing.T) {
	myCustomStruct := customStruct{}
	values := getHydratorValues()
	values["IntVal"] = "10"
	err := hydrator.Hydrate(&myCustomStruct, &values)
	if nil == err {
		t.Error("Incorrect return value. Expected error, got nil")
	}
}

func TestHydrateWithUnexportedField(t *testing.T) {
	myCustomStruct := customStruct{}
	values := getHydratorValues()
	values["unexported"] = 1
	err := hydrator.Hydrate(&myCustomStruct, &values)
	if nil == err {
		t.Error("Incorrect return value. Expected error, got nil")
	}
}

func getHydratorValues() map[string]interface{} {
	values := make(map[string]interface{})
	values["IntVal"] = 10

	return values
}
