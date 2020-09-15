package testing

import (
	"fmt"
	"reflect"
	"testing"
)

// Expected is used in test assertions, the Value parameter should be the value the
// test is expecting to be correct
type Expected struct {
	*testing.T
	Value interface{}
}

// AssertEqual takes two values of the same type and compares the value of built-in types (https://tour.golang.org/basics/11)
func (e Expected) AssertEqual(actual interface{}) bool {
	if !typeCheck(e.Value, actual) {
		e.Fatalf("AssertEqual must be performed on the same types, expected type: %v, actual type: %v", reflect.TypeOf(e.Value), reflect.TypeOf(actual))
	}

	return e.Value == actual
}

// quick check to see if interfaces are of the same type
func typeCheck(a interface{}, b interface{}) bool {
	fmt.Println(fmt.Sprintf("%v, %v", reflect.TypeOf(a), reflect.TypeOf(b)))
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
