package veil_test

// getInputData prepares the test dataset.
func getInputData() []interface{} {

	var data []interface{}

	// string
	data = append(data, "hello 123")

	// struct
	data = append(data, struct {
		Name  string
		Phone string
		ph    string
	}{"Test", "123", "123"})

	// map
	data = append(data, map[string]int{
		"phone": 123,
	})

	return data
}
