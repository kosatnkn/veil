package veil

import "regexp"

// Rule represents a de-identification rule.
type Rule struct {
	name    string
	pattern *regexp.Regexp
	action  ActionFn
}

// NewRule creates a new de-identification rule.
func NewRule(name, pattern string, fn ActionFn) Rule {
	return Rule{
		name:    name,
		pattern: regexp.MustCompile(pattern),
		action:  fn,
	}
}

// Name returns the name of the rule.
func (r *Rule) Name() string {
	return r.name
}

// Pattern returns the pattern matcher of the rule.
func (r *Rule) Pattern() *regexp.Regexp {
	return r.pattern
}

// ActionFunc returns the function containing the action to perform against the rule.
func (r *Rule) ActionFunc() ActionFn {
	return r.action
}
