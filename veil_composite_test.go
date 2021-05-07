package veil_test

import (
	"fmt"
	"testing"

	"github.com/kosatnkn/veil"
)

func TestStruct(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", veil.PatternNumber, veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	d := getStructData()

	// process
	o, err := v.Process(d...)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	fmt.Printf("Output: %+v\n", o)
}

func TestMap(t *testing.T) {
	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", veil.PatternNumber, veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	m := make(map[string]string)
	m["key1"] = "123"
	m["key2"] = "abc"
	m["key3"] = "abc123"

	// process
	o, err := v.Process(m)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	fmt.Printf("Output: %+v\n", o)
}
