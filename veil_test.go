package veil_test

import (
	"fmt"
	"testing"

	"github.com/kosatnkn/veil"
)

// TestNewVeilWithDuplicateRules checks for duplicate rule validation when creating a new instance of veil.
func TestNewVeilWithDuplicateRules(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "[0-9]+", veil.ActionObscureFunc))
	rules = append(rules, veil.NewRule("email", "[0-9]+", veil.ActionMaskFunc))
	rules = append(rules, veil.NewRule("address", "[a-z]+", veil.ActionObscureFunc))

	_, err := veil.NewVeil(rules)

	need := fmt.Errorf("veil: there are more than one occurrence of the '%s' rule", "[0-9]+")
	if err.Error() != need.Error() {
		t.Errorf(`Need: "%s", got: "%s"`, need, err)
	}
}

// TestObscure tests for the data obscure action performed against given data.
func TestObscure(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "[0-9]+", veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %+v\n", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("Output: %+v\n", o)
}

// TestMask tests for the data mask action performed against given data.
func TestMask(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "[0-9]+", veil.ActionMaskFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %+v\n", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("Output: %+v\n", o)
}

// TestAll tests for all possibilities of the veil.Process functionality.
func TestAll(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "[0-9]+", veil.ActionMaskFunc))
	rules = append(rules, veil.NewRule("email", "[0-9]+", veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %+v\n", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("Output: %+v\n", o)
}

// getInputData prepares the test dataset.
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
