package veil_test

import (
	"fmt"
	"testing"

	"github.com/kosatnkn/veil"
)

func TestNewVeil(t *testing.T) {
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", "123", veil.ActionObscure))

	v := veil.NewVeil(rules)
	fmt.Printf("%+v", v)
}
