package veil

// ActionFn represents function to run against the rule.
type ActionFn func(in string) string

// ActionObscureFunc is the default obscure function.
// It replaces the matching string with '***'.
var ActionObscureFunc = func(in string) string {
	return "***"
}

// ActionMaskFunc is the default masking function.
// It masks the input string as follows.
//
// | Input			| Output	|
// | -------------- | --------- |
// | 1				| ****1		|
// | 12				| ****23	|
// | 123  			| ****123	|
// | 1234 			| ****234	|
// | somelongstring | ****ing	|
var ActionMaskFunc = func(in string) string {

	mask := "****"

	if len(in) <= 3 {
		return mask + in
	}

	return mask + string(in[len(in)-3])
}
