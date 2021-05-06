package veil

import "regexp"

// Rule represents a de-identification rule.
type Rule struct {
	name      string
	pattern   string
	patternRx *regexp.Regexp
	action    ActionFn
}

// NewRule creates a new de-identification rule.
func NewRule(name, pattern string, fn ActionFn) Rule {
	return Rule{
		name:      name,
		pattern:   pattern,
		patternRx: regexp.MustCompile(pattern),
		action:    fn,
	}
}
