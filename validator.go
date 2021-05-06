package veil

// validate validates the rule set.
func validate(rules []Rule) error {

	// check for duplicates
	err := checkDuplicateRules(rules)
	if err != nil {
		return err
	}

	return nil
}

// checkDuplicateRules checks to see whether the same regexp pattern is used in more than one rules.
func checkDuplicateRules(rules []Rule) error {

	for _, rule := range rules {
		if isPatternDuplicated(rules, rule.pattern) {
			return errDuplicateRule(rule.pattern)
		}
	}

	return nil
}

// isPatternDuplicated checks whether there are more than one occurrence of the pattern in the rule set.
func isPatternDuplicated(rules []Rule, pattern string) bool {

	count := 0

	for _, rule := range rules {
		if pattern == rule.pattern {
			count++
		}
	}

	return count > 1
}
