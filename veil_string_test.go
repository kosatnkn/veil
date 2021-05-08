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

// TestNumberRegex tests for default regex to identify any number.
func TestNumberRegex(t *testing.T) {

	tcs := []struct {
		name string
		in   interface{}
		out  string
	}{
		// non existant
		{"int as string", "123456", "***"},
		{"neg int as string", "-123456", "***"},
		{"int in string start ", "123456abcdef", "***abcdef"},
		{"int in string mid ", "abc123456def", "abc***def"},
		{"int in string end ", "abcdef123456", "abcdef***"},
		{"neg int in string start ", "123456abcdef", "***abcdef"},
		{"neg int in string mid ", "abc123456def", "abc***def"},
		{"neg int in string end ", "abcdef123456", "abcdef***"},
		{"int", 123456, "***"},
		{"neg int", -123456, "***"},
		{"float", 123.456, "***"},
		{"neg float", -123.456, "***"},
	}

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("number", veil.PatternNumber, veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	for _, tc := range tcs {

		t.Run(tc.name, func(t *testing.T) {

			out, err := v.Process(tc.in)
			if err != nil {
				t.Error(err)
			}

			got, ok := out[0].(string)
			if !ok {
				t.Error("assertion error")
			}

			t.Logf("output: %v", out)

			if got != tc.out {
				t.Errorf("need: %v, got: %v", tc.out, got)
			}
		})
	}
}
