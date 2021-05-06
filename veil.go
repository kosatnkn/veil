package veil

import (
	"fmt"
)

// Veil represents a veil instance.
type Veil struct {
	rules []Rule
}

// NewVeil creates a new veil instance.
func NewVeil(rules []Rule) Veil {
	return Veil{
		rules: rules,
	}
}

// Process returns a processed set of inputs against the rule set.
func (v *Veil) Process(inputs ...interface{}) (outputs []interface{}, err error) {
	for _, input := range inputs {
		s, err := v.processString(fmt.Sprintf("%+v", input))
		if err != nil {
			return nil, err
		}

		outputs = append(outputs, s)
	}

	return
}

// processString processes the given string against the attached rule set.
func (v *Veil) processString(input string) (string, error) {
	for _, rule := range v.rules {
		input = rule.Pattern().ReplaceAllStringFunc(input, rule.ActionFunc())
	}

	return input, nil
}
