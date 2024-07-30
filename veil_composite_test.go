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

func TestSlice(t *testing.T) {
	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", veil.PatternNumber, veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	var s []string
	s = append(s, "123")
	s = append(s, "abc")
	s = append(s, "abc123")

	// process
	o, err := v.Process(s)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	fmt.Printf("Output: %+v\n", o)
}

func TestArray(t *testing.T) {
	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", veil.PatternNumber, veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	var a [3]string
	a[0] = "123"
	a[1] = "abc"
	a[2] = "abc123"

	// process
	o, err := v.Process(a)
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}

	fmt.Printf("Output: %+v\n", o)
}
