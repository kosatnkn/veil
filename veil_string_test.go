package veil_test

import (
	"fmt"
	"testing"

	"github.com/kosatnkn/veil"
)

// TestStringObscure tests for the data obscure action performed against given string.
func TestStringObscure(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("number", veil.PatternNumber, veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %#v\n", d)

	// process
	o, _ := v.Process(d...)
	fmt.Printf("Output: %#v\n", o)
}

// TestStringMask tests for the data mask action performed against given string.
func TestStringMask(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("number", veil.PatternNumber, veil.ActionMaskFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %#v\n", d)

	// process
	o, _ := v.Process(d...)
	fmt.Printf("Output: %#v\n", o)
}
