package veil

import (
	"fmt"
)

func errDuplicateRule(rule string) error {
	return fmt.Errorf("veil: there are more than one occurrence of the '%s' rule", rule)
}
