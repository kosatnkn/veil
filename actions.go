package veil

type ActionFn func(in string) string

var ActionObscureFunc = func(in string) string {
	return "####"
}
