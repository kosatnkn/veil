package veil

import (
	"fmt"
	"reflect"
)

// Veil represents a veil instance.
type Veil struct {
	rules []Rule
}

// NewVeil creates a new veil instance.
func NewVeil(rules []Rule) (Veil, error) {

	err := validate(rules)
	if err != nil {
		return Veil{}, err
	}

	return Veil{
		rules: rules,
	}, nil
}

// Process returns a processed set of inputs against the rule set.
func (v *Veil) Process(inputs ...interface{}) ([]interface{}, error) {

	var outputs []interface{}

	for _, input := range inputs {
		p, err := v.process(input)
		if err != nil {
			return nil, err
		}
		outputs = append(outputs, p)
	}

	return outputs, nil
}

// process processes the given type.
func (v *Veil) process(input interface{}) (interface{}, error) {

	switch input.(type) {
	case string,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return v.processString(fmt.Sprintf("%v", input))
	default:
		return v.processComposite(input)
	}
}

// ProcessString processes the given string against the rule set.
func (v *Veil) processString(input string) (string, error) {

	for _, rule := range v.rules {
		input = rule.patternRx.ReplaceAllStringFunc(input, rule.action)
	}

	return input, nil
}

// processComposite processes composite types such as structs, maps and slices.
func (v *Veil) processComposite(input interface{}) (interface{}, error) {

	switch reflect.TypeOf(input).Kind() {
	case reflect.Struct:
		return v.processStruct(input)
	case reflect.Map:
		return v.processMap(input)
	case reflect.Array,
		reflect.Slice:
		return v.processSlice(input)
	default:
		return v.process(fmt.Sprintf("%+v", input))
	}
}

// processStruct process the given struct.
//
// Structs will be converted to maps for processing convenience.
// Precedence is given to veil tags of struct fields if there are any.
func (v *Veil) processStruct(input interface{}) (interface{}, error) {

	s := make(map[string]interface{})

	typ := reflect.TypeOf(input)
	val := reflect.ValueOf(input)

	for i := 0; i < val.NumField(); i++ {

		fTyp := typ.Field(i)
		fVal := val.Field(i)

		// process tags
		switch fTyp.Tag.Get("veil") {
		case tagHide:
			continue
		case tagObscure:
			s[fTyp.Name] = ActionObscureFunc("")
			continue
		case tagMask:
			s[fTyp.Name] = ActionMaskFunc(fmt.Sprintf("%v", fVal))
			continue
		default:
			var f interface{}

			if fVal.CanInterface() {
				f = fVal.Interface()
			} else {
				f = fmt.Sprintf("%v", fVal)
			}

			out, err := v.process(f)
			if err != nil {
				return nil, err
			}

			s[fTyp.Name] = out
		}
	}

	return s, nil
}

// processMap process the given map.
func (v *Veil) processMap(input interface{}) (interface{}, error) {

	m := make(map[interface{}]interface{})

	iter := reflect.ValueOf(input).MapRange()
	for iter.Next() {
		key := iter.Key()
		val := iter.Value()

		out, err := v.process(val.Interface())
		if err != nil {
			return nil, err
		}

		m[key.Interface()] = out
	}

	return m, nil
}

// processSlice process the given slice.
func (v *Veil) processSlice(input interface{}) (interface{}, error) {

	var s []interface{}

	val := reflect.ValueOf(input)

	for i := 0; i < val.Len(); i++ {
		out, err := v.process(val.Index(i).Interface())
		if err != nil {
			return nil, err
		}

		s = append(s, out)
	}

	return s, nil
}
