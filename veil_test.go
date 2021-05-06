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
	rules = append(rules, veil.NewRule("phone", veil.PatternNumber, veil.ActionObscureFunc))
	rules = append(rules, veil.NewRule("email", veil.PatternNumber, veil.ActionMaskFunc))
	rules = append(rules, veil.NewRule("address", "[a-z]+", veil.ActionObscureFunc))

	_, err := veil.NewVeil(rules)

	need := fmt.Errorf("veil: there are more than one occurrence of the '%s' rule", veil.PatternNumber)
	if err.Error() != need.Error() {
		t.Errorf(`Need: "%s", got: "%s"`, need, err)
	}
}

// TestProcess tests for all possibilities of the veil.Process functionality.
func TestProcess(t *testing.T) {

	// define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", veil.PatternNumber, veil.ActionMaskFunc))
	rules = append(rules, veil.NewRule("email", veil.PatternEmail, veil.ActionObscureFunc))

	// create new veil instance
	v, _ := veil.NewVeil(rules)

	// input data
	d := getInputData()
	fmt.Printf("Data: %+v\n", d)

	// process
	o, _ := v.Process(d)
	fmt.Printf("Output: %+v\n", o)
}
