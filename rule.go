package veil

// Rule represents a de-identification rule.
type Rule struct {
	Name    string
	Pattern string // regex pattern
	Action  action
}

// NewRule creates a new de-identification rule.
func NewRule(name, pattern string, action action) Rule {
	return Rule{
		Name:    name,
		Pattern: pattern,
		Action:  action,
	}
}
