package veil_test

import (
	"fmt"
	"testing"

	"github.com/kosatnkn/veil"
)

func TestNewVeil(t *testing.T) {
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "123", veil.ActionObscureFunc))

	v := veil.NewVeil(rules)
	fmt.Printf("%+v", v)
}

func TestProcess(t *testing.T) {
	// define rules
	var rules []veil.Rule
	fn := func(in string) string { return "#*#*#*" }
	rules = append(rules, veil.NewRule("phone", "123", fn))

	// create new veil instance
	v := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("%+v", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("%+v", o)
}

func getInputData() []interface{} {
	return []interface{}{
		"hello 123",
		struct {
			Name  string
			Phone string
			ph    string
		}{"Test", "123", "123"},
		map[string]int{
			"phone": 123,
		},
	}
}
