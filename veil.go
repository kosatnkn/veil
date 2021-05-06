package veil

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
		outputs = append(outputs, input)
	}

	return
}
