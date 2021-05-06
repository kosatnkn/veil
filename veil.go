package veil

import (
	"fmt"
)

// Veil represents a veil instance.
type Veil struct {
	rules []Rule
}

// NewVeil creates a new veil instance.
func NewVeil(rules []Rule) (Veil, error) {

	err := validate(rules)
	if err != nil {
		return Veil{}, err
	}

	return Veil{
		rules: rules,
	}, nil
}

// Process returns a processed set of inputs against the rule set.
func (v *Veil) Process(inputs ...interface{}) ([]interface{}, error) {

	var outputs []interface{}

	for _, input := range inputs {

		s, err := v.ProcessString(fmt.Sprintf("%+v", input))
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, s)
	}

	return outputs, nil
}

// ProcessString processes the given string against the rule set.
func (v *Veil) ProcessString(input string) (string, error) {

	for _, rule := range v.rules {
		input = rule.patternRx.ReplaceAllStringFunc(input, rule.action)
	}

	return input, nil
}

// // ProcessStruct processes the given struct against the rule set.
// func (v *Veil) ProcessStruct(input struct{}) (struct{}, error) {

// }
