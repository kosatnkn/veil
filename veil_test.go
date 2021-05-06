package veil_test

import (
	"fmt"
	"testing"

	"github.com/kosatnkn/veil"
)

func TestObscure(t *testing.T) {
	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "[0-9]+", veil.ActionObscureFunc))

	// create new veil instance
	v := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %+v\n", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("Output: %+v\n", o)
}

func TestMask(t *testing.T) {
	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "[0-9]+", veil.ActionMaskFunc))

	// create new veil instance
	v := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %+v\n", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("Output: %+v\n", o)
}

func TestAll(t *testing.T) {
	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "[0-9]+", veil.ActionMaskFunc))
	rules = append(rules, veil.NewRule("email", "[0-9]+", veil.ActionObscureFunc))

	// create new veil instance
	v := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %+v\n", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("Output: %+v\n", o)
}

func getInputData() []interface{} {
	var data []interface{}

	// string
	data = append(data, "hello 123")

	// struct
	data = append(data, struct {
		Name  string
		Phone string
		ph    string
	}{"Test", "123", "123"})

	// map
	data = append(data, map[string]int{
		"phone": 123,
	})

	return data
}
