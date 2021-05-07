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
