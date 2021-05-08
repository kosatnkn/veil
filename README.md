# Veil

![Open Issues](https://img.shields.io/github/issues/kosatnkn/veil)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/kosatnkn/veil)
[![Go Reference](https://pkg.go.dev/badge/github.com/kosatnkn/veil.svg)](https://pkg.go.dev/github.com/kosatnkn/veil)
[![Go Report Card](https://goreportcard.com/badge/github.com/kosatnkn/veil)](https://goreportcard.com/report/github.com/kosatnkn/veil)

A data de-identification library written in Go.

Veil can be used to obscure sensitive data in data structures when they are being persisted. This is especially useful
when writing to log files.

## Usage

In order to use `veil` you need to create a veil instance. To create a new `veil` instance first you need to create
a set of `veil.Rule`s.

A `veil,Rule` consists of a **name**, a **regex pattern** to match against and an **action function** that will be
executed against the match.
```go
rule := veil.NewRule("phone", veil.PatternNumber, veil.ActionObscureFunc)
```

`Veil` has a set of pre-defined **regex patterns** and **action functions** that you can use out of the box.
You can also use your own regex patterns and action functions.
```go
rule := veil.NewRule("email",
    `^([a-zA-Z0-9._%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6})*$`,
    func(in string) string {
        mask := "****"
        l := len(in)
        if l <= 3 {
            return mask + in
        }
        return mask + string(in[l-3:])
    })
```

A new `veil` instance is created by passing in a slice of `veil.Rule`s in to `veil.NewVeil` function.
Once a new instance is created use the `Process` function to process data.
```go
package main

func main() {
    // define rules
	var rules []veil.Rule
	rules = append(rules, veil.NewRule("phone", veil.PatternNumber, veil.ActionObscureFunc))
	rules = append(rules, veil.NewRule("email", veil.PatternEmail, veil.ActionMaskFunc))

	// create new veil instance
	v, err := veil.NewVeil(rules)
    if err != nil {
        panic(err)
    }

    // process data
    o, err := v.Process("This text contains a phone number 0712345678",
        "This text contains an email address test@example.com")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%#v", o)
}
```

This will output
```text
[]interface {}{"This text contains a phone number ***", "This text contains an email address ****com"} 
```

## Limitations

`veil.Process` will not return the same data type that you pass in to it.
Following is the list of **in**, **out** types. 

| In Type   | Out Type                      |
| ---       | ---                           |
| string    | string                        |
| int       | string                        |
| uint      | string                        |
| float     | string                        |
| struct    | map[string]interface{}        |
| map       | map[interface{}]interface{}   |
| array     | []interface{}                 |
| slice     | []interface{}                 |

Any other data type will be first converted in to a string using `fmt.Sprintf("%+v", input)` and then processed.

## Using Struct Tags

`veil` supports struct tags. You can use them to define how a field is de-identified.
```go
type User struct {
    User  string
    Pwd   string `veil:"obscure"`
    group float32
}
```

Following are the tag options that are available.
- `obscure` to obscure the field value by using the **default obscure function**
- `mask` to mask the field value by using the **default masking function**
- `hide` to remove the field from the returned data structure

## Execution Precedence

`veil` will use the following precedence when de-identifying data.
- If a `stringer` interface is implemented for the data type it will be used
- If the type is a struct and `veil` struct tags are used use the struct tag to process the field
- Use the ruleset passed when creating the `veil` instance
